package ws

import (
	"fmt"
	"os"
	"os/exec"
	"vscode-workspace-cli/internal/config"
	"vscode-workspace-cli/internal/exception"
	"vscode-workspace-cli/internal/form"
	"vscode-workspace-cli/internal/log"
	"vscode-workspace-cli/internal/utils"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:   "ws",
	Short: "A simple CLI tool to quickly open VSCode Workspace",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig(debug)
		logger := log.InitLogger(debug)
		workspaceDir := cfg.GetWorkspacePath()

		if err := ensureWorkspaceDir(workspaceDir); err != nil {
			logger.Error().Msg(err.Error())
			if err != exception.ErrUserAborted {
				os.Exit(1)
			}
			return
		}

		logger.Debug().Msg("Application is running in debug mode")
		logger.Debug().Str("envVar", "VSCODE_WS_PATH").Msg("Using the default workspace directory as the environment variable was not found")
		logger.Debug().Str("workspacePath", workspaceDir).Msg("")

		workspace, err := form.ChooseWorkspace(workspaceDir)
		if err != nil && err != huh.ErrUserAborted {
			logger.Error().Msg(err.Error())
			os.Exit(1)
		}

		if workspace != "" {
			logger.Debug().Str("file", workspace).Msg("Selected workspace")
			logger.Info().Str("file", workspace).Msg("Opening workspace")
			cmd := exec.Command("code", workspace)
			if err := cmd.Run(); err != nil {
				logger.Error().Msg(err.Error())
				os.Exit(1)
			}
		}
	},
}

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s from Git SHA %s)", version, date, commit)
}

func ensureWorkspaceDir(path string) error {
	if err := utils.ValidateWorkspacePath(path); err != nil {
		if err == exception.ErrNotExist {
			confirm, err := form.ConfirmCreateDirectory(path)
			if err != nil || !confirm {
				if err == huh.ErrUserAborted || !confirm {
					return exception.ErrUserAborted
				}
				return err
			}

			if err = os.MkdirAll(path, 0755); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug mode")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing ws '%s'\n", err)
		os.Exit(1)
	}
}
