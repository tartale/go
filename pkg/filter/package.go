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
// and you would like to be able to filter a list of
// these objects based on Title or MovieYear. To use this package,
// you would define a corresponding struct:
//
// 		type MovieFilter struct {
// 			Title     *filter.Operator `json:"title,omitempty"`
// 			MovieYear *filter.Operator `json:"movieYear,omitempty"`
// 			And       []*MovieFilter   `json:"and,omitempty"`
// 			Or        []*MovieFilter   `json:"or,omitempty"`
// 		}
//
// This struct can be used in your API like this:
//
// 	POST /movies/search
// 	{
// 		"filters": [
// 			"title": {
// 				"eq": "Back to the Future"
// 			},
// 			"and": {
// 				"movieYear": {
// 					"eq": 1985
// 				}
// 			}
// 		]
// 	}
//
// Then, the resolver for this API call would look something like this:
//
// 	func MovieSearch(ctx context.Context, filters: []MovieFilter) []Movie {
// 		allMovies := getAllMovies()
//
// 		expression := filter.GetExpression(filters)
// 		for _, movie := range allMovies {
// 			values := filter.GetValues(filters, movie)
// 			shouldInclude, _ := gval.Evaluate(expression, values)
// 			if shouldInclude.(bool) {
// 				// movie has passed the filter
// 			}
// 		}
// 	}
//
