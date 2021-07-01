package hstore_test

import (
	"github.com/SafeBike/hstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditionalQuery(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		m := []map[string]string{
			{
				"amenity": "restaurant",
			},
		}
		expectedQuery := "(tags->'amenity' = 'restaurant')"
		columnName := "tags"
		query, err := hstore.ConditionalQuery(columnName, m)
		assert.NoError(t, err)
		assert.Equal(t, expectedQuery, query)
	})

	t.Run("success", func(t *testing.T) {
		m := []map[string]string{
			{
				"amenity": "restaurant",
			},
			{
				"amenity": "bench",
			},
		}
		expectedQuery := "((tags->'amenity' = 'restaurant') OR (tags->'amenity' = 'bench'))"
		columnName := "tags"
		query, err := hstore.ConditionalQuery(columnName, m)
		assert.NoError(t, err)
		assert.Equal(t, expectedQuery, query)
	})

	t.Run("success", func(t *testing.T) {
		m := []map[string]string{
			{
				"amenity":    "restaurant",
				"diet:vegan": "yes",
			},
			{
				"amenity":  "bench",
				"material": "wood",
			},
		}
		expectedQuery := "((tags->'amenity' = 'restaurant' AND tags->'diet:vegan' = 'yes') OR (tags->'amenity' = 'bench' AND tags->'material' = 'wood'))"
		expectedQuery2 := "((tags->'diet:vegan' = 'yes' AND tags->'amenity' = 'restaurant') OR (tags->'material' = 'wood' AND tags->'amenity' = 'bench'))"
		expectedQuery3 := "((tags->'diet:vegan' = 'yes' AND tags->'amenity' = 'restaurant') OR (tags->'amenity' = 'bench' AND tags->'material' = 'wood'))"
		expectedQuery4 := "((tags->'amenity' = 'restaurant' AND tags->'diet:vegan' = 'yes') OR (tags->'material' = 'wood' AND tags->'amenity' = 'bench'))"
		columnName := "tags"
		query, err := hstore.ConditionalQuery(columnName, m)
		assert.NoError(t, err)
		assert.Condition(t, func() (success bool) {
			return query == expectedQuery || query == expectedQuery2 || query == expectedQuery3 || query == expectedQuery4
		})
	})

	t.Run("error", func(t *testing.T) {
		m := make([]map[string]string, 0)

		columnName := "tags"
		query, err := hstore.ConditionalQuery(columnName, m)
		assert.Error(t, err, hstore.ErrEmptyMaps)
		assert.Empty(t, query)
	})
}

func TestToMap(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tags := `"seats"=>"5", "amenity"=>"bench", "backrest"=>"yes", "material"=>"wood"`
		expectedRes := map[string]string{
			"seats":    "5",
			"amenity":  "bench",
			"backrest": "yes",
			"material": "wood",
		}
		res, err := hstore.ToMap(tags)
		assert.NoError(t, err)
		assert.Equal(t, expectedRes, res)
	})
}
