package filter

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
	Kind        *Operator      `json:"kind,omitempty"`
	Title       *Operator      `json:"title,omitempty"`
	Description *Operator      `json:"description,omitempty"`
	MovieYear   *Operator      `json:"movieYear,omitempty"`
	And         []*MovieFilter `json:"and,omitempty"`
	Or          []*MovieFilter `json:"or,omitempty"`
}

var testMovie = Movie{
	Kind:        MOVIE,
	Title:       "Back to the Future",
	Description: "The time travel adventures of Doc Brown and Marty McFly",
	MovieYear:   1985,
}
