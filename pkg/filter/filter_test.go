package filter

import (
	"testing"
)

func TestFilter(t *testing.T) {
	// dynamicFilterJson := `[{"movieYear": {"eq": 2014}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`
	// dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)
	// filteredMovies := dynamicFilter.Filter(slices.Values(testMovies))
	// filteredMovieValues := slices.Collect(filteredMovies)

	// assert.Len(t, filteredMovieValues, 2)
	// assert.Equal(t, testMovies[0], filteredMovieValues[0])
	// assert.Equal(t, testMovies[2], filteredMovieValues[1])
}

func TestFilterWithPredefinedType(t *testing.T) {
	// dynamicFilterJson := `[{"movieYear": {"eq": 2014}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`
	// predefinedFilter := []MovieFilter{}
	// jsonx.MustUnmarshalFromString(dynamicFilterJson, &predefinedFilter)
	// dynamicFilter := DynamicFilter[Movie]{predefinedFilter}
	// filteredMovies := dynamicFilter.FilterAll(testMovies)

	// assert.Len(t, filteredMovies, 2)
	// assert.Equal(t, testMovies[0], filteredMovies[0])
	// assert.Equal(t, testMovies[2], filteredMovies[1])
}
