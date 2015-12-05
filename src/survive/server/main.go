package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"survive/server/conf"
	"survive/server/modules/game"
	"survive/server/modules/gate"
	"survive/server/modules/login"
	"survive/server/logger"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath

	logger.GetLogger().Infof("prepare Run:")
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
	logger.GetLogger().Infof("prepare Exit")
}
