package utils

import (
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"vscode-workspace-cli/internal/exception"
)

func FindWorkspaceFiles(directory string) ([]string, error) {
	var workspaces []string

	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".code-workspace" {
			workspaces = append(workspaces, path)
		}

		return nil
	})

	return workspaces, err
}

func ValidateWorkspacePath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return exception.ErrNotExist
		}
		return exception.ErrUnhandled
	}

	if !info.IsDir() {
		return exception.ErrNotDirectory
	}

	return nil
}

func ShortenPath(path string) string {
	path = filepath.ToSlash(path)
	parts := strings.Split(path, "/")

	// shorten name of parrent directories:
	for i := 0; i < len(parts)-2; i++ {
		if parts[i] != "" {
			parts[i] = string(parts[i][0])
		}
	}

	// remove file extension .code-workspace:
	lastPart := parts[len(parts)-1]
	parts[len(parts)-1] = strings.TrimSuffix(lastPart, ".code-workspace")

	return strings.Join(parts, "/")
}

func NormalizePath(path string) (string, error) {
	usr, _ := user.Current()
	if len(path) > 1 && path[:2] == "~/" {
		// remove ~/:
		path = filepath.Join(usr.HomeDir, path[2:])
	}

	// replace environment variables:
	path = os.ExpandEnv(path)

	return filepath.Abs(filepath.Clean(path))
}
