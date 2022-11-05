package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (app *application) startMetrics(listener string) (*http.Server, <-chan error) {

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:    listener,
		Handler: mux,
		//TLSConfig:    tlsConfig,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
		//ErrorLog:     app.rootLogger.GetErrorLogger(),
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
