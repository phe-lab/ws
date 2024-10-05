package form

import (
	"path/filepath"
	"strings"
	"vscode-workspace-cli/internal/log"
	"vscode-workspace-cli/internal/utils"

	"github.com/charmbracelet/huh"
)

func shortenPath(path string) string {
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

func convertToOptions(workspaces []string, path string) []huh.Option[string] {
	options := make([]huh.Option[string], len(workspaces))

	for i, workspace := range workspaces {
		relativePath, err := filepath.Rel(path, workspace)
		if err != nil {
			relativePath = workspace
		}
		options[i] = huh.NewOption("./"+shortenPath(relativePath), workspace)
	}

	return options
}

func ChooseWorkspace(path string) (string, error) {
	var selectedFile string
	var logger = log.GetLogger()

	workspaces, err := utils.FindWorkspaceFiles(path)
	if err != nil {
		return selectedFile, err
	}

	if len(workspaces) == 0 {
		logger.Info().Str("workspacePath", path).Msg("The workspace directory is currently empty")
		return selectedFile, nil
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a workspace to open: ").
				Options(convertToOptions(workspaces, path)...).
				Value(&selectedFile),
		),
	)

	if err = form.Run(); err != nil {
		if err != huh.ErrUserAborted {
			return selectedFile, err
		}
	}

	return selectedFile, nil
}
