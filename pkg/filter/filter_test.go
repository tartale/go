package filter

import (
	"encoding/json"

	"github.com/PaesslerAG/gval"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tartale/go/pkg/jsonx"
)

type Movie struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	MovieYear   int    `json:"movieYear"`
}

type MovieFilter struct {
	Title       *Operator      `json:"title,omitempty"`
	Description *Operator      `json:"description,omitempty"`
	MovieYear   *Operator      `json:"movieYear,omitempty"`
	And         []*MovieFilter `json:"and,omitempty"`
	Or          []*MovieFilter `json:"or,omitempty"`
}

var _ = Describe("Filtering", func() {

	Context("for syntactically correct filters", func() {
		movie := Movie{
			Title:       "Back to the Future",
			Description: "The time travel adventures of Doc Brown and Marty McFly",
			MovieYear:   1985,
		}

		DescribeTable("can be evaluated against an input that should return true",
			func(movieFiltersJson string, movie Movie) {

				var movieFilters []*MovieFilter
				err := jsonx.StrictUnmarshal([]byte(movieFiltersJson), &movieFilters)
				Expect(err).ToNot(HaveOccurred())

				expression := GetExpression(movieFilters)
				values := GetValues(movieFilters, movie)
				eval, err := gval.Evaluate(expression, values)

				Expect(err).ToNot(HaveOccurred())
				Expect(eval.(bool)).To(Equal(true))
			},

			Entry("simple filter",
				`[{"title": {"eq": "Back to the Future"}}]`,
				movie,
			),

			Entry("multi-field filter with implied logical 'and'",
				`[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1985}}]`,
				movie,
			),

			Entry("multi-field filter with explicit logical 'or'",
				`[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "Back to the .*"}}]}]`,
				movie,
			),

			Entry("multi-field nested filter",
				`[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`,
				movie,
			),
		)

		DescribeTable("can be evaluated against an input that should return false",
			func(movieFiltersJson string, movie Movie) {

				var movieFilters []*MovieFilter
				err := json.Unmarshal([]byte(movieFiltersJson), &movieFilters)
				Expect(err).ToNot(HaveOccurred())

				expression := GetExpression(movieFilters)
				values := GetValues(movieFilters, movie)
				eval, err := gval.Evaluate(expression, values)

				Expect(err).ToNot(HaveOccurred())
				Expect(eval.(bool)).To(Equal(false))
			},

			Entry("simple filter",
				`[{"title": {"eq": "The Shawshank Redemption"}}]`,
				movie,
			),

			Entry("multi-field filter with implied logical 'and'",
				`[{"movieYear": {"eq": 1955}}, {"title": {"matches": "Back to the .*"}}]`,
				movie,
			),

			Entry("multi-field filter with explicit logical 'or'",
				`[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": ".*Shawshank.*"}}]}]`,
				movie,
			),

			Entry("multi-field nested filter",
				`[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "The Shawshank Redemption"}}]}]}]`,
				movie,
			),
		)
	})
})
