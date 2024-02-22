package logger

import (
	"fmt"
	"ops_client/pkg/util"

	"os"
)

type Logger struct {
}

var logger *Logger

func Log() *Logger {
	return logger
}

// 日志写入函数
func (logvar *Logger) Panic(service string, handler string, m ...any) {
	msg := fmt.Sprint("[Panic] "+"["+handler+"] ", fmt.Sprint(m...))
	util.CommonLog(service, msg)
	os.Exit(0)
}

func (logvar *Logger) Error(service string, handler string, m ...any) {
	msg := fmt.Sprint("[Error] "+"["+handler+"] ", fmt.Sprint(m...))
	util.CommonLog(service, msg)
}

func (logvar *Logger) Warning(service string, handler string, m ...any) {
	msg := fmt.Sprint("[Warning] "+"["+handler+"] ", fmt.Sprint(m...))
	util.CommonLog(service, msg)
}

func (logvar *Logger) Info(service string, handler string, m ...any) {
	msg := fmt.Sprint("[Info] "+"["+handler+"] ", fmt.Sprint(m...))
	util.CommonLog(service, msg)
}
