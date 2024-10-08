package main

import (
	"github.com/phe-lab/ws/cmd/ws"
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
