package redis

import "strings"

// ExtractQueryFromLine get the monitor output and extracts the query
func ExtractQueryFromLine(line string) string {
	querySlice := strings.Split(line, " ")
	querySlice = querySlice[3:]
	query := strings.Join(querySlice, " ")
	query = strings.Replace(query[:len(query)-1], "\"", "", -1)

	return query
}
