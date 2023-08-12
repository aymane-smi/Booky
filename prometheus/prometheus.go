package prometheus_book

import "github.com/prometheus/client_golang/prometheus"

var TotalRequest = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "book_request_count",
		Help: "No of request handled by the handler",
	},
)

var TotalErros = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "book_error_count",
		Help: "No of errors that raised during execution",
	},
	[]string{"type", "method", "package"},
)