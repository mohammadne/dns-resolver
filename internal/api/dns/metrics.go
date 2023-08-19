package dns

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	requestsTotal      *prometheus.CounterVec
	requestsInProgress *prometheus.GaugeVec
	requestsDuration   *prometheus.HistogramVec
}

func newMetrics() *metrics {
	constLabels := prometheus.Labels{
		"api": "dns",
	}

	totalRequests := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:        "total_requests",
		Help:        "This is counter for total requests from clients.",
		ConstLabels: constLabels,
	}, []string{})

	requestsInProgress := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:        "requests_in_progress",
		Help:        "All the requests in progress.",
		ConstLabels: constLabels,
	}, []string{})

	requestsDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        "request_duration_seconds",
		Help:        "Duration of all DNS requests by remote.",
		ConstLabels: constLabels,
	}, []string{})

	// registering metric collectors
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(requestsInProgress)
	prometheus.MustRegister(requestsDuration)

	return &metrics{
		requestsTotal:      totalRequests,
		requestsInProgress: requestsInProgress,
		requestsDuration:   requestsDuration,
	}
}
