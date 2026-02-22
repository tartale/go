package filter

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/jsonx"
)

type ShowKind string

const (
	MOVIE  ShowKind = "MOVIE"
	SERIES ShowKind = "SERIES"
)

type Movie struct {
	Kind        ShowKind `json:"kind"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	MovieYear   int      `json:"movieYear"`
}

type MovieFilter struct {
	Kind        *Operator      `json:"kind,omitempty"`
	Title       *Operator      `json:"title,omitempty"`
	Description *Operator      `json:"description,omitempty"`
	MovieYear   *Operator      `json:"movieYear,omitempty"`
	And         []*MovieFilter `json:"and,omitempty"`
	Or          []*MovieFilter `json:"or,omitempty"`
}

func TestNewTypeFilter(t *testing.T) {
	typeFilterJsonIn := `{"kind":{"eq":"MOVIE"}}`
	typeFilter := NewTypeFilter[Movie](typeFilterJsonIn)
	typeFilterJsonOut := jsonx.MustMarshalToString(typeFilter)

	assert.Equal(t, typeFilterJsonIn, typeFilterJsonOut)
}

func TestShouldInclude(t *testing.T) {
	typeFilterJson := `{"kind": {"eq": "MOVIE"}}`
	typeFilter := NewTypeFilter[Movie](typeFilterJson)
	jsonx.MustUnmarshalFromString(typeFilterJson, &typeFilter)
	movie := Movie{
		Kind:        MOVIE,
		Title:       "Back to the Future",
		Description: "The time travel adventures of Doc Brown and Marty McFly",
		MovieYear:   1985,
	}

	result := typeFilter.ShouldInclude(movie)

	assert.True(t, result)
}

var _ = Describe("Filtering", func() {
	Context("for syntactically correct filters", func() {
		movie := Movie{
			Kind:        MOVIE,
			Title:       "Back to the Future",
			Description: "The time travel adventures of Doc Brown and Marty McFly",
			MovieYear:   1985,
		}

		DescribeTable("can be evaluated against an input that should return true",
			func(movieFiltersJson string, movie Movie) {
				var movieFilters []*MovieFilter
				err := jsonx.StrictUnmarshal([]byte(movieFiltersJson), &movieFilters)
				Expect(err).ToNot(HaveOccurred())
			},

			Entry("simple enum filter",
				`[{"kind": {"eq": "MOVIE"}}]`,
				movie,
			),
		)
	})
})

////////////// legacy //////////////

// var _ = Describe("Filtering", func() {
// 	Context("for syntactically correct filters", func() {
// 		movie := Movie{
// 			Kind:        MOVIE,
// 			Title:       "Back to the Future",
// 			Description: "The time travel adventures of Doc Brown and Marty McFly",
// 			MovieYear:   1985,
// 		}

// 		DescribeTable("can be evaluated against an input that should return true",
// 			func(movieFiltersJson string, movie Movie) {
// 				var movieFilters []*MovieFilter
// 				err := jsonx.StrictUnmarshal([]byte(movieFiltersJson), &movieFilters)
// 				Expect(err).ToNot(HaveOccurred())

// 				expression := GetExpression(movieFilters)
// 				values := GetValues(movieFilters, movie)
// 				eval, err := gval.Evaluate(expression, values)

// 				Expect(err).ToNot(HaveOccurred())
// 				Expect(eval.(bool)).To(Equal(true))
// 			},

// 			Entry("simple enum filter",
// 				`[{"kind": {"eq": "MOVIE"}}]`,
// 				movie,
// 			),

// 			Entry("simple string filter",
// 				`[{"title": {"eq": "Back to the Future"}}]`,
// 				movie,
// 			),

// 			Entry("simple number filter",
// 				`[{"movieYear": {"eq": 1985}}]`,
// 				movie,
// 			),

// 			Entry("multi-field filter with implied logical 'and'",
// 				`[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1985}}]`,
// 				movie,
// 			),

// 			Entry("multi-field filter with explicit logical 'or'",
// 				`[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "Back to the .*"}}]}]`,
// 				movie,
// 			),

// 			Entry("multi-field nested filter",
// 				`[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`,
// 				movie,
// 			),
// 		)

// 		DescribeTable("can be evaluated against an input that should return false",
// 			func(movieFiltersJson string, movie Movie) {
// 				var movieFilters []*MovieFilter
// 				err := json.Unmarshal([]byte(movieFiltersJson), &movieFilters)
// 				Expect(err).ToNot(HaveOccurred())

// 				expression := GetExpression(movieFilters)
// 				values := GetValues(movieFilters, movie)
// 				eval, err := gval.Evaluate(expression, values)

// 				Expect(err).ToNot(HaveOccurred())
// 				Expect(eval.(bool)).To(Equal(false))
// 			},

// 			Entry("simple enum filter",
// 				`[{"kind": {"eq": "SERIES"}}]`,
// 				movie,
// 			),

// 			Entry("simple string filter",
// 				`[{"title": {"eq": "The Shawshank Redemption"}}]`,
// 				movie,
// 			),

// 			Entry("simple number filter",
// 				`[{"movieYear": {"eq": 1955}}]`,
// 				movie,
// 			),

// 			Entry("multi-field filter with implied logical 'and'",
// 				`[{"movieYear": {"eq": 1955}}, {"title": {"matches": "Back to the .*"}}]`,
// 				movie,
// 			),

// 			Entry("multi-field filter with explicit logical 'or'",
// 				`[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": ".*Shawshank.*"}}]}]`,
// 				movie,
// 			),

// 			Entry("multi-field nested filter",
// 				`[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "The Shawshank Redemption"}}]}]}]`,
// 				movie,
// 			),
// 		)
// 	})
// })
