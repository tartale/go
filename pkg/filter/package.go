package filter

// Copyright 2023 Tom Artale. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// Package filter provides the capability to do generic
// filtering of objects by their fields. The filter
// expression can be serialized to/from json, making
// it useful for use in APIs.
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
// fields using a JSON document that describes a boolean
// expression; for example:
//
// 		movies := GetAllMovies()
//
// 		movieFilter := NewTypeFilter[Movie]
// 		movieFilterJson := `[{"title": {"eq": "Back to the Future"}}]`
//		jsonx.MustUnmarshalFromString(movieFilterJson, &movieFilter)
// 		filteredMovies := typeFilter.Filter(slices.Values(movies))
// 		filteredMovieValues := slices.Collect(filteredMovies)
//
// Note: there is also a convenience function to create and populate
// the filter in one step:
//
//    movieFilter := NewTypeFilterFromJson[Movie](`[{"title": {"eq": "Interstellar"}}]`)
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
// 		// Wrap the MovieFilter object in a TypeFilter, then it can
// 		// be used just like the dynamically-created one.
// 		typeFilter := TypeFilter[Movie]{filters}
// 		filteredMovies := typeFilter.Filter(slices.Values(movies))
// 		filteredMovieValues := slices.Collect(filteredMovies)
// 	}
//
