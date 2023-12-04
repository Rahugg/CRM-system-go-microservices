package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	statusError = "error"
	statusOk    = "ok"
)

func mustRegister(collectors ...prometheus.Collector) {
	prometheus.DefaultRegisterer.MustRegister(collectors...)
}

func newHistogramVec(name, help string, buckets []float64, labelValues ...string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "crm_core",
			Name:      name,
			Help:      help,
			Buckets:   buckets,
		},
		labelValues,
	)
}

func newCounterVec(name, help string, labelValues ...string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "crm_core",
			Name:      name,
			Help:      help,
		},
		labelValues,
	)
}
