package common

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// myWriter 自定义writer的对象
var myWriter dateFileWriter

// targetFile 当前要写入日志的文件
var targetFile *os.File

// dateFileWriter 自定义一个writer专门用于写日志
type dateFileWriter struct {
	io.Writer
}

// Write 为自定义writer实现Write接口
func (b *dateFileWriter) Write(p []byte) (n int, err error) {
	return targetFile.Write(p)
}

// RefreshLogFileUsage 刷新指向的日志文件
func RefreshLogFileUsage() {
	fileName := "log/" + time.Now().Format("2006_01_02") + ".log"
	tryFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		err := os.WriteFile(fileName, []byte(""), 0777)
		if err != nil {
			return
		}
		targetFile, _ = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	} else {
		targetFile = tryFile
	}
}

// LogWriter 返回自定义的 writer
func LogWriter() io.Writer {
	RefreshLogFileUsage()
	return &myWriter
}

func Log(format string, values ...any) {
	now := time.Now().Format("2006/01/02 - 15:04:05")
	f := fmt.Sprintf("[LOG] %s %s\n", now, format)
	_, err := fmt.Fprintf(gin.DefaultWriter, f, values...)
	if err != nil {
		return
	}
}
