package utils

import (
	"io/fs"
	"os"
	"path/filepath"
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
