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

	qqq = promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "kube_link",
		Name:      "qqq",
		Help:      "qqq",
	}, func() float64 {
		return 999.9
	})

	ttt = promauto.NewCounterFunc(prometheus.CounterOpts{
		Namespace: "kube_link",
		Name:      "ttt",
		Help:      "ttt",
	}, func() float64 {
		return 999.8
	})
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
	//http.Handle("/metrics", promhttp.Handler())
	//http.ListenAndServe(":2112", nil)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	svc := &http.Server{
		Addr:              ":2113",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	svc.ListenAndServe()
}
