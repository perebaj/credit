package bureaus

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

// This test just validate that the metrics are being registered correctly.
// and all labels are being set.
func TestMetrics(t *testing.T) {
	registry := prometheus.NewRegistry()
	MustRegistryBureauMetrics(registry)

	bureauCounter.WithLabelValues("200", "receita_federal").Inc()
	bureauDuration.WithLabelValues("200", "receita_federal").Observe(0.1)

	m, err := registry.Gather()
	if err != nil {
		t.Fatalf("error gathering metrics: %v", err)
	}

	bureauCounterOK := false
	bureauDurationOK := false
	for _, metric := range m {
		if metric.GetName() == "bureau_requests_total" {
			bureauCounterOK = true
		}

		if metric.GetName() == "bureau_request_duration_seconds" {
			bureauDurationOK = true
		}
	}

	if !bureauCounterOK {
		t.Errorf("bureau_requests_total metric not found")
	}

	if !bureauDurationOK {
		t.Errorf("bureau_request_duration_seconds metric not found")
	}
}
