package stringz

import (
	"regexp"
)

var SpecialCharsRE = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func ToAlphaNumeric(s string) string {

	return SpecialCharsRE.ReplaceAllString(s, "")
}
