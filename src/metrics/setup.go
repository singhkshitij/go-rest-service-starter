package metrics

import (
	"github.com/cyberdelia/go-metrics-graphite"
	"github.com/rcrowley/go-metrics"
	"github.com/singhkshitij/golang-rest-service-starter/src/config"
	"github.com/singhkshitij/golang-rest-service-starter/src/logger"
	"log"
	"net"
	"os"
	"time"
)

var UserAPICount metrics.Counter

func Setup() {
	UserAPICount = metrics.NewCounter()
	metrics.Register("user-api-count", UserAPICount)

	emitMetrics()
}

func emitMetrics() {
	if config.GetEnv() == "prod" {
		addr, err := net.ResolveTCPAddr("tcp", "34.117.7.29:80")
		if err != nil {
			logger.Error("Failed to resolve metrics datasource ", logger.KV("error", err))
		}
		go graphite.Graphite(
			metrics.DefaultRegistry,
			5*time.Second,
			"metrics: ", addr)
	} else {
		go metrics.Log(metrics.DefaultRegistry,
			2*time.Minute,
			log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
	}
}

func Increment(counter metrics.Counter) {
	counter.Inc(1)
}
