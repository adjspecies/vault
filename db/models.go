// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package db

// Source holds a data source such as an organization or collection of surveys.
// It may contain surveys and other sources
type Source struct{}

// Survey holds a set of responses from a single survey.
type Survey struct{}

// Summary holds information about a data summary such as an article or
// visualization.
type Summary struct{}

// Response holds all of the answers provided by a respondent to a single
// survey.
type Response struct{}

// Respondent represents a user who created a response to one or more surveys.
type Respondent struct{}

// Question holds information about a question in a survey, including type and
// any sub-questions.
type Question struct{}

// Answer holds a response to a single question.
type Answer struct{}

// Touchpoint holds information about an action a user took on a survey.
type Touchpoint struct{}
