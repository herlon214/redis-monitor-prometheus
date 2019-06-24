package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/herlon214/redis-monitor-prometheus/pkg/prometheus"
	"github.com/herlon214/redis-monitor-prometheus/pkg/redis"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// PORT is the port that will be used by the webserver
	PORT = os.Getenv("PORT")
)

func main() {
	log.Println("[ Redis Monitor Prometheus ]")
	promWriter := &prometheus.Writer{}
	watcher := &redis.Watcher{
		Writer: promWriter,
	}

	log.Println("-> Running watcher...")
	go watcher.Run()

	// Starts a web server and sets the prometheus handler
	log.Println("-> Starting webserver...")

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}
