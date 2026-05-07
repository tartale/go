package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/jsonx"
)

func TestNewDynamicFilterFromJson(t *testing.T) {
	dynamicFilterJsonIn := `[{"kind":{"eq":"MOVIE"}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJsonIn)
	dynamicFilterJsonOut := jsonx.MustMarshalToString(dynamicFilter)

	assert.Equal(t, dynamicFilterJsonIn, dynamicFilterJsonOut)
}

func TestNewDynamicFilterFromJson_ImpliedLogicalAndFilter(t *testing.T) {
	dynamicFilterJsonIn := `[{"title":{"matches":"Back to the.*"}},{"movieYear":{"eq":1985}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJsonIn)
	dynamicFilterJsonOut := jsonx.MustMarshalToString(dynamicFilter)

	assert.Equal(t, dynamicFilterJsonIn, dynamicFilterJsonOut)
}

func TestNewDynamicFilterFromJson_ExplicitLogicalAndFilter(t *testing.T) {
	dynamicFilterJsonIn := `[{"movieYear":{"eq":1985}},{"and":[{"title":{"matches":"Back to the.*"}}]}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJsonIn)
	dynamicFilterJsonOut := jsonx.MustMarshalToString(dynamicFilter)

	assert.Equal(t, dynamicFilterJsonIn, dynamicFilterJsonOut)
}

func TestNewDynamicFilterFromJson_ComplexNestedFilter(t *testing.T) {
	dynamicFilterJsonIn := `[{"movieYear":{"eq":1955}},{"or":[{"movieYear":{"eq":1985}},{"and":[{"title":{"eq":"BacktotheFuture"}}]}]}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJsonIn)
	dynamicFilterJsonOut := jsonx.MustMarshalToString(dynamicFilter)

	assert.Equal(t, dynamicFilterJsonIn, dynamicFilterJsonOut)
}

func TestShouldInclude_SimpleEnumFilter(t *testing.T) {
	dynamicFilterJson := `[{"kind": {"eq": "MOVIE"}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_SimpleStringFilter(t *testing.T) {
	dynamicFilterJson := `[{"title": {"eq": "Back to the Future"}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_SimpleNumberFilter(t *testing.T) {
	dynamicFilterJson := `[{"movieYear": {"eq": 1985}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ImpliedLogicalAndFilter(t *testing.T) {
	dynamicFilterJson := `[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1985}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ExplicitLogicalAndFilter(t *testing.T) {
	dynamicFilterJson := `[{"movieYear": {"eq": 1985}}, {"and": [{"title": {"matches": "Back to the .*"}}]}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ExplicitLogicalOrFilter(t *testing.T) {
	dynamicFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "Back to the .*"}}]}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.True(t, result)
}

func TestShouldInclude_ComplexNestedFilter(t *testing.T) {
	dynamicFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)
	assert.True(t, result)
}

func TestShouldNotInclude_SimpleEnumFilter(t *testing.T) {
	dynamicFilterJson := `[{"kind": {"eq": "SERIES"}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_SimpleStringFilter(t *testing.T) {
	dynamicFilterJson := `[{"title": {"eq": "The Shawshank Redemption"}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_SimpleNumberFilter(t *testing.T) {
	dynamicFilterJson := `[{"movieYear": {"eq": 1955}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ImpliedLogicalAndFilter(t *testing.T) {
	dynamicFilterJson := `[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1955}}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ExplicitLogicalAndFilter(t *testing.T) {
	dynamicFilterJson := `[{"movieYear": {"eq": 1955}}, {"and": [{"title": {"matches": "Back to the .*"}}]}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ExplicitLogicalOrFilter(t *testing.T) {
	dynamicFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "The Shawshank .*"}}]}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)

	assert.False(t, result)
}

func TestShouldNotInclude_ComplexNestedFilter(t *testing.T) {
	dynamicFilterJson := `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "The Shawshank Redemption"}}]}]}]`
	dynamicFilter := NewDynamicFilterFromJson[Movie](dynamicFilterJson)

	result := dynamicFilter.ShouldInclude(testMovie)
	assert.False(t, result)
}
