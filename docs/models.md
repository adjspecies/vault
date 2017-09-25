# Models

## Source

* Name [String]
* Slug [Slug] (unique, primary key)
* Description [String]
* Link [URL]
* Parent? -> Source

## Survey

* Title [String]
* Slug [Slug] (source slug plus survey slug unique)
* Date Start [Date]
* Date End [Date]
* Source -> Source
* Description [String]
* Link [URL]

## Summary

* Title [String]
* Slug [Slug] (source slug plus summary slug unique)
* Source -> Source
* Surveys? -> Survey[]
* Link [URL]
* Description [String]
* Content? [String (Jupyter? Zeppelin?)]

## Respondent

* ID [String] (unique, primary key)
* Metadata? [Object]

## Response

* Survey -> Survey
* Respondent -> Respondent
* Date [Date]

## Touchpoint

* Type [String: 'start', 'end', 'page', 'answer', 'other']
* Value [String]
* DateTime [DateTime]

## Question

* Survey -> Survey
* Key [slug] (unique)
* Text [String]
* Options? [Object] (JSONSchema)

## IAnswer

* Question -> Question
* Response -> Response

## IListAnswer implements IAnswer

## TextAnswer implements IAnswer

* Value [String]

## TextListAnswer implements IListAnswer

* Values [String[]]

## NumberAnswer implements IAnswer

* Value [Number]

## NumberListAnswer implements IAnswer

* Values [Number[]]

## ObjectAnswer implements IAnswer

* Value [Object]

## ObjectListAnswer implements IAnswer

* Values [Object[]]
