package redis

import (
	"testing"

	"github.com/herlon214/redis-monitor-prometheus/pkg/redis"
)

func TestExtractQueryFromLine(t *testing.T) {
	line := `1561384669.058000 [0 10.2.30.157:57984] "hgetall" "my_key"`
	query, _ := redis.ExtractQueryFromLine(line, 0)

	if query != "hgetall my_key" {
		t.Errorf("Invalid extracted query output: %s", query)
	}
}

func TestExtractQueryFromLineGranularity1(t *testing.T) {
	line := `1561384669.058000 [0 10.2.30.157:57984] "hgetall" "my_key"`
	query, _ := redis.ExtractQueryFromLine(line, 1)

	if query != "hgetall" {
		t.Errorf("Invalid extracted query output: %s", query)
	}
}

func TestExtractQueryFromLineGranularity2(t *testing.T) {
	line := `1561384669.058000 [0 10.2.30.157:57984] "hgetall" "my_key"`
	query, _ := redis.ExtractQueryFromLine(line, 2)

	if query != "hgetall my_key" {
		t.Errorf("Invalid extracted query output: %s", query)
	}
}

func TestExtractQueryFromLineWithTwoSpaces(t *testing.T) {
	line := "one two three"
	// was panic(slice bounds out of range) if splitted slice length == 3
	_, _ = redis.ExtractQueryFromLine(line, 2)
}
