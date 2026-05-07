package filter

import (
	"encoding/json"
	"testing"

	"github.com/PaesslerAG/gval"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tartale/go/pkg/jsonx"
)

type filterTestCase struct {
	name       string
	filterJSON string
	strict     bool
	shouldPass bool
}

func TestGetExpression_Evaluates(t *testing.T) {
	movie := Movie{
		Kind:        MOVIE,
		Title:       "Back to the Future",
		Description: "The time travel adventures of Doc Brown and Marty McFly",
		MovieYear:   1985,
	}

	cases := []filterTestCase{
		{
			name:       "simple enum filter",
			filterJSON: `[{"kind": {"eq": "MOVIE"}}]`,
			strict:     true,
			shouldPass: true,
		},
		{
			name:       "simple string filter",
			filterJSON: `[{"title": {"eq": "Back to the Future"}}]`,
			strict:     true,
			shouldPass: true,
		},
		{
			name:       "simple number filter",
			filterJSON: `[{"movieYear": {"eq": 1985}}]`,
			strict:     true,
			shouldPass: true,
		},
		{
			name:       "multi-field filter with implied logical 'and'",
			filterJSON: `[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1985}}]`,
			strict:     true,
			shouldPass: true,
		},
		{
			name:       "multi-field filter with explicit logical 'or'",
			filterJSON: `[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": "Back to the .*"}}]}]`,
			strict:     true,
			shouldPass: true,
		},
		{
			name:       "multi-field nested filter",
			filterJSON: `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "Back to the Future"}}]}]}]`,
			strict:     true,
			shouldPass: true,
		},
		{
			name:       "simple enum filter false",
			filterJSON: `[{"kind": {"eq": "SERIES"}}]`,
			strict:     false,
			shouldPass: false,
		},
		{
			name:       "simple string filter false",
			filterJSON: `[{"title": {"eq": "The Shawshank Redemption"}}]`,
			strict:     false,
			shouldPass: false,
		},
		{
			name:       "simple number filter false",
			filterJSON: `[{"movieYear": {"eq": 1955}}]`,
			strict:     false,
			shouldPass: false,
		},
		{
			name:       "multi-field filter with implied logical 'and' false",
			filterJSON: `[{"movieYear": {"eq": 1955}}, {"title": {"matches": "Back to the .*"}}]`,
			strict:     false,
			shouldPass: false,
		},
		{
			name:       "multi-field filter with explicit logical 'or' false",
			filterJSON: `[{"movieYear": {"eq": 1955}}, {"or": [{"title": {"matches": ".*Shawshank.*"}}]}]`,
			strict:     false,
			shouldPass: false,
		},
		{
			name:       "multi-field nested filter false",
			filterJSON: `[{"movieYear": {"eq": 1955}}, {"or": [{"movieYear": {"eq": 1985}}, {"and": [{"title": {"eq": "The Shawshank Redemption"}}]}]}]`,
			strict:     false,
			shouldPass: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var movieFilters []*MovieFilter
			var err error

			if tc.strict {
				err = jsonx.StrictUnmarshal([]byte(tc.filterJSON), &movieFilters)
			} else {
				err = json.Unmarshal([]byte(tc.filterJSON), &movieFilters)
			}

			require.NoError(t, err)

			expression := GetExpression(movieFilters)
			values := GetValues(movieFilters, movie)
			eval, err := gval.Evaluate(expression, values)
			require.NoError(t, err)

			result, ok := eval.(bool)
			require.True(t, ok, "expected boolean result")

			if tc.shouldPass {
				assert.True(t, result)
			} else {
				assert.False(t, result)
			}
		})
	}
}
