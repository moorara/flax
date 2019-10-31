package main

import (
	"flag"
	"os"

	"github.com/moorara/konfig"
	"github.com/moorara/observe/log"
	"github.com/moorara/observe/xhttp"

	"github.com/moorara/flax/cmd/config"
	"github.com/moorara/flax/cmd/server"
	"github.com/moorara/flax/cmd/version"
	"github.com/moorara/flax/internal/service"
	"github.com/moorara/flax/internal/spec"
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

	// Set up mock service
	mockService := service.NewMockService(logger)
	for i := range s.HTTPMocks {
		mockService.Add(&s.HTTPMocks[i])
	}
	for i := range s.RESTMocks {
		mockService.Add(&s.RESTMocks[i])
	}

	// Set up http middleware
	mid := xhttp.NewServerMiddleware(
		xhttp.ServerLogging(logger),
	)

	// Set up api server
	router := mockService.Router()
	handler := mid.Logging(router.ServeHTTP)
	apiServer := server.NewAPIServer(logger, s.Config.HTTPPort, handler)
	apiServer.Start()
}
