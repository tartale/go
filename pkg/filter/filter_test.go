package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/jsonx"
)

type filterTestCase struct {
	name          string
	filterJSON    string
	shouldInclude bool
}

func TestShouldInclude(t *testing.T) {
	movie := Movie{
		Kind:        MOVIE,
		Title:       "Back to the Future",
		Description: "The time travel adventures of Doc Brown and Marty McFly",
		MovieYear:   1985,
	}

	cases := []filterTestCase{
		{
			name:          "simple enum filter should include",
			filterJSON:    `[{"kind": {"eq": "MOVIE"}}]`,
			shouldInclude: true,
		},
		{
			name:          "simple string filter should include",
			filterJSON:    `[{"title": {"eq": "Back to the Future"}}]`,
			shouldInclude: true,
		},
		{
			name:          "simple number filter should include",
			filterJSON:    `[{"movieYear": {"eq": 1985}}]`,
			shouldInclude: true,
		},
		{
			name:          "multi-field filter with implied logical 'and' should include",
			filterJSON:    `[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1985}}]`,
			shouldInclude: true,
		},
		{
			name:          "multi-field filter with explicit logical 'and' should include",
			filterJSON:    `[{"title": {"matches": "Back to the .*"}}, {"and": [{"movieYear": {"eq": 1985}}]}]`,
			shouldInclude: true,
		},
		{
			name:          "multi-field filter with explicit logical 'or' should include",
			filterJSON:    `[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "Back to the .*"}}]}]`,
			shouldInclude: true,
		},
		{
			name:          "multi-field nested filter should include",
			filterJSON:    `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`,
			shouldInclude: true,
		},
		{
			name:          "simple enum filter should exclude",
			filterJSON:    `[{"kind": {"eq": "SERIES"}}]`,
			shouldInclude: false,
		},
		{
			name:          "simple string filter should exclude",
			filterJSON:    `[{"title": {"eq": "The Shawshank Redemption"}}]`,
			shouldInclude: false,
		},
		{
			name:          "simple number filter should exclude",
			filterJSON:    `[{"movieYear": {"eq": 1955}}]`,
			shouldInclude: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var movieFilters []*MovieFilter
			jsonx.MustUnmarshal([]byte(tc.filterJSON), &movieFilters)

			result := ShouldInclude(movieFilters, movie)

			assert.Equal(t, tc.shouldInclude, result)
		})
	}
}
