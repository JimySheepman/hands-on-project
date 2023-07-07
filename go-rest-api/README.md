# THE CHALLENGE

We’d like you to create a RESTful API with two endpoints.

## Required Information

1. One of them that fetches the data in the provided MongoDB collection and returns the results
in the requested format.

- **Request Payload**

    The request payload of the first endpoint will include a JSON with 4 fields.
  - “startDate” and “endDate” fields will contain the date in a “YYYY-MM-DD” format.
    You should filter the data using “createdAt”
  - “minCount” and “maxCount” are for filtering the data. Sum of the “count” array in the documents should be between “minCount” and “maxCount”.

Sample:

```Json
{
    "startDate": "2016-01-26",
    "endDate": "2018-02-02",
    "minCount": 2700,
    "maxCount": 3000
}
```

- **Response Payload**

    Response payload should have 3 main fields.
  - “code” is for status of the request. 0 means success. Other values may be used   for errors that you define.
  - “msg” is for description of the code. You can set it to “success” for successful
requests. For unsuccessful requests, you should use explanatory messages.
  - “records” will include all the filtered items according to the request. This array
should include items of “key”, “createdAt” and “totalCount” which is the sum the
“counts” array in the document.

Sample:

```Json
{
    "code":0,
    "msg":"Success",
    "records": [
        {
            "key":"TAKwGc6Jr4i8Z487",
            "createdAt":"2017-01-28T01:22:14.398Z",
            "totalCount":2800
        },
        {
        "key":"NAeQ8eX7e5TEg7oH",
        "createdAt":"2017-01-27T08:19:14.135Z",
        "totalCount":2900
        }
    ]
}
```

2. Second endpoint is to create(POST) and fetch(GET) data from an in-memory database.

- POST Endpoint
  - **Request Payload**

    The request payload of POST endpoint will include a JSON with 2 fields.
    - “key” fields holds the key (any key in string type)
    - “value” fields holds the value (any value in string type)
Sample:

```Json
{
"key": "active-tabs",
"value": "getir"
}
```

- **Response Payload**

    Response payload should return echo of the request or error (if any).

- GET Endpoint
  - **Request Payload**

    The request payload of GET endpoint will include 1 query parameter. That is “key”
    param holds the key (any key in string type)

Sample:

`http://localhost/in-memory?key=active-tabs`

- **Response Payload**

    Response payload of GET endpoint should return a JSON with 2 fields or error (if any)
  - “key” fields holds the key
  - “value” fields holds the value

Sample:

```Json
{
"key": "active-tabs",
"value": "getir"
}
```

## Endpoint List

```Bash
/api/v1/fetch-data  # POST
/api/v1/in-memory   # POST
/api/v1/in-memory   # GET
```

## Quick Start

```Bash
# make file 
$ make -f Makefile

# make file  run for windows
$ make buildw

# command line start
$ go run main.go

# Dockerfile
$ docker build -t app .
$ docker run --name app -d -p 8000:8000 app:latest 
```

## Project Layout

```Bash
.
├── config
│   ├── db
│   │   ├── db.go
│   │   └── db_test.go
│   └── env
│       ├── config.go
│       └── config_test.go
├── internal
│   ├── handler
│   │   ├── fetchDataHandler.go
│   │   ├── fetchDataHandler_test.go
│   │   ├── inMemoryHandler.go
│   │   └── inMemoryHandler_test.go
│   ├── router
│   │   └── routes.go
│   ├── model
│   │   ├── memory.go
│   │   └── records.go
│   └── times
│       ├── timeConverter.go
│       ├── timeConverter_test.go
│       ├── timeFormatValidator.go
│       └── timeFormatValidator_test.go
├── Dockerfile
├── example.env
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md
```

## Test Coverage

```Bash
$ go test -cover

PASS
coverage: 100.0% of statements
ok      github.com/JimySheepman/go-rest-api/config/env 

PASS
coverage: 33.3% of statements
ok      github.com/JimySheepman/go-rest-api/config/db   0.002s

PASS
coverage: 88.9% of statements
ok      github.com/JimySheepman/go-rest-api/internal/times    0.001s

PASS
coverage: 85.1% of statements
ok      github.com/JimySheepman/go-rest-api/internal/handler    5.759s
```
