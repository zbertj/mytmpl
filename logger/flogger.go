package logger

// 向文件里写日志

import (
	"fmt"
	"os"
	"path"
	"time"
)

//FileLogger 定义FileLogger类
type FileLogger struct {
	Level       LogLevel
	Mode        string
	FPath       string
	FName       string
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
}

// writeFile
func (f *FileLogger) splitFile(fp *os.File) (*os.File, bool) {
	fileInfo, err := fp.Stat()
	if err != nil {
		fmt.Printf("get file info failed %v\n", err.Error())
		return nil, false
	}
	// 1.判断是否切割日志文件
	if fileInfo.Size() >= f.maxFileSize {
		fileName := fileInfo.Name()

		//2. 关闭当前文件
		fp.Close()

		//3. 备份当前文件
		timeStr := time.Now().Format("20060102150405000")
		logName := path.Join(f.FPath, fileName)
		bakLogName := fmt.Sprintf("%s.bak%s", logName, timeStr)
		err = os.Rename(logName, bakLogName)
		if err != nil {
			fmt.Println(err.Error())
		}
		// 4.用来的名字重新打开一个文件
		newfp, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("new open file failed %v\n", err.Error())
			return nil, false
		}
		return newfp, true
	}
	return nil, true
}

func (f *FileLogger) log(lv string, format string, a ...interface{}) {
	if f.enable(stringToLogLevel(lv)) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		// 内部判断是否切割，如果是就执行切割
		newFileObj, ok := f.splitFile(f.fileObj)
		// 判断splitFile是否出现错误
		if ok {
			// 有新的文件对象就将新的文件对象赋值给类的fileObj
			if newFileObj != nil {
				f.fileObj = newFileObj
			}
			fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), lv, fileName, funcName, lineNo, msg)
		}
		if f.Mode == "f+" {
			if stringToLogLevel(lv) >= ERROR {
				newErrFileObj, ok := f.splitFile(f.errFileObj)
				if ok {
					if newErrFileObj != nil {
						f.errFileObj = newErrFileObj
					}
					fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), lv, fileName, funcName, lineNo, msg)
				}
			}
		}
	}
}

// enable 判断是否大于等于限制等级
func (f *FileLogger) enable(loglevel LogLevel) bool {
	return loglevel >= f.Level
}

// Close 关闭日志文件对象
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

// Debug 输出打印
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log("debug", format, a...)
}

// Trace 输出打印
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log("trace", format, a...)
}

// Info 输出打印
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log("info", format, a...)
}

// Warning 输出打印
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log("warning", format, a...)
}

// Error 输出打印
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log("error", format, a...)
}

// Fatal 输出打印
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log("fatal", format, a...)
}
