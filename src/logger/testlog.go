package main

import "logger"

func log() {
	log := logger.NewLogger("f", "warning")
	log.Error("接口 文件 输出x")
}

func consolelog() {
	log := logger.NewConsoleLogger("warning")
	log.Debug("调试！%s\n", "wdf")
	log.Info("提示！%s\n", "wdf")
	log.Warning("警告！%s\n", "wdf")
	log.Error("出错！%s\n", "wdf")
	log.Fatal("致命错误！%s\n", "wdf")
}

func filelog() {
	log := logger.NewFileLogger("Info", "./", "log.log")
	// for {
	log.Debug("调试！")
	log.Info("提示！")
	log.Warning("警告！")
	log.Error("出错！")
	log.Fatal("致命错误！")
	// }
}
