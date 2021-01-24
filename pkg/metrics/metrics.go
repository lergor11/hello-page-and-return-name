package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	Request = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request",
			Help: "How many requests to server.",
		},
		[]string{"method"},
	)
)

func init() {
	prometheus.MustRegister(Request)

}
