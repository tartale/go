// Package errorx
// Copyright 2019 Tom Artale. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// Package errorx implements functions to manipulate errors.
package errorx

import (
	"errors"
	"fmt"
	"strings"
)

// Errors is an abstraction for a slice of errors
type Errors []error

// Combine takes the slice of errors and combines it into a single
// error that represents the errors in a single message.
func (o Errors) Combine(message string, separator string) error {
	var errorStrings []string
	for _, err := range o {
		if err == nil {
			continue
		}
		errorStrings = append(errorStrings, err.Error())
	}
	if len(errorStrings) == 0 {
		return nil
	}

	if message == "" {
		return errors.New(strings.Join(errorStrings, separator))
	}

	return errors.New(fmt.Sprintf("%s%s", message, strings.Join(errorStrings, separator)))
}

func (o Errors) Error() string {
	if len(o) == 0 {
		return ""
	}
	return o.Combine("", "; ").Error()
}
