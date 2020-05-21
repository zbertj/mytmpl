package logger

// 往终端里写日志

import (
	"fmt"
	"time"
)

// ConsoleLogger 定义ConsoleLogger类
type ConsoleLogger struct {
	Level LogLevel
}

func (c *ConsoleLogger) log(lv string, format string, a ...interface{}) {
	if c.enable(stringToLogLevel(lv)) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), lv, fileName, funcName, lineNo, msg)
	}
}

func (c *ConsoleLogger) enable(loglevel LogLevel) bool {
	return loglevel >= c.Level
}

// Close 顺从接口
func (c *ConsoleLogger) Close() {
	//do nothing
}

// Debug 输出打印
func (c *ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log("debug", format, a...)
}

// Trace 输出打印
func (c *ConsoleLogger) Trace(format string, a ...interface{}) {
	c.log("trace", format, a...)
}

// Info 输出打印
func (c *ConsoleLogger) Info(format string, a ...interface{}) {
	c.log("info", format, a...)
}

// Warning 输出打印
func (c *ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log("warning", format, a...)
}

// Error 输出打印
func (c *ConsoleLogger) Error(format string, a ...interface{}) {
	c.log("error", format, a...)
}

// Fatal 输出打印
func (c *ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log("fatal", format, a...)
}
