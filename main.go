package main

import (
	"github.com/moorara/flax/cmd/config"
	"github.com/moorara/flax/cmd/version"
	"github.com/moorara/goto/log"
)

func main() {
	// Create logger
	opts := log.Options{Name: config.Config.Name, Level: config.Config.LogLevel}
	logger := log.NewLogger(opts)
	logger = logger.With(
		config.Config.Name, map[string]string{
			"version":   version.Version,
			"revision":  version.Revision,
			"branch":    version.Branch,
			"goVersion": version.GoVersion,
			"buildTool": version.BuildTool,
			"buildTime": version.BuildTime,
		},
	)

	logger.Info("message", "Hello, World!")
}
