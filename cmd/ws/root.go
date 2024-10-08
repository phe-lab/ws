package ws

import (
	"fmt"
	"os"

	"github.com/phe-lab/ws/internal/config"
	"github.com/phe-lab/ws/internal/exception"
	"github.com/phe-lab/ws/internal/form"
	"github.com/phe-lab/ws/internal/log"
	"github.com/phe-lab/ws/internal/utils"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var debug bool
var customUsage string = `Usage:
  {{.UseLine}}

Arguments:
  [filename]  Optional filename to be opened

Flags:
{{.Flags.FlagUsages | trimTrailingWhitespaces}}

Examples:
  # List the workspaces
  ws

  # Open the workspace with the filename "simple-scrollspy.code-workspace"
  ws simple-scrollspy

  # Set the logging level to "debug"
  ws --debug
`

var rootCmd = &cobra.Command{
	Use:   "ws [filename]",
	Short: "A simple CLI tool to quickly open VSCode Workspace",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig(debug)
		logger := log.InitLogger(debug)
		workspaceDir := cfg.GetWorkspacePath()

		filename := ""
		if len(args) > 0 {
			filename = fmt.Sprintf("%s.code-workspace", args[0])
			logger.Debug().Str("filename", filename).Msg("Set the target workspace matching the file")
		}

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

		workspace, err := form.ChooseWorkspace(filename, workspaceDir)
		if err != nil && err != huh.ErrUserAborted {
			logger.Error().Msg(err.Error())
			os.Exit(1)
		}

		if workspace != "" {
			log.Logger.Debug().Str("file", workspace).Msg("Selected workspace")
			utils.OpenWorkspace(workspace)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug mode")
	rootCmd.SetUsageTemplate(customUsage)
}

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s from Git SHA %s)", version, date, commit)
}

func ensureWorkspaceDir(path string) error {
	log.Logger.Debug().Msg("Validate workspace path")
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing ws '%s'\n", err)
		os.Exit(1)
	}
}
