package metrics

import (
	"github.com/penglongli/gin-metrics/ginmetrics"
)

var monitor *ginmetrics.Monitor
var UserMetric *ginmetrics.Metric

func Setup() {

	//SET GIN Metrics observability
	// get global Monitor object
	monitor = ginmetrics.GetMonitor()
	monitor.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	monitor.SetSlowTime(2)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	monitor.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	initialiseMetrics()
}

func GetMetricsMonitor() *ginmetrics.Monitor {
	return monitor
}

func initialiseMetrics(){
	UserMetric = &ginmetrics.Metric{
		Type:        ginmetrics.Counter,
		Name:        "user_api_count",
		Description: "an example of counter type metric",
		Labels:      []string{"user"},
	}

	// Add metric to global monitor object
	_ = ginmetrics.GetMonitor().AddMetric(UserMetric)
}

func Increment(metric string, labels []string) {
	ginmetrics.GetMonitor().GetMetric(metric).Inc(labels)
}
