package main

import (
	"vscode-workspace-cli/cmd/ws"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	ws.SetVersionInfo(version, commit, date)
	ws.Execute()
}
