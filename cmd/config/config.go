package config

import (
	"github.com/moorara/goto/config"
)

const (
	defaultLogLevel        = "info"
	defaultControlPort     = ":9000"
	defaultJaegerAgentAddr = "localhost:6831"
	defaultJaegerLogSpans  = false
)

// Config defines the configuration values
var Config = struct {
	LogLevel        string
	ControlPort     string
	JaegerAgentAddr string
	JaegerLogSpans  bool
}{
	LogLevel:        defaultLogLevel,
	ControlPort:     defaultControlPort,
	JaegerAgentAddr: defaultJaegerAgentAddr,
	JaegerLogSpans:  defaultJaegerLogSpans,
}

func init() {
	config.Pick(&Config)
}
