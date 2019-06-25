package prometheus

import (
	"log"
	"strings"

	"github.com/herlon214/redis-monitor-prometheus/pkg/redis"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// ProcessedQueriesGauge is the total number of queries executed in prometheus
	ProcessedQueriesGauge = promauto.NewGauge(prom.GaugeOpts{
		Name: "redis_processed_queries_total",
		Help: "The total number of processed events",
	})
	// QueryGauge is the query executed prometheus counter
	QueryGauge = promauto.NewGaugeVec(prom.GaugeOpts{
		Name: "redis_query_executed",
		Help: "Total number that a query was executed",
	}, []string{"query"})
	// StartedScrapping flag used to check if the scrapping started
	StartedScrapping = false // Used to show the scrapping started message
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
		queries := strings.Split(body, "\n")

		for _, queryLine := range queries {
			// Parse the line to extract only the command part
			query, err := redis.ExtractQueryFromLine(queryLine)

			if err != nil {
				log.Printf("Error found: %+v\n", err)
			} else {
				// Increase the query execution
				QueryGauge.With(prom.Labels{"query": query}).Add(1)

				// Increase the processed events
				ProcessedQueriesGauge.Add(1)
			}
		}

		if !StartedScrapping {
			log.Println("-> Started scrapping redis queries...")
			StartedScrapping = true
		}
	} else {
		log.Printf("[!] Redis not command output: %s", body)
	}

	// Always returning ok
	return len(line), nil
}
