package httpservice

import (
	"fmt"
	"strings"

	"github.com/hero-soft/web-scanner/pkg/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
)

type HttpService struct {
	baseUrl           string
	permissiveHeaders bool
	logger            *zap.SugaredLogger
	counters          map[string]prometheus.Counter
	SendChan          chan<- websocket.SendTo
}

func NewResponderService(baseUrl string, permissiveHeaders bool, logger *zap.SugaredLogger, counters map[string]prometheus.Counter) *HttpService {

	counters["not_found_hits"] = newHitCounter("not found")
	counters["method_not_allowed_hits"] = newHitCounter("method not allowed")
	counters["index_hits"] = newHitCounter("index")
	counters["health_hits"] = newHitCounter("health")

	counters["directory_hits"] = newHitCounter("directory")

	return &HttpService{
		baseUrl:           baseUrl,
		permissiveHeaders: permissiveHeaders,
		logger:            logger,
		counters:          counters,
	}
}

func newHitCounter(name string) prometheus.Counter {
	return promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "certmgr",
		Subsystem: "handler",
		Name:      fmt.Sprintf("%s_hits", strings.ReplaceAll(strings.ToLower(name), " ", "_")),
		Help:      fmt.Sprintf("Number of hits for handler %s", name),
	})
}
