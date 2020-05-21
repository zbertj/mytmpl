package logger

/*
功能描述：
自定义的分级别logger，支持输出终端和文件（按照文件大小分割）

调用方式：
log = logger.NewLog("Warning", "f+", "./", "test.log", 5*1024)
log = logger.NewLog("Warning", "", "", "", 0)

*/

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

const (
	// UNKNOWN : 0
	UNKNOWN LogLevel = iota
	// DEBUG : 1
	DEBUG
	// TRACE : 2
	TRACE
	// INFO : 3
	INFO
	// WARNING 4
	WARNING
	// ERROR 5
	ERROR
	// FATAL 6
	FATAL
)

// LogLevel 自定义uint6类型
type LogLevel uint16

// Logger 定义Logger接口
type Logger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Info(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	Close()
}

// stringToLogLevel 将传进来的string类型的level 转化为LogLevel，也就是uint16
func stringToLogLevel(level string) LogLevel {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return DEBUG

	case "trace":
		return TRACE

	case "info":
		return INFO

	case "warning":
		return WARNING

	case "error":
		return ERROR

	case "fatal":
		return FATAL

	default:
		return UNKNOWN
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed!")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

// NewLog 创建并返回一个Logger实例
func NewLog(level, mode, fpath, fname string, fileSize int64) Logger {
	loglevel := stringToLogLevel(level)
	fileLogger := &FileLogger{
		Level:       loglevel,
		Mode:        mode,
		FPath:       fpath,
		FName:       fname,
		maxFileSize: fileSize,
	}

	fullFileName := path.Join(fpath, fname)
	if mode == "f" {
		fileObj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		fileLogger.fileObj = fileObj
		return fileLogger
	}
	if mode == "f+" {
		fileObj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		errFileObj, err := os.OpenFile(fullFileName+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		fileLogger.fileObj = fileObj
		fileLogger.errFileObj = errFileObj
		return fileLogger
	}

	return &ConsoleLogger{
		Level: loglevel,
	}

}
