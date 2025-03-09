package logger

import (
	"compress/gzip"
	"gitee.com/quant1x/gox/api"
	"io"
	"os"
	"path/filepath"
	"time"
)

type DateWriter struct {
	logpath  string
	name     string
	dateType DateType
	num      int
	currDate string
	currFile *os.File
	openTime int64
}

func NewDateWriter(logpath, name string, dateType DateType, num int) *DateWriter {
	w := &DateWriter{
		logpath:  logpath,
		name:     name,
		num:      num,
		dateType: dateType,
	}
	w.currDate = w.getCurrDate()
	return w
}

type DateType uint8

// 压缩 使用gzip压缩成tar.gz
func gzipFile(source string) error {
	dest := source + ".gz"
	_ = os.Remove(dest)
	newfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer api.CloseQuietly(newfile)

	file, err := os.Open(source)
	if err != nil {
		return err
	}

	zw := gzip.NewWriter(newfile)

	filestat, err := file.Stat()
	if err != nil {
		return nil
	}

	zw.Name = filestat.Name()
	zw.ModTime = filestat.ModTime()
	_, err = io.Copy(zw, file)
	if err != nil {
		return nil
	}

	_ = zw.Flush()
	if err := zw.Close(); err != nil {
		return nil
	}
	return nil
}

func (w *DateWriter) Write(v []byte) {
	fullPath := filepath.Join(w.logpath, w.name+".log")
	unixTime := timeRotate.GetUnixTime()
	if w.currFile == nil || w.openTime+10 < unixTime {
		reopenFile(fullPath, &w.currFile, &w.openTime)
		w.currDate = w.getFileDate()
	}
	if w.currFile == nil {
		return
	}

	// 获取当前日期
	currDate := w.getCurrDate()
	if w.currDate != currDate {
		// 1. 文件改名
		sourceFile := fullPath
		destFile := filepath.Join(w.logpath, w.name+".log."+w.currDate)
		// 1.2 关闭当前文件
		api.CloseQuietly(w.currFile)
		w.currFile = nil
		// 1.3 删除目标文件, 预防无法改名
		err := os.Remove(destFile)
		// 1.4 文件改名, 原文件名增加后缀带日期
		err = os.Rename(sourceFile, destFile)
		if err != nil {
			// 改名失败, 源文件裁减长度为0
			_ = os.Truncate(sourceFile, 0)
		} else {
			// 改名成功
		}
		w.currDate = currDate

		_ = gzipFile(destFile)
		w.cleanOldLogs()
		reopenFile(fullPath, &w.currFile, &w.openTime)
		// 清理旧文件 [wangfeng on 2018/12/25 12:39]
		_ = os.Remove(destFile)
	}
	_, _ = w.currFile.Write(v)
}

func (w *DateWriter) NeedPrefix() bool {
	return true
}

func (w *DateWriter) getFormat() string {
	format := timeFmtDay
	if w.dateType == HOUR {
		format = timeFmtHour
	}
	return format
}

func (w *DateWriter) cleanOldLogs() {
	format := timeFmtDay
	duration := -time.Hour * 24
	if w.dateType == HOUR {
		format = timeFmtHour
		duration = -time.Hour
	}

	t := time.Now()
	t = t.Add(duration * time.Duration(w.num))
	for i := 0; i < 30; i++ {
		t = t.Add(duration)
		k := t.Format(format)
		fullPath := filepath.Join(w.logpath, w.name+".log."+k+".gz")
		if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
			_ = os.Remove(fullPath)
		}
	}
	return
}

func (w *DateWriter) getCurrDate() string {
	if w.dateType == HOUR {
		return timeRotate.GetDateHour()
	}
	return timeRotate.GetDateDay() // DAY
}

func (w *DateWriter) getFileDate() string {
	fi, err := w.currFile.Stat()
	if err == nil {
		return fi.ModTime().Format(w.getFormat())
	} else {
		return ""
	}
}
