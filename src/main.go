package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/herlon214/redis-monitor-prometheus/pkg/prometheus"
	"github.com/herlon214/redis-monitor-prometheus/pkg/redis"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// PORT is the port that will be used by the webserver
	PORT = os.Getenv("PORT")
	// redisURI is used to connect to redis
	// you can put many URI, separated by semicolon
	redisURI = os.Getenv("REDIS_URI")
)

func main() {
	log.Println("[ Redis Monitor Prometheus ]")
	promWriter := &prometheus.Writer{}
	watcher := &redis.Watcher{
		Writer: promWriter,
	}

	// Parse redis uri
	servers := strings.Split(redisURI, ";")

	log.Printf("Found %d servers on the REDIS_URI env...", len(servers))

	// Run a monitor for each server
	for i, server := range servers {
		log.Printf("-> Running watcher [%d]...", i)
		go watcher.Run(server)
	}

	// Starts a web server and sets the prometheus handler
	log.Println("-> Starting webserver...")

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}
