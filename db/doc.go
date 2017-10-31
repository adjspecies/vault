// Copyright 2017 [adjective][species], Ltd
// Licensed under the MIT license, see the LICENSE file for details.

// Package db manages database connections and object relational models.
//
// Data models
//
// Vault stores data in three main models:
//
// * Sources describe collections of surveys, nominally from an organization
// * Surveys describe a collection of questions and their responses
// * Summaries describe and link to a summary providing information gleaned from
// one or more surveys.
//
// When it comes to responses to surveys, data is stored in two main models:
//
// * Respondents represent someone who has answered one or more surveys within a
// source.
// * Responses represent the list of answers someone has provided to a survey.
//
// Responses are broken down into three main models:
//
// * Questions represent a question that was asked on the survey. They can be
// simple or complex (lists of answers, structured answers, ...)
// * Answers represent an answer to a question and are tied to a response.
// * Touchpoints represent an action that a user took during the survey (such as
// answering a question, clicking submit, moving to the next page) along with
// a timestamp.
//
// Vault data is stored in a hierarchy.
//
// Data is stored as data sets called `surveys`. These are grouped together in
// `sources`, and these sources may belong to a hierarchy of other sources.
// Additionally, `summaries` of data may belong to these sources.
//
// The hierarchy allows for a descriptive means of storing data from multiple
// sources. For example, [a][s] runs several surveys, so we might have an
// `adjspecies` source. Several surveys are grouped together under the Furry
// Poll heading, which might be named `furrypoll`. These surveys happen every
// year. Given this, our hierarchy might look like so:
//  adjspecies (source)
//    furrypoll (source)
//      2018 (source)
//        yearly-summary (summary)
//        demographics (survey)
//        interests (survey)
//        ...
package db
