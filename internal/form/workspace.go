package form

import (
	"path/filepath"
	"vscode-workspace-cli/internal/log"
	"vscode-workspace-cli/internal/utils"

	"github.com/charmbracelet/huh"
)

func ChooseWorkspace(basename string, path string) (string, error) {
	var selectedFile string

	workspaces, err := utils.FindWorkspaceFiles(path, basename)
	if err != nil {
		return "", err
	}

	if len(workspaces) == 0 {
		logEvent := log.Logger.Info().Str("workspacePath", path)
		if basename != "" {
			logEvent.Str("filename", basename)
		}
		logEvent.Msg("No workspace files found")
		return "", nil
	}

	if basename != "" && len(workspaces) == 1 {
		return workspaces[0], nil
	}

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a workspace to open: ").
				Filtering(true).
				Options(convertToOptions(workspaces, path)...).
				Value(&selectedFile),
		),
	).WithTheme(t).Run()

	if err != nil {
		return "", err
	}

	return selectedFile, nil
}

func convertToOptions(workspaces []string, path string) []huh.Option[string] {
	options := make([]huh.Option[string], len(workspaces))

	for i, workspace := range workspaces {
		relativePath, err := filepath.Rel(path, workspace)
		if err != nil {
			relativePath = workspace
		}
		options[i] = huh.NewOption(utils.ShortenPath(relativePath), workspace)
	}

	return options
}
