package logger

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/quant1x/gox/api"
)

const (
	DAY DateType = iota
	HOUR
)

type LogWriter interface {
	Write(v []byte)
	NeedPrefix() bool
}

// 重新打开文件
func reopenFile(name string, currFile **os.File, openTime *int64) {
	*openTime = timeRotate.GetUnixTime()
	if *currFile != nil {
		_ = (*currFile).Close()
	}
	of, err := os.OpenFile(name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, api.CACHE_FILE_MODE)
	if err == nil {
		*currFile = of
	} else {
		fmt.Println("open log file error", err)
	}
}

// 压缩 使用gzip压缩成gz
func gzipFile(source string) (err error) {
	// 打开源文件
	srcFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("无法打开源文件: %w", err)
	}
	defer api.CloseQuietly(srcFile)

	// 获取文件元数据
	fileStat, err := srcFile.Stat()
	if err != nil {
		return fmt.Errorf("获取文件信息失败: %w", err)
	}

	// 创建目标文件
	dest := source + ".gz"
	gzFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("创建压缩文件失败: %w", err)
	}
	defer api.CloseQuietly(gzFile)

	// 初始化gzip写入器（带错误检查）
	gzWriter, err := gzip.NewWriterLevel(gzFile, gzip.DefaultCompression)
	if err != nil {
		return fmt.Errorf("初始化压缩器失败: %w", err)
	}
	defer api.CloseQuietly(gzWriter)

	// 设置压缩文件头信息
	gzWriter.Name = filepath.Base(source) // 确保文件名不含路径
	gzWriter.ModTime = fileStat.ModTime() // 保留原始修改时间

	// 创建缓冲写入层（32KB缓冲区）
	bufWriter := bufio.NewWriterSize(gzWriter, 32*1024)
	defer func() {
		// 确保缓冲区数据刷新，并捕获相关错误
		if flushErr := bufWriter.Flush(); flushErr != nil && err == nil {
			err = fmt.Errorf("缓冲区刷新失败: %w", flushErr)
		}
	}()

	// 执行压缩操作
	if _, err = io.Copy(bufWriter, srcFile); err != nil {
		return fmt.Errorf("压缩过程失败: %w", err)
	}

	return nil
}
