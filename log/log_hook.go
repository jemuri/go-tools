package log

import (
	"fmt"
	"github.com/jemuri/go-tools/config"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
	"time"
)

var file *os.File
var onceDo sync.Once

func obtainFile() *os.File {
	if file == nil {
		// 读取一次文件 提高性能
		onceDo.Do(func() {
			name := fmt.Sprintf(config.CertainString("log/file"), time.Now().Format("2006_01_02"))
			var err error
			if file, err = os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0755); err != nil {
				fmt.Println("打开日志文件出错: ", err)
			}
		})
	}

	return file
}

// FileHook to send logs via syslog.
type FileHook struct {
	Writer  *os.File
	LogPath string
}

// NewFileHook 日志文件
func NewFileHook(logPath string) *FileHook {
	if file == nil {
		file = obtainFile()
	}
	return &FileHook{file, logPath}
}

func (hook *FileHook) Fire(entry *logrus.Entry) error {
	writer := io.MultiWriter(hook.Writer, os.Stdout)
	entry.Logger.SetOutput(writer)

	return nil
}

func (hook *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
