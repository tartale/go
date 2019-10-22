// Copyright 2019 Tom Artale. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// Package errors implements functions to manipulate errors.
package errors

import (
	"errors"
	"fmt"
	"strings"
)

type Errors []error

func (o *Errors) Combine(message string, separator string) error {
	if len(*o) == 0 {
		return nil
	}
	var errorStrings []string
	for _, err := range *o {
		errorStrings = append(errorStrings, err.Error())
	}

	if message == "" {
		return errors.New(strings.Join(errorStrings, separator))
	}

	return errors.New(fmt.Sprintf("%s: %s", message, strings.Join(errorStrings, separator)))
}

func (o *Errors) Error() string {
	if len(*o) == 0 {
		return ""
	}
	return o.Combine("", "; ").Error()
}

func (o *Errors) Try(f func() error) {
	err := f()
	if err != nil {
		*o = append(*o, err)
	}
}
