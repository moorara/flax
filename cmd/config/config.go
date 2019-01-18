package config

import (
	"github.com/moorara/goto/config"
)

const (
	defaultName            = "flax"
	defaultLogLevel        = "info"
	defaultControlPort     = ":9999"
	defaultJaegerAgentAddr = "localhost:6831"
	defaultJaegerLogSpans  = false
)

// Config defines the configuration values
var Config = struct {
	Name            string `flag:"-" env:"-" file:"-"`
	LogLevel        string
	ControlPort     string
	JaegerAgentAddr string
	JaegerLogSpans  bool
}{
	Name:            defaultName,
	LogLevel:        defaultLogLevel,
	ControlPort:     defaultControlPort,
	JaegerAgentAddr: defaultJaegerAgentAddr,
	JaegerLogSpans:  defaultJaegerLogSpans,
}

func init() {
	config.Pick(&Config)
}
