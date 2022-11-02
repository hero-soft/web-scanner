package httpservice

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.uber.org/zap"
)

type HttpService struct {
	baseUrl           string
	permissiveHeaders bool
	logger            *zap.SugaredLogger
	counters          map[string]prometheus.Counter
}

func NewResponderService(baseUrl string, permissiveHeaders bool, logger *zap.SugaredLogger, counters map[string]prometheus.Counter) *HttpService {

	counters["not_found_hits"] = newHitCounter("not found")
	counters["method_not_allowed_hits"] = newHitCounter("method not allowed")
	counters["index_hits"] = newHitCounter("index")
	counters["health_hits"] = newHitCounter("health")

	counters["directory_hits"] = newHitCounter("directory")
	counters["issuing_certificate_hits"] = newHitCounter("issuing certificate")
	counters["account_hits"] = newHitCounter("new account")
	counters["nonce_hits"] = newHitCounter("new_nonce ")
	counters["new_order_hits"] = newHitCounter("new order")
	counters["revoke_cert_hits"] = newHitCounter("revoke cert")
	counters["key_change_hits"] = newHitCounter("key change")
	counters["authz_hits"] = newHitCounter("authz")
	counters["challenge_hits"] = newHitCounter("challenge")
	counters["finalize_order_hits"] = newHitCounter("order finalize")
	counters["get_order_hits"] = newHitCounter("get order")
	counters["post_order_hits"] = newHitCounter("post order")
	counters["cert_hits"] = newHitCounter("cert")

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
