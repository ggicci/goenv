# goenv

Populate go struct instance with OS environment variables.

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
