package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hero-soft/web-scanner/pkg/httpservice"
	"github.com/hero-soft/web-scanner/pkg/settings"
	"github.com/hero-soft/web-scanner/pkg/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// BUILD is the build number of the application
	BUILD = 100
	// COMMIT is the Git commit hash
	COMMIT = "XXXXXXXXXXXXXXXXXXXXXXXX"
)

type application struct {
	logger *zap.SugaredLogger
	//quit              chan struct{}
	metricsPort       string
	serviceHTTPPort   string
	permissiveHeaders bool
	counters          map[string]prometheus.Counter
	SendChan          chan<- websocket.SendTo
}

func main() {

	mySettings, err := settings.New()

	if err != nil {
		fmt.Printf("Could not get settings: %v\n", err)
		os.Exit(1)
	}

	app := application{
		permissiveHeaders: mySettings.GetBool("server.permissive_headers"),
		metricsPort:       mySettings.GetString("server.metrics_port"),
		serviceHTTPPort:   mySettings.GetString("server.service_http_port"),
		counters:          make(map[string]prometheus.Counter),
	}

	app.setupLogging()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		osSignal := make(chan os.Signal)
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

		for {
			select {
			case <-osSignal:
				app.logger.Infof("Quit signal received...")
				cancel()
				//app.wg.Done()
				return
			}
		}
	}()

	if err := app.run(ctx); err != nil {
		log.Fatal(err)
	}
}

func (app *application) run(ctx context.Context) error {

	app.logger.Infof("Starting...")
	app.logger.Infof("BUILD: %v", BUILD)
	app.logger.Infof("COMMIT: %s", COMMIT[len(COMMIT)-7:])

	metricsListener := fmt.Sprintf(":%v", app.metricsPort)

	app.logger.Infof("Metrics listening on %s", metricsListener)
	app.startMetrics(metricsListener)

	// ctx, cancel := context.WithCancel(context.Background())

	// Create audio directory if it does not exist
	if _, err := os.Stat("audio"); os.IsNotExist(err) {
		err = os.Mkdir("audio", 0755)

		if err != nil {
			app.logger.Errorf("Could not create audio directory: %v", err)
		}
	}

	for {

		app.logger.Infof("Starting HTTP server on :%s", app.serviceHTTPPort)

		hub := websocket.NewHub(app.logger)
		go hub.Run()

		app.SendChan = hub.Send

		httpService := httpservice.NewHttpService(app.permissiveHeaders, app.logger, app.counters)

		httpService.SendChan = app.SendChan

		router := httpService.NewRouter(hub)

		httpServer, httpErrorChan := app.startHTTPServer(app.serviceHTTPPort, router)

		httpShutdownCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		defer httpServer.Shutdown(httpShutdownCtx)

		// Wait for SIGINT and SIGTERM (HIT CTRL-C)

		select {
		case <-ctx.Done():
			return nil
		case err := <-httpErrorChan:
			app.logger.Errorf("HTTP Server error: %v", err)
			time.Sleep(5 * time.Second)
		}
	}

}

func (app *application) startHTTPServer(listener string, handler http.Handler) (*http.Server, <-chan error) {

	logger, err := zap.NewStdLogAt(app.logger.Desugar(), zap.ErrorLevel)

	if err != nil {
		app.logger.Errorf("could not create standard error logger from sugared logger: %v", err)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.serviceHTTPPort),
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
		ErrorLog:     logger,
	}

	errorChan := make(chan error)

	go func() {

		if err := srv.ListenAndServe(); err != nil {
			errorChan <- fmt.Errorf("http server error: %v", err)

		}
	}()

	// returning reference so caller can call Shutdown()
	return srv, errorChan
}

func (app *application) setupLogging() {

	// The bundled Config struct only supports the most common configuration
	// options. More complex needs, like splitting logs between multiple files
	// or writing to non-file outputs, require use of the zapcore package.
	//
	// In this example, imagine we're both sending our logs to Kafka and writing
	// them to the console. We'd like to encode the console output and the Kafka
	// topics differently, and we'd also like special treatment for
	// high-priority logs.

	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	standardOut := zapcore.Lock(os.Stdout)

	consoleEncoder := func() zapcore.Encoder {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

		if viper.GetBool("server.console_logs") {
			return zapcore.NewConsoleEncoder(encoderConfig)
		}

		return zapcore.NewJSONEncoder(encoderConfig)
	}()

	// consoleEncoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, standardOut, highPriority),
		zapcore.NewCore(consoleEncoder, standardOut, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	app.logger = logger.Sugar()

}
