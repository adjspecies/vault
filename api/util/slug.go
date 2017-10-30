// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package util

import (
	"regexp"
	"strings"
)

var (
	replaceCharacters = regexp.MustCompile(`[\W\s]`)
	dashes            = regexp.MustCompile(`-+`)
	dash              = []byte(`-`)
)

// Slug represents a string that is safe to use in a URL.
type Slug string

// NewSlug generates a slug string from a given input.
func NewSlug(from string) Slug {
	slug := Slug(
		strings.Trim(
			string(
				dashes.ReplaceAll(
					replaceCharacters.ReplaceAll(
						[]byte(from),
						dash),
					dash)),
			string(dash)))
	return slug
}
