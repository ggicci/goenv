package goenv

import (
	"fmt"
	"os"
	"reflect"

	"github.com/ggicci/owl"
)

func nameDirective(rtm *owl.DirectiveRuntime) error {
	if len(rtm.Directive.Argv) == 0 {
		return nil
	}
	if value, ok := os.LookupEnv(rtm.Directive.Argv[0]); ok {
		if rtm.Value.Elem().Type().Kind() != reflect.String {
			return fmt.Errorf("only string field can be tagged with env")
		}
		rtm.Value.Elem().SetString(value)
	}
	return nil
}

func init() {
	owl.UseTag("env")
	owl.RegisterDirectiveExecutor("name", owl.DirectiveExecutorFunc(nameDirective))
}

// Load populates the fields of the value with the values of the environment variables.
func Load(value interface{}) error {
	resolver, err := owl.New(value)
	if err != nil {
		return err
	}
	return resolver.ResolveTo(value)
}
