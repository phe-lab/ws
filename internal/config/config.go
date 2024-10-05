package config

import (
	"os/user"
	"path/filepath"
)

const DEFAULT_WORKSPACE_DIR = ".code-workspaces"

type Config struct {
	Debug         bool
	WorkspacePath string
}

func DefaultWorkspacePath() string {
	currentUser, err := user.Current()
	if err != nil {
		return DEFAULT_WORKSPACE_DIR
	}

	return filepath.Join(currentUser.HomeDir, DEFAULT_WORKSPACE_DIR)
}

func (c *Config) GetWorkspacePath() string {
	if c.WorkspacePath != "" {
		return c.WorkspacePath
	}

	return DefaultWorkspacePath()
}
