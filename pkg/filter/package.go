package filter

// Copyright 2023 Tom Artale. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// Package filter provides the capability to do generic
// filtering of lists of objects by their fields, using a JSON
// boolean expression language, powered by the gval library.
//
// This package includes both a set of functions that can
// orchestrate the filtering of slices, as well as
// a way to create this capability dynamically for
// any given struct.
//
// Example:
// Suppose you have the following data structure:
//
// 		type Movie struct {
// 			Title       string   `json:"title"`
// 			Description string   `json:"description"`
// 			MovieYear   int      `json:"movieYear"`
// 		}
//
// You can create an object that is able to filter any of these
// fields using a JSON string that describes a boolean
// expression; for example:
//
// 		movies := GetAllMovies()
//
// 		movieFilterJson := `[{"title": {"eq": "Back to the Future"}}]`
// 		structFilter := NewStructFilter[Movie](movieFilterJson)
// 		filteredMovies := filter.FilterAllFor[Movie](structFilter, movies)
//
// The JSON expression operators (e.g. "eq") are defined in the filter.Operator
// struct. You can also create compound filters (using "and" and "or" operators).
// The tests include several examples.
//
// Additionally, you can "bring your own filter" by defining a struct that uses the
// filter.Operator type. This is useful if you want the API to have a well-known
// structure that can be used for swagger documentation, code generation, graphQL
// definitions, etc.  For example:
//
// 		type MovieFilter struct {
// 			Title     *filter.Operator `json:"title,omitempty"`
// 			MovieYear *filter.Operator `json:"movieYear,omitempty"`
// 			And       []*MovieFilter   `json:"and,omitempty"`
// 			Or        []*MovieFilter   `json:"or,omitempty"`
// 		}
//
// Then, the resolver for this API call would look something like this:
//
// 	func MovieSearch(ctx context.Context, filters: []MovieFilter) []Movie {
// 		movies := GetAllMovies()
//
//    filterJson := jsonx.MustMarshalToString(filters)
// 		structFilter := NewStructFilter[Movie](filterJson)
// 		filteredMovies := filter.FilterAllFor[Movie](structFilter, movies)
// 	}
//
