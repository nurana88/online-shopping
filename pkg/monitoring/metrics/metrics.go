package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var RegisterRequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "register_request_counter",
		Help: "Number of request handled by register handler",
	},
)
