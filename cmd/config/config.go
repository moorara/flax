package config

import (
	"github.com/moorara/goto/config"
)

const (
	defaultLogLevel        = "info"
	defaultHTTPPort        = ":8080"
	defaultHTTPSPort       = ":8443"
	defaultJaegerAgentAddr = "localhost:6831"
	defaultJaegerLogSpans  = false
)

// Config defines the configuration values
var Config = struct {
	LogLevel        string
	HTTPPort        string
	HTTPSPort       string
	JaegerAgentAddr string
	JaegerLogSpans  bool
}{
	LogLevel:        defaultLogLevel,
	HTTPPort:        defaultHTTPPort,
	HTTPSPort:       defaultHTTPSPort,
	JaegerAgentAddr: defaultJaegerAgentAddr,
	JaegerLogSpans:  defaultJaegerLogSpans,
}

func init() {
	config.Pick(&Config)
}
