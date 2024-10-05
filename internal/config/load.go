package config

import (
	"os"
	"sync"
)

var (
	configInstance *Config
	once           sync.Once
)

func LoadConfig(debug bool) *Config {
	once.Do(func() {
		workspacePath := os.Getenv("VSCODE_WS_PATH")

		configInstance = &Config{
			WorkspacePath: workspacePath,
			Debug:         debug || os.Getenv("VSCODE_WS_DEBUG") == "true",
		}
	})

	return configInstance
}
