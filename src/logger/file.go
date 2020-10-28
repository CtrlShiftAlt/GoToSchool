package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 文件操作类
type FileLogger struct {
	Level      Level
	fileObj    *os.File
	errFileObj *os.File
	filePath   string
	fileName   string
	maxSize    int64
}

// NewFileLogger 实例化文件操作类
func NewFileLogger(levelStr string, filePath string, fileName string) (f *FileLogger) {
	level := strToLevel(levelStr)
	f = &FileLogger{
		Level:    level,
		filePath: filePath,
		fileName: fileName,
		maxSize:  10 * 1024 * 1024,
	}
	f.openFile()
	return
}

// // SaveFilePath 设置保存文件路径 默认 "./"
// func (f *FileLogger) SaveFilePath(filePath string) {
// 	f.filePath = filePath
// }

// // SaveFileName 设置保存文件路径 默认 "log.log"
// func (f *FileLogger) SaveFileName(fileName string) {
// 	f.filePath = fileName
// }

// openFile 打开文件
func (f *FileLogger) openFile() {
	// 打开日志文件
	pathStr := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(pathStr, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file fail, err: %v\n", err)
		return
	}
	// 打开错误日志文件
	errPathStr := path.Join(f.filePath, "err."+f.fileName)
	errFileObj, err := os.OpenFile(errPathStr, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err file fail, err: %v\n", err)
		return
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return
}

func (f *FileLogger) splitFile(fail *os.File) (*os.File, error) {
	// 判断文件大小如果文件达到maxSize则切割
	fileInfo, err := fail.Stat()
	if err != nil {
		fmt.Printf("get fileInfo failed, err: %v\n", err)
	}
	if fileInfo.Size() > f.maxSize {
		pathStr := path.Join(f.filePath, fileInfo.Name())
		timeStr := time.Now().Format("20060102150405000")
		// 关闭文件
		fail.Close()
		// 复制文件到bak
		err = os.Rename(pathStr, pathStr+timeStr)
		if err != nil {
			fmt.Printf("file Rename failed,err: %v\n", err)
			return nil, err
		}
		// 打开新文件
		fail, err = os.OpenFile(pathStr, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("open file fail, err: %v\n", err)
			return nil, err
		}
	}
	return fail, nil
}

// logOrNot 验证输出等级 符合的输出
func (f *FileLogger) logOrNot(level Level) bool {
	return level >= f.Level
}

// log 写入文件操作
func (f *FileLogger) log(level Level, format string, a ...interface{}) {
	// 验证等级 符合的输出
	if f.logOrNot(level) {
		msg := fmt.Sprintf(format, a...)
		levelStr := levelToStr(level)
		// 执行处的文件、函数、行号信息
		funcname, filename, line := getFileInfo(3)
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		// 写入文件
		fileObj, err := f.splitFile(f.fileObj)
		if err != nil {
			fmt.Printf("splitFile file fail, err: %v\n", err)
			return
		}
		f.fileObj = fileObj
		fmt.Fprintf(f.fileObj, "[%s] [%s|%s|%d] [%s] %s\n", timeStr, funcname, filename, line, levelStr, msg)
		if level >= ERROR {
			errFileObj, err := f.splitFile(f.errFileObj)
			if err != nil {
				fmt.Printf("splitFile file fail, err: %v\n", err)
				return
			}
			f.errFileObj = errFileObj
			fmt.Fprintf(f.errFileObj, "[%s] [%s|%s|%d] [%s] %s\n", timeStr, funcname, filename, line, levelStr, msg)
		}
	}
}

// Debug ...
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info ...
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning ...
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error ...
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal ...
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
