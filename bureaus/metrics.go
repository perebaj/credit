// Package bureaus metrics.go provides a metric client to collect metrics from the external bureaus that
// we are integrating with.
package bureaus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

// svcName is used to identify the service in the metrics.
// All metrics will be tagged with this value.
const svcName = "credit-api"

// MustRegistryBureauMetrics registers the metrics for the bureau client.
//
// New metrics should be added here.
func MustRegistryBureauMetrics(registry prometheus.Registerer) {
	prometheus.WrapRegistererWith(prometheus.Labels{"serviceName": svcName}, registry).MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		bureauCounter,
		bureauDuration,
	)
}

var (
	bureauCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "bureau_requests_total",
		Help: "Total of bureau requests",
	}, []string{"status_code", "bureau_name"})

	bureauDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "bureau_request_duration_seconds",
		Help:    "Duration of bureau requests",
		Buckets: []float64{.1, .2, .3, .4, .5, 1, 2, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60},
	}, []string{"status_code", "bureau_name"})
)
