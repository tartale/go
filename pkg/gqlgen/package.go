package gqlgen

// Copyright 2023 Tom Artale. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// Package gqlgen contains some utility functions to
// help work with the context provided by the generated
// gqlgen server. The provided context contains information
// about the query, such as the arguments that were
// passed to sub-queries (that aren't available without
// a separate resolver).
// For more information about the gqlgen Go GraphQL generator,
// see https://gqlgen.com
