package main

import (
	"github.com/moorara/chia/cmd/config"
	"github.com/moorara/chia/cmd/version"
	"github.com/moorara/chia/pkg/log"
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
