package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	vvv = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "kube_link",
		Name:      "vvv",
		Help:      "vvv",
	}, []string{"method", "path"})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()

			vvv.With(prometheus.Labels{
				"method": "get",
				"path":   "/m",
			}).Add(12)
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

}
