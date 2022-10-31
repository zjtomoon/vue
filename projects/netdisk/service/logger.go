package service

import "github.com/yddeng/utils/log"

var Logger *log.Logger

func InitLogger(basePath string, fileName string) *log.Logger {
	Logger = log.NewLogger(basePath, fileName, 1024*1024*4)
	Logger.Debugf("%s logger init", fileName)
	return Logger
}
