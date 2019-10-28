package config

// Global is the single global store for all configuration values
var Global = struct {
	Name        string `flag:"-" env:"-" file:"-"`
	LogLevel    string
	ControlPort string
}{
	Name:        "flax",
	LogLevel:    "info",
	ControlPort: ":9999",
}
