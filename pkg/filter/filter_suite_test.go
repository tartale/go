package filter_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tartale/go/pkg/filter"
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

type MovieFilter struct {
	Kind        *filter.Operator `json:"kind,omitempty"`
	Title       *filter.Operator `json:"title,omitempty"`
	Description *filter.Operator `json:"description,omitempty"`
	MovieYear   *filter.Operator `json:"movieYear,omitempty"`
	And         []*MovieFilter   `json:"and,omitempty"`
	Or          []*MovieFilter   `json:"or,omitempty"`
}

var testMovie = Movie{
	Kind:        MOVIE,
	Title:       "Back to the Future",
	Description: "The time travel adventures of Doc Brown and Marty McFly",
	MovieYear:   1985,
}

var testMovies = []Movie{
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

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filter Test Suite")
}
