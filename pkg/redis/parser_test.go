package redis

import (
	"testing"

	"github.com/herlon214/redis-monitor-prometheus/pkg/redis"
)

func TestExtractQueryFromLine(t *testing.T) {
	line := `1561384669.058000 [0 10.2.30.157:57984] "hgetall" "my_key"`
	query := redis.ExtractQueryFromLine(line)

	if query != "hgetall my_key" {
		t.Errorf("Invalid extracteed query output: %s", query)
	}
}
