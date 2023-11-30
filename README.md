# gqlgen-tracing-sentry

[![Test](https://github.com/benevolent0505/gqlgen-tracing-sentry/actions/workflows/test.yml/badge.svg)](https://github.com/benevolent0505/gqlgen-tracing-sentry/actions/workflows/test.yml)

[Sentry](https://sentry.io) Tracer for [gqlgen](https://gqlgen.com/)

## Installation

```sh
go get github.com/benevolent0505/gqlgen-tracing-sentry@latest
```

## Usage

```go
package main

import (
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/benevolent0505/gqlgen-tracing-sentry/sentrytracing"
)

func main() {
    srv := handler.NewDefaultServer(...)
    srv.Use(sentrytracing.Tracer{})
}
```
