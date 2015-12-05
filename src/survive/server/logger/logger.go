package logger

import (
	log "github.com/cihub/seelog"
	"fmt"
	"errors"
)

type Logger interface {
	log.LoggerInterface
}

//私有属性
var defaultLogger  Logger

//package 初始化
func init() {
	//读取配置文件，创建logger
	newLogger,err := log.LoggerFromConfigAsFile("conf/log.xml")
	if err != nil{
		fmt.Println("logger initial failed!")
		return
	}
	log.ReplaceLogger(newLogger)
	defaultLogger = log.Current
	defaultLogger.Info("logger initial success!")
}
/**
	在主程序的defer里调用此方法，可保证日志始终记录
 */
func FlushAllLogs() {
	log.Flush()
}
/**
	获取全局的logger
 */
func GetLogger() Logger {
	if defaultLogger == nil{
		panic(errors.New("logger initial failed! can not get logger"))
	}
	return defaultLogger
}