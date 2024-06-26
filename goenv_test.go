package goenv_test

import (
	"errors"
	"os"
	"testing"

	"github.com/ggicci/goenv"
)

type MyAppConfig struct {
	Workspace string `env:"name=MYAPP_HOME"`
	User      string `env:"name=MYAPP_USER"`
	Debug     string `env:"name=MYAPP_DEBUG"`
}

func TestGoenv(t *testing.T) {
	// Set environment variables
	_ = os.Setenv("MYAPP_HOME", "/home/myapp")
	_ = os.Setenv("MYAPP_USER", "admin")
	_ = os.Setenv("MYAPP_DEBUG", "1")

	// Load configuration
	config := &MyAppConfig{}
	err := goenv.Load(config)
	if err != nil {
		t.FailNow()
	}
	if config.Workspace != "/home/myapp" {
		t.FailNow()
	}
	if config.User != "admin" {
		t.FailNow()
	}
	if config.Debug != "1" {
		t.FailNow()
	}

	os.Clearenv()
}

func TestGoenv_IgnoreUnsetEnvVar(t *testing.T) {
	// Set environment variables
	_ = os.Setenv("MYAPP_HOME", "/home/myapp")

	// Load configuration
	config := &MyAppConfig{
		User: "myapp",
	}
	err := goenv.Load(config)
	if err != nil {
		t.FailNow()
	}
	if config.Workspace != "/home/myapp" {
		t.FailNow()
	}
	if config.User != "myapp" {
		t.FailNow()
	}
	if config.Debug != "" {
		t.FailNow()
	}

	os.Clearenv()
}

func TestGoenv_ErrMissingName(t *testing.T) {
	type ConfigWithoutName struct {
		Workspace string `env:"name"`
	}

	// Load configuration
	config := &ConfigWithoutName{}
	err := goenv.Load(config)
	if err == nil {
		t.FailNow()
	}
	if !errors.Is(err, goenv.ErrMissingName) {
		t.FailNow()
	}
}

func TestGoenv_ErrInvalidType(t *testing.T) {
	type ConfigWithInvalidType struct {
		Workspace int  `env:"name=MYAPP_HOME"`
		Debug     bool `env:"name=MYAPP_DEBUG"`
	}

	// Set environment variables
	_ = os.Setenv("MYAPP_HOME", "/home/myapp")
	_ = os.Setenv("MYAPP_DEBUG", "true")

	// Load configuration
	config := &ConfigWithInvalidType{}
	err := goenv.Load(config)
	if err == nil {
		t.FailNow()
	}
	if !errors.Is(err, goenv.ErrInvalidType) {
		t.FailNow()
	}

	os.Clearenv()
}

func TestGoenv_ErrCreatingResolver(t *testing.T) {
	var config int
	err := goenv.Load(config)
	if err == nil {
		t.FailNow()
	}
}
