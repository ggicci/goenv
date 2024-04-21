# goenv

Populate go struct instance with OS environment variables.

[![Go](https://github.com/ggicci/goenv/actions/workflows/go.yaml/badge.svg)](https://github.com/ggicci/goenv/actions/workflows/go.yaml)
[![codecov](https://codecov.io/gh/ggicci/goenv/graph/badge.svg?token=IYVTITZZXS)](https://codecov.io/gh/ggicci/goenv)
[![Go Report Card](https://goreportcard.com/badge/github.com/ggicci/goenv)](https://goreportcard.com/report/github.com/ggicci/goenv)
[![Go Reference](https://pkg.go.dev/badge/github.com/ggicci/goenv.svg)](https://pkg.go.dev/github.com/ggicci/goenv)

**LIMITATION**: only string fields can be tagged with `env` tag.

```go
type MyAppConfig struct {
	Workspace string `env:"name=MYAPP_HOME"`
	User      string `env:"name=MYAPP_USER"`
	Debug     string `env:"name=MYAPP_DEBUG"`
}

var config = new(MyAppConfig)
if err := goenv.Load(config); err != nil { ... }

if config.Debug == "1" {
    // enable debug output
}
```
