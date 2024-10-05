package ws

import (
	"fmt"
	"os"
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
	Use:     "ws",
	Version: "0.1.0",
	Short:   "A simple CLI tool to quickly open VSCode Workspace",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig(debug)
		logger := log.InitLogger(debug)
		workspaceDir := cfg.GetWorkspacePath()

		if err := ensureWorkspaceDir(workspaceDir); err != nil {
			if err != exception.ErrUserAborted {
				logger.Error().Msg(err.Error())
				logger.Debug().Msg("application stopped")
				os.Exit(1)
			}
			return
		}

		logger.Debug().Msg("Application is running in debug mode")
		logger.Debug().Str("envVar", "VSCODE_WS_PATH").Msg("Using the default workspace directory as the environment variable was not found")
		logger.Debug().Str("workspacePath", workspaceDir).Msg("")

		workspace, err := form.ChooseWorkspace(workspaceDir)
		if err != nil {
			logger.Error().Msg(err.Error())
			os.Exit(1)
		}

		if workspace != "" {
			logger.Debug().Str("file", workspace).Msg("Selected workspace")
			logger.Debug().Str("file", workspace).Msg("Opening workspace")
		}
	},
}

func ensureWorkspaceDir(path string) error {
	if err := utils.ValidateWorkspacePath(path); err != nil {
		if err == exception.ErrNotExist {
			yes, err := form.ConfirmCreateDirectory(path)
			if err != nil || !yes {
				if err == huh.ErrUserAborted || !yes {
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
