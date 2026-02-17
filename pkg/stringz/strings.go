// Package stringz provides helpers for working with strings.
package stringz

import "regexp"

// SpecialCharsRE matches any non-alphanumeric character.
var SpecialCharsRE = regexp.MustCompile(`[^a-zA-Z0-9]+`)

// ToAlphaNumeric removes all non-alphanumeric characters from s.
//
// Example:
//
//	out := stringz.ToAlphaNumeric("a-b_c!") // "abc"
func ToAlphaNumeric(s string) string {
	return SpecialCharsRE.ReplaceAllString(s, "")
}
