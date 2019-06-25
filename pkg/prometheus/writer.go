package prometheus

import (
	"log"
	"strings"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/herlon214/redis-monitor-prometheus/pkg/redis"
)

var (
	processedQueriesCounter = promauto.NewCounter(prom.CounterOpts{
		Name: "redis_processed_queries_total",
		Help: "The total number of processed events",
	})
	queryCounter = promauto.NewCounterVec(prom.CounterOpts{
		Name: "redis_query_executed",
		Help: "Total number that a query was executed",
	}, []string{"query"})
	startScrapping = false // Used to show the scrapping started message
)

// Writer is responsible for write prometheus metrics
type Writer struct{}

// Write writes prometheus metrics from the given line
func (p *Writer) Write(line []byte) (n int, err error) {
	body := strings.TrimSuffix(string(line), "\n")

	// Standard output for redis command:
	// 1561384669.058000 [0 10.2.30.157:57984] "hgetall" "my_key" [...]
	isCommand := strings.Contains(body, "[") && strings.Contains(body, "]")

	// Only do some action if it's a query
	if isCommand {
		// Parse the line to extract only the command part
		query := redis.ExtractQueryFromLine(body)

		// Increase the query execution
		queryCounter.With(prom.Labels{"query": query}).Inc()

		// Increase the processed events
		processedQueriesCounter.Inc()

		if !startScrapping {
			log.Println("-> Started scrapping redis queries...")
			startScrapping = true
		}
	} else {
		log.Printf("[!] Redis not command output: %s", body)
	}

	// Always returning ok
	return len(line), nil
}
