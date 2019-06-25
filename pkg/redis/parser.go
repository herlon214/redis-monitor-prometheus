package redis

import (
	"fmt"
	"strings"
)

// ExtractQueryFromLine get the monitor output and extracts the query
func ExtractQueryFromLine(line string) (string, error) {
	querySlice := strings.Split(line, " ")

	if len(querySlice) < 3 {
		return "", fmt.Errorf("Not a query line: %s", line)
	}

	querySlice = querySlice[3:]
	query := strings.Join(querySlice, " ")
	query = strings.Replace(query[:len(query)-1], "\"", "", -1)

	return query, nil
}
