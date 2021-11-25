package redis

import (
	"fmt"
	"strings"
)

// ExtractQueryFromLine get the monitor output and extracts the query
func ExtractQueryFromLine(line string, granularity int) (string, error) {
	querySlice := strings.Split(line, " ")

	if len(querySlice) <= 3 {
		return "", fmt.Errorf("Not a query line: %s", line)
	}

	querySlice = querySlice[3:]
	query := strings.Join(querySlice, " ")
	query = strings.Replace(query[:len(query)-1], "\"", "", -1)

	// Return based on granularity
	if granularity == 0 {
		return query, nil
	}

	granularQuery := strings.Split(query, " ")

	// Avoid access array out of bounds
	if len(granularQuery) < granularity {
		return query, nil
	}

	granularQuery = granularQuery[0:granularity]

	return strings.Join(granularQuery, " "), nil
}
