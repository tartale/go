package filter

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/jsonx"
)

type ShowKind string

const (
	MOVIE  ShowKind = "MOVIE"
	SERIES ShowKind = "SERIES"
)

type Movie struct {
	Kind        ShowKind `json:"kind,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	MovieYear   int      `json:"movieYear,omitempty"`
}

var testMovie = Movie{
	Kind:        MOVIE,
	Title:       "Back to the Future",
	Description: "The time travel adventures of Doc Brown and Marty McFly",
	MovieYear:   1985,
}

func TestNewTypeFilterFromJson(t *testing.T) {
	typeFilterJsonIn := `[{"kind":{"eq":"MOVIE"}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJsonIn)
	typeFilterJsonOut := jsonx.MustMarshalToString(typeFilter)

	assert.Equal(t, typeFilterJsonIn, typeFilterJsonOut)
}

func TestNewTypeFilterFromJson_ImpliedLogicalAndFilter(t *testing.T) {
	typeFilterJsonIn := `[{"title":{"matches":"Back to the.*"}},{"movieYear":{"eq":1985}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJsonIn)
	typeFilterJsonOut := jsonx.MustMarshalToString(typeFilter)

	assert.Equal(t, typeFilterJsonIn, typeFilterJsonOut)
}

func TestNewTypeFilterFromJson_ExplicitLogicalAndFilter(t *testing.T) {
	typeFilterJsonIn := `[{"movieYear":{"eq":1985}},{"and":[{"title":{"matches":"Back to the.*"}}]}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJsonIn)
	typeFilterJsonOut := jsonx.MustMarshalToString(typeFilter)

	assert.Equal(t, typeFilterJsonIn, typeFilterJsonOut)
}

func TestNewTypeFilterFromJson_ComplexNestedFilter(t *testing.T) {
	typeFilterJsonIn := `[{"movieYear":{"eq":1955}},{"or":[{"movieYear":{"eq":1985}},{"and":[{"title":{"eq":"BacktotheFuture"}}]}]}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJsonIn)
	typeFilterJsonOut := jsonx.MustMarshalToString(typeFilter)

	assert.Equal(t, typeFilterJsonIn, typeFilterJsonOut)
}

func TestShouldInclude_SimpleEnumFilter(t *testing.T) {
	typeFilterJson := `[{"kind": {"eq": "MOVIE"}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_SimpleStringFilter(t *testing.T) {
	typeFilterJson := `[{"title": {"eq": "Back to the Future"}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_SimpleNumberFilter(t *testing.T) {
	typeFilterJson := `[{"movieYear": {"eq": 1985}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ImpliedLogicalAndFilter(t *testing.T) {
	typeFilterJson := `[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1985}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ExplicitLogicalAndFilter(t *testing.T) {
	typeFilterJson := `[{"movieYear": {"eq": 1985}}, {"and": [{"title": {"matches": "Back to the .*"}}]}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ExplicitLogicalOrFilter(t *testing.T) {
	typeFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "Back to the .*"}}]}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ComplexNestedFilter(t *testing.T) {
	typeFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)
	assert.True(t, result)
}

func TestShouldNotInclude_SimpleEnumFilter(t *testing.T) {
	typeFilterJson := `[{"kind": {"eq": "SERIES"}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_SimpleStringFilter(t *testing.T) {
	typeFilterJson := `[{"title": {"eq": "The Shawshank Redemption"}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_SimpleNumberFilter(t *testing.T) {
	typeFilterJson := `[{"movieYear": {"eq": 1955}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ImpliedLogicalAndFilter(t *testing.T) {
	typeFilterJson := `[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1955}}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ExplicitLogicalAndFilter(t *testing.T) {
	typeFilterJson := `[{"movieYear": {"eq": 1955}}, {"and": [{"title": {"matches": "Back to the .*"}}]}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ExplicitLogicalOrFilter(t *testing.T) {
	typeFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "The Shawshank .*"}}]}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ComplexNestedFilter(t *testing.T) {
	typeFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "The Shawshank Redemption"}}]}]}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	result := typeFilter.ShouldInclude(testMovie)
	assert.False(t, result)
}

func TestFilter(t *testing.T) {
	typeFilterJson := `[{"movieYear": {"eq": 2014}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`
	typeFilter := NewTypeFilterFromJson[Movie](typeFilterJson)

	testMovies := []Movie{
		testMovie,
		{
			Kind:        MOVIE,
			Title:       "The Shawshank Redemption",
			Description: "Andy DuFresne escapes from prison.",
			MovieYear:   1995,
		},
		{
			Kind:        MOVIE,
			Title:       "Interstellar",
			Description: "Matt Damon is the bad guy.",
			MovieYear:   2014,
		},
	}

	filteredMovies := typeFilter.Filter(slices.Values(testMovies))
	filteredMovieValues := slices.Collect(filteredMovies)

	assert.Len(t, filteredMovieValues, 2)
	assert.Equal(t, testMovies[0], filteredMovieValues[0])
	assert.Equal(t, testMovies[2], filteredMovieValues[1])
}
