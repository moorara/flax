package main

import (
	"flag"
	"os"

	"github.com/moorara/flax/cmd/config"
	"github.com/moorara/flax/cmd/server"
	"github.com/moorara/flax/internal/service"
	"github.com/moorara/flax/internal/spec"
	"github.com/moorara/flax/version"
	"github.com/moorara/konfig"
	"github.com/moorara/log"
)

const (
	specErr = 10
)

func main() {
	// Reading configuration values
	_ = konfig.Pick(&config.Global)

	// Populating flags
	flag.Parse()

	// Create an instance logger
	logger := log.NewKit(log.Options{
		Name:  config.Global.Name,
		Level: config.Global.LogLevel,
	})

	// Log binary information
	logger = logger.With(
		"bin.version", version.Version,
		"bin.commit", version.Commit,
		"bin.branch", version.Branch,
		"bin.goVersion", version.GoVersion,
		"bin.buildTool", version.BuildTool,
		"bin.buildTime", version.BuildTime,
	)

	// Reading spec file
	s, err := spec.ReadSpec(config.Global.SpecFile)
	if err != nil {
		logger.Errorf("error while reading spec file: %s", err)
		os.Exit(specErr)
	}

	// Set up mock service
	mockService := service.NewMockService(logger)
	for i := range s.HTTPMocks {
		mockService.Add(&s.HTTPMocks[i])
	}
	for i := range s.RESTMocks {
		mockService.Add(&s.RESTMocks[i])
	}

	// Set up api server
	router := mockService.Router()
	apiServer := server.NewAPIServer(logger, s.Config.HTTPPort, router)
	apiServer.Start()
}
