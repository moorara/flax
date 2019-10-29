package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moorara/konfig"
	"github.com/moorara/observe/log"

	"github.com/moorara/flax/cmd/config"
	"github.com/moorara/flax/cmd/server"
	"github.com/moorara/flax/cmd/version"
	"github.com/moorara/flax/internal/service"
	"github.com/moorara/flax/internal/spec"
)

const specErr = 10

func main() {
	// Reading configuration values
	_ = konfig.Pick(&config.Global)

	// Populating flags
	flag.Parse()

	// Create an instance logger
	logger := log.NewLogger(log.Options{
		Name:  config.Global.Name,
		Level: config.Global.LogLevel,
	})

	// Log binary information
	logger = logger.With(
		"bin.version", version.Version,
		"bin.revision", version.Revision,
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

	mockService := service.NewMockService(logger)

	// FIXME:
	for _, m := range s.HTTPMocks {
		mockService.Add(&m)
	}

	// FIXME:
	for _, m := range s.RESTMocks {
		mockService.Add(&m)
	}

	port := fmt.Sprintf(":%d", s.Config.HTTPPort)
	router := mockService.Router()
	httpMockServer := server.NewHTTPMockServer(logger, port, router)
	httpMockServer.Start()
}
