package utils

import (
	"os"
	"os/exec"

	"github.com/phe-lab/ws/internal/log"
)

func OpenWorkspace(workspace string) {
	log.Logger.Info().Str("file", workspace).Msg("Opening workspace")
	cmd := exec.Command("code", workspace)
	if err := cmd.Run(); err != nil {
		log.Logger.Error().Msg(err.Error())
		os.Exit(1)
	}
	log.Logger.Info().Str("file", workspace).Msg("Workspace is opened")
}
