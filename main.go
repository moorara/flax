package main

import (
	"github.com/moorara/flax/cmd/config"
	"github.com/moorara/flax/cmd/version"
	"github.com/moorara/flax/pkg/log"
)

func main() {
	logger := log.NewLogger("flax", config.Config.LogLevel)

	logger.Info(
		"version", version.Version,
		"revision", version.Revision,
		"branch", version.Branch,
		"buildTime", version.BuildTime,
		"message", "Flax started.",
	)
}
