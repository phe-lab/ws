package form

import (
	"path/filepath"
	"vscode-workspace-cli/internal/log"
	"vscode-workspace-cli/internal/utils"

	"github.com/charmbracelet/huh"
)

func ChooseWorkspace(path string) (string, error) {
	var selectedFile string

	workspaces, err := utils.FindWorkspaceFiles(path)
	if err != nil {
		return "", err
	}

	if len(workspaces) == 0 {
		log.Logger.Info().Str("workspacePath", path).Msg("The workspace directory is currently empty")
		return "", nil
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
