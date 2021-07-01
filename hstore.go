package hstore

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyMaps   = errors.New("cannot create the query of hstore data type because the maps provided in parameters are empty")
	ErrTagsParsing = errors.New("cannot parse tags stored as hstore")
)

// ConditionalQuery transforms a map into a conditional query by recursion of hstore.
func ConditionalQuery(column string, ms []map[string]string) (string, error) {
	if len(ms) == 0 {
		return "", ErrEmptyMaps
	}
	query := "("

	for i, m := range ms {
		if len(ms) > 1 {
			query += "("
		}
		j := 0
		for key, value := range m {
			query += fmt.Sprintf("%s->'%s' = '%s'", column, key, value)
			if j < len(m) - 1 {
				query += " AND "
			}
			j += 1
		}
		if len(ms) > 1 {
			query += ")"
		}
		if i < len(ms) - 1 {
			query += " OR "
		}
		i += 1
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
