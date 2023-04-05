# gqlgen-tracing-sentry ![Test](https://github.com/benevolent0505/gqlgen-sentry-tracing/.github/workflows/test.yml/badge.svg)

[Sentry](https://sentry.io) Tracer for [gqlgen](https://gqlgen.com/)

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
