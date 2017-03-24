plasma-go
==========

[![Circle CI](https://circleci.com/gh/openfresh/plasma-go.svg?style=shield&circle-token=)](https://circleci.com/gh/openfresh/plasma-go)
[![Language](https://img.shields.io/badge/language-go-brightgreen.svg?style=flat)](https://golang.org/)
[![issues](https://img.shields.io/github/issues/openfresh/plasma-go.svg?style=flat)](https://github.com/openfresh/plasma-go/issues?state=open)
[![License: MIT](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/openfresh/plasma-go?status.png)](https://godoc.org/github.com/openfresh/plasma-go)

[openfresh/plasma](https://github.com/openfresh/plasma) client library for Golang 


## Installation

Install:

```shell
go get -u github.com/openfresh/plasma-go
```

Import:

```go
import "github.com/openfresh/plasma-go"
```

## Usage

```go
package main

import (
	"encoding/json"
	"log"

	"github.com/openfresh/plasma-go"
	"github.com/openfresh/plasma-go/config"
	"github.com/openfresh/plasma-go/event"
)

func main() {
	config := config.Config{
		Type: plasma_client.TypeRedis,
	}

	client, err := plasma_client.New(config)
	if err != nil {
		log.Fatal(err)
	}

	payload := event.Payload{
		Meta: event.MetaData{
			Type: "video:1234:views",
		},
		Data: json.RawMessage(`{"data":55301}`),
	}

	err = client.Publish(payload)
	if err != nil {
		log.Fatal(err)
	}
}

```

## Documents

[GoDoc](https://godoc.org/github.com/openfresh/plasma-go)

License
===
See [LICENSE](LICENSE).

Copyright © CyberAgent, Inc. All Rights Reserved.
