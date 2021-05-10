package hstore

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyMap    = errors.New("cannot create the query of hstore data type because the map provided in parameters is empty")
	ErrTagsParsing = errors.New("cannot parse tags stored as hstore")
)

func getNewQuery(query string, column string, key string, value string) string {
	queryBase := fmt.Sprintf("%s->'%s' = '%s'", column, key, value)
	if query == "(" {
		query += queryBase
	} else {
		query += fmt.Sprintf(" OR %s", queryBase)
	}
	return query
}

// ConditionalQuery transforms a map into a conditional query by recursion of hstore.
func ConditionalQuery(column string, m map[string][]string) (string, error) {
	if len(m) == 0 {
		return "", ErrEmptyMap
	}
	query := "("

	for key, values := range m {
		for _, v := range values {
			query = getNewQuery(query, column, key, v)
		}
	}
	query += ")"
	return query, nil
}

func cleanSet(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"")
	return s
}

// ToMap transforms hstore to a map.
func ToMap(tags string) (map[string]string, error) {
	res := make(map[string]string)
	values := strings.Split(tags, "\",")

	for _, v := range values {
		pairs := strings.Split(v, "=>")
		if len(pairs) != 2 {
			return nil, ErrTagsParsing
		}
		key := cleanSet(pairs[0])
		value := cleanSet(pairs[1])
		res[key] = value
	}
	return res, nil
}
