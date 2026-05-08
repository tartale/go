package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/jsonx"
)

func TestNewStructFilter(t *testing.T) {
	structFilterJsonIn := `[{"kind":{"eq":"MOVIE"}}]`
	structFilter := NewStructFilter[Movie](structFilterJsonIn)
	structFilterJsonOut := jsonx.MustMarshalToString(structFilter)

	assert.Equal(t, structFilterJsonIn, structFilterJsonOut)
}

func TestNewStructFilter_ImpliedLogicalAndFilter(t *testing.T) {
	structFilterJsonIn := `[{"title":{"matches":"Back to the.*"}},{"movieYear":{"eq":1985}}]`
	structFilter := NewStructFilter[Movie](structFilterJsonIn)
	structFilterJsonOut := jsonx.MustMarshalToString(structFilter)

	assert.Equal(t, structFilterJsonIn, structFilterJsonOut)
}

func TestNewStructFilter_ExplicitLogicalAndFilter(t *testing.T) {
	structFilterJsonIn := `[{"movieYear":{"eq":1985}},{"and":[{"title":{"matches":"Back to the.*"}}]}]`
	structFilter := NewStructFilter[Movie](structFilterJsonIn)
	structFilterJsonOut := jsonx.MustMarshalToString(structFilter)

	assert.Equal(t, structFilterJsonIn, structFilterJsonOut)
}

func TestNewStructFilter_ComplexNestedFilter(t *testing.T) {
	structFilterJsonIn := `[{"movieYear":{"eq":1955}},{"or":[{"movieYear":{"eq":1985}},{"and":[{"title":{"eq":"BacktotheFuture"}}]}]}]`
	structFilter := NewStructFilter[Movie](structFilterJsonIn)
	structFilterJsonOut := jsonx.MustMarshalToString(structFilter)

	assert.Equal(t, structFilterJsonIn, structFilterJsonOut)
}

func TestShouldInclude_SimpleEnumFilter(t *testing.T) {
	structFilterJson := `[{"kind": {"eq": "MOVIE"}}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_SimpleStringFilter(t *testing.T) {
	structFilterJson := `[{"title": {"eq": "Back to the Future"}}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_SimpleNumberFilter(t *testing.T) {
	structFilterJson := `[{"movieYear": {"eq": 1985}}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ImpliedLogicalAndFilter(t *testing.T) {
	structFilterJson := `[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1985}}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ExplicitLogicalAndFilter(t *testing.T) {
	structFilterJson := `[{"movieYear": {"eq": 1985}}, {"and": [{"title": {"matches": "Back to the .*"}}]}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ExplicitLogicalOrFilter(t *testing.T) {
	structFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "Back to the .*"}}]}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ComplexNestedFilter(t *testing.T) {
	structFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)
	assert.True(t, result)
}

func TestShouldNotInclude_SimpleEnumFilter(t *testing.T) {
	structFilterJson := `[{"kind": {"eq": "SERIES"}}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_SimpleStringFilter(t *testing.T) {
	structFilterJson := `[{"title": {"eq": "The Shawshank Redemption"}}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_SimpleNumberFilter(t *testing.T) {
	structFilterJson := `[{"movieYear": {"eq": 1955}}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ImpliedLogicalAndFilter(t *testing.T) {
	structFilterJson := `[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1955}}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ExplicitLogicalAndFilter(t *testing.T) {
	structFilterJson := `[{"movieYear": {"eq": 1955}}, {"and": [{"title": {"matches": "Back to the .*"}}]}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ExplicitLogicalOrFilter(t *testing.T) {
	structFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "The Shawshank .*"}}]}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ComplexNestedFilter(t *testing.T) {
	structFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "The Shawshank Redemption"}}]}]}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := structFilter.ShouldInclude(testMovie)
	assert.False(t, result)
}

func TestFilterAllFor(t *testing.T) {
	structFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`
	structFilter := NewStructFilter[Movie](structFilterJson)

	result := FilterAllFor(structFilter, testMovieList)

	assert.Len(t, result, 1)
	assert.Equal(t, "Back to the Future", result[0].Title)
}
