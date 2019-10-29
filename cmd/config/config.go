package config

import "time"

// Global is the single global store for all configuration values
var Global = struct {
	Name        string `flag:"-" env:"-" file:"-"`
	LogLevel    string
	ControlPort uint16
	SpecFile    string
	GracePeriod time.Duration
}{
	Name:        "flax",
	LogLevel:    "debug",
	ControlPort: 9999,
	SpecFile:    "flax.yaml",
	GracePeriod: 30 * time.Second,
}
