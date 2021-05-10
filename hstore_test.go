package hstore_test

import (
	"github.com/SafeBike/hstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditionalQuery(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		m := map[string][]string{
			"amenity": {"restaurant", "bench"},
		}

		expectedQuery := "(tags->'amenity' = 'restaurant' OR tags->'amenity' = 'bench')"
		columnName := "tags"
		query, err := hstore.ConditionalQuery(columnName, m)
		assert.NoError(t, err)
		assert.Equal(t, expectedQuery, query)
	})

	t.Run("success", func(t *testing.T) {
		m := map[string][]string{
			"amenity": {"restaurant", "bench"},
			"name":    {"some name"},
		}

		expectedQuery := "(tags->'amenity' = 'restaurant' OR tags->'amenity' = 'bench' OR tags->'name' = 'some name')"
		expectedQuery2 := "(tags->'name' = 'some name' OR tags->'amenity' = 'restaurant' OR tags->'amenity' = 'bench')"
		columnName := "tags"
		query, err := hstore.ConditionalQuery(columnName, m)
		assert.NoError(t, err)
		assert.Condition(t, func() (success bool) {
			return query == expectedQuery || query == expectedQuery2
		})
	})

	t.Run("error", func(t *testing.T) {
		m := make(map[string][]string)

		columnName := "tags"
		query, err := hstore.ConditionalQuery(columnName, m)
		assert.Error(t, err, hstore.ErrEmptyMap)
		assert.Empty(t, query)
	})
}

func TestToMap(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tags := `"seats"=>"5", "amenity"=>"bench", "backrest"=>"yes", "material"=>"wood"`
		expectedRes := map[string]string{
			"seats": "5",
			"amenity": "bench",
			"backrest": "yes",
			"material": "wood",
		}
		res, err := hstore.ToMap(tags)
		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})
}