# API Reference

## REST API

### GET /api/v1

Provides an overview of the system: A status and message, plus a simplified hierarchy of sources (not surveys).

    {
        "status": 200,
        "message": "ok",
        "result": {
            "sources": [
                {
                    "name": "[adjective][species]",
                    "slug": "adjspecies",
                    "sources": [
                        {
                            "name": "The Furry Poll",
                            "slug": "furrypoll"
                            "sources": {
                                "name": "Mega polls",
                                "slug": "megapolls",
                                "sources": {
                                    "name": "Mega polls - 2018",
                                    "slug": "megapolls2018"
                                }
                            }
                        },
                        {
                            "name": "Microsurveys",
                            "slug": "microsurveys"
                        }
                        ...
                    ]
                },
                ...
            ],
            "summaries": [
                {
                    "name": "..."
                    "slug": "why-furries-rock",
                    "source": {
                        "name": "Mega polls - 2018",
                        "slug": "megapolls2018"
                    },
                    "surveys": {},
                    "description": "...",
                    "link": "..."
                },
                ...
            ]
        }
    }

### GET /api/v1/summary/<source>/<summary>

...

### PUT /api/v1/summary/<source>

### GET /api/v1/source/<source>

Provides an in-depth view of a source and all of its component surveys, plus a hierarchy of sources (not surveys).

    {
        "status": 200,
        "message": "ok"
        "result": {
            "name": "Mega polls - 2018",
            "slug": "megapolls2018",
            "description": "...",
            "link": "http://mega.furrypoll.com//2018",
            "surveys": [
                {
                    "name": "Demographics",
                    "slug": "demographics",
                    "start_date": ...,
                    "end_date": ...,
                    "description": "..."
                },
                {
                    "name": "Species",
                    "slug": "species",
                    "start_date": ...,
                    "end_date": ...,
                    "description": "..."
                },
                ...
            ],
            "hierarchy": {
                "parents": [
                    {
                        "type": "source",
                        "name": "[adjective][species]",
                        "slug": "adjspecies"
                    },
                    {
                        "type": "source",
                        "name": "The Furry Poll",
                        "slug": "furrypoll"
                    },
                    {
                        "type": "source",
                        "name": "Mega polls",
                        "slug": "megapolls"
                    }
                ]
                "siblings": [
                    {
                        "type": "source",
                        "name": "Mega polls - 2019",
                        "slug": "megapolls2019"
                    }
                ],
                "children": []
            }
        }
    }

Here is an example with no surveys but a complex hierarchy:

    {
        "status": 200,
        "message": "ok"
        "result": {
            "name": "Mega polls",
            "slug": "megapolls",
            "description": "...",
            "link": "http://mega.furrypoll.com",
            "surveys": [],
            "hierarchy": {
                "parents": [
                    {
                        "type": "source",
                        "name": "[adjective][species]",
                        "slug": "adjspecies"
                    },
                    {
                        "type": "source",
                        "name": "The Furry Poll",
                        "slug": "furrypoll"
                    }
                ],
                "siblings": [
                    ...
                ]
                "children": [
                    {
                        "type": "source",
                        "name": "Mega polls - 2018",
                        "slug": "megapolls2019"
                    },
                    {
                        "type": "source",
                        "name": "Mega polls - 2019",
                        "slug": "megapolls2019"
                    }
                ]
            ]
        }
    }

### GET /api/v1/survey/<source>/<survey>

Provides an overview of a survey and its hierarchy.

    {
        "status": 200,
        "message": "ok",
        "result": {
            "name": "Demographics",
            "slug": "demograpahics",
            "start_date": ...,
            "end_date": ...,
            "description": "...",
            "response_count": "329",
            "questions": [
                {
                    "type": "text",
                    "key": "name",
                    "text": "What is your name?"
                },
                {
                    "type": "number",
                    "key": "age",
                    "text": "What is your age?"
                },
                ...
            ],
            "hierarchy": {
                "parents": [
                    ...
                ],
                "siblings": [
                    {
                        "type": "survey",
                        "name": "Species",
                        "slug": "species"
                    }
                ]
            }
        }
    }

### GET /api/v1/survey/<source>/<survey>.json

Params:
* `touchpoints` (bool default=false): include touchpoints
* `limit`(int default=1000): limit results to value

Dumps the questions and answers for a survey (example with `touchpoints` set to true).

    {
        "status": 200,
        "message": "ok",
        "results": {
            "questions": [
                {
                    "type": "text",
                    "key": "name",
                    "text": "What is your name?"
                },
                {
                    "type": "number",
                    "key": "age",
                    "text": "What is your age?"
                },
                ...
            ],
            "responses": [
                {
                    "respondent": {
                        "id": "...",
                        "metadata": {}
                    },
                    "answers": [
                        {
                            "question": "name",
                            "type": "text",
                            "value": "Rose"
                        },
                        {
                            "question": "age",
                            "type": "number",
                            "value": 42
                        },
                        ...
                    ]
                }
            ],
            "touchpoints" [
                {
                    "type": "start",
                    "value": "",
                    "date_time": ...
                },
                {
                    "type": "answer",
                    "value": "1", //question 1
                    "date_time": ...
                }
            ]
        }
    }

### GET /api/v1/survey/<source>/<survey>.csv

Params:
* `lists` (bool default=false): include list responses as JSON
* `objects` (bool default=false): include object responses as JSON
* `limit`(int default=1000): limit results to value

Dumps the questions and answers for a survey into CSV format. Lists and objects will be omitted unless requested, then they will be represented as JSON.

### GET /api/v2/respondent/<respondent>

Retrieves information about a respondent.

    {
        "status": 200,
        "message": "ok",
        "result": {
            "id": "...",
            "metadata": {},
            "responses": [
                {
                    "survey": {
                        "name": "Demographics",
                        "slug": "demographics"
                    },
                    "source": {
                        "name": "Mega polls - 2018",
                        "slug": "megapolls2018"
                    }
                    "url": "megapolls2018/demographics"
                }
            ]
        }
    }

### POST /api/v1/query

Queries the database. (Use something similar to ffmddb's query language?)

## GraphQL API

### POST /api/graphql/v1

...
