# JSON-2-Go

[![Pipeline](https://gitlab.com/ahmedalhulaibi/json2go/badges/master/pipeline.svg)](https://gitlab.com/ahmedalhulaibi/json2go/pipelines)[![codecov](https://codecov.io/gl/ahmedalhulaibi/json2go/branch/master/graph/badge.svg)](https://codecov.io/gl/ahmedalhulaibi/json2go)

This tool converts JSON into a Go type definition. This is a Go implementation of Matt Holt's JSON-to-Go.

# Installation
This project is hosted on gitlab and mirrored to github

`go get -u gitlab.com/ahmedalhulaibi/json2go/...`

# Usage

```
curl http://api.open-notify.org/astros.json -s | json2go-cli --typename Astronaut

type Astronaut struct {
	Message string `json:"message"`
	Number  int    `json:"number"`
	People  []struct {
		Craft string `json:"craft"`
		Name  string `json:"name"`
	} `json:"people"`
}


```

Nested JSON objects are resolved to inline type definitions