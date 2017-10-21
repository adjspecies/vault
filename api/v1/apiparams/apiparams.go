// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

package apiparams

import (
	"net/url"
	"time"

	"github.com/adjspecies/vault/internal"
)

// APIResponse holds a JSON response along with an HTTP status and a message string.
type APIResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Result  responseContent `json:"result"`
}

type responseContent interface{}

// Sluggable is an embedded type to allow for an object to have a type, name, and slug field.
type Sluggable struct {
	Type string        `json:"type"`
	Name string        `json:"string"`
	Slug internal.Slug `json:"slug"`
}

// ShortSource is the bare necessity required to reference a data source. It may have children.
type ShortSource struct {
	Sluggable
	Sources []ShortSource `json:"sources"`
}

// Source represents a data source. It has metadata, along with a listing of its hierarchy.
type Source struct {
	Sluggable
	Description string        `json:"description"`
	Link        url.URL       `json:"link"`
	Surveys     []ShortSurvey `json:"surveys"`

	// Hierarchy contains the data source's parents, siblings, and children.
	Hierarchy struct {

		// Parents is an ordered list of parent sources up to the top, parentless source.
		Parents []Sluggable `json:"parents"`

		// Siblings is an unordered list of sources and surveys who share the same parent as this source.
		Siblings []Sluggable `json:"siblings"`

		// Children is an unordered list of sources (surveys fit, but, by convention, are not included) which claim this source as a parent.
		Children []Sluggable `json:"children"`
	} `json:"hierarchy"`
}

// ShortSurvey is the bare necessity required to reference a survey.
type ShortSurvey struct {
	Sluggable
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}

// Survey represents a survey. It has metadata, questions, and a hierarchy.
type Survey struct {
	Sluggable
	StartDate     time.Time  `json:"start_date"`
	EndDate       time.Time  `json:"end_date"`
	Description   string     `json:"description"`
	ResponseCount int        `json:"response_count"`
	Questions     []Question `json:"questions"`

	// Hierarchy contains the data source's parents and siblings.
	Hierarchy struct {
		// Parents is an ordered list of parent sources up to the top, parentless source.
		Parents []Sluggable `json:"parents"`

		// Siblings is an unordered list of sources and surveys who share the same parent as this source.
		Siblings []Sluggable `json:"siblings"`
	} `json:"hierarchy"`
}

// SurveyDump is a list of questions and responses for a particular survey.
type SurveyDump struct {
	Questions []Question `json:"questions"`
	Responses []Response `json:"responses"`
}

// Response represents a single response to a survey. This includes the respondent, their answers, and their actions (touchpoints) from the process of answering the survey.
type Response struct {
	Respondent  ShortRespondent `json:"respondent"`
	Answers     []Answer        `json:"answers"`
	Touchpoints []Touchpoint    `json:"touchpoints"`
}

// ShortRespondent is the bare necessity required to reference a respondent.
type ShortRespondent struct {
	ID       string                 `json:"id"`
	Metadata map[string]interface{} `json:"metadata"`
}

// Respondent represents a respondent, all of the surveys they have taken, and where to find them.
type Respondent struct {
	ShortRespondent
	Responses []struct {
		Survey Sluggable `json:"survey"`
		Source Sluggable `json:"source"`
		URL    url.URL   `json:"url"`
	} `json:"responses"`
}

// Question represents a question, its text, and optionally its possible values and any child questions.
type Question struct {
	Type     string     `json:"type"`
	Key      string     `json:"key"`
	Text     string     `json:"text"`
	Items    []string   `json:"items"`
	Children []Question `json:"children"`
}

// Answer represents an answer to a question.
type Answer struct {
	Question string `json:"question"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

// Touchpoint represents a respondent's interaction with a survey. This can be answering a question, starting or ending the survye, hitting the "next page" button, etc.
type Touchpoint struct {
	Type  string    `json:"type"`
	Value string    `json:"value"`
	When  time.Time `json:"when"`
}

// Summary represents a writeup or summary of a data source.
type Summary struct {
	Sluggable
	Source      Source  `json:"source"`
	Description string  `json:"description"`
	Link        url.URL `json:"link"`
}

// Overview contains some basic information of what's in the vault.
type Overview struct {
	Sources   []ShortSource `json:"sources"`
	Summaries []Summary     `json:"summaries"`
}
