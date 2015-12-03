package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"survive/server/conf"
	"survive/server/game"
	"survive/server/gate"
	"survive/server/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
