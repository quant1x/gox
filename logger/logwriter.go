package logger

import (
	"compress/gzip"
	"fmt"
	"gitee.com/quant1x/gox/api"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	DAY DateType = iota
	HOUR
)

type LogWriter interface {
	Write(v []byte)
	NeedPrefix() bool
}

type ConsoleWriter struct {
}

type RollFileWriter struct {
	logpath  string
	name     string
	num      int
	size     int64
	currSize int64
	currFile *os.File
	openTime int64
}

type DateWriter struct {
	logpath  string
	name     string
	dateType DateType
	num      int
	currDate string
	currFile *os.File
	openTime int64
}

type HourWriter struct {
}

type DateType uint8

func reOpenFile(path string, currFile **os.File, openTime *int64) {
	*openTime = currUnixTime
	if *currFile != nil {
		_ = (*currFile).Close()
	}
	of, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		*currFile = of
	} else {
		fmt.Println("open log file error", err)
	}
}

func (w *ConsoleWriter) Write(v []byte) {
	_, _ = os.Stdout.Write(v)
}

func (w *ConsoleWriter) NeedPrefix() bool {
	return true
}

func (w *RollFileWriter) Write(v []byte) {
	if w.currFile == nil || w.openTime+10 < currUnixTime {
		fullPath := filepath.Join(w.logpath, w.name+".log")
		reOpenFile(fullPath, &w.currFile, &w.openTime)
	}
	if w.currFile == nil {
		return
	}
	n, _ := w.currFile.Write(v)
	w.currSize += int64(n)
	if w.currSize >= w.size {
		w.currSize = 0
		for i := w.num - 1; i >= 1; i-- {
			var n1, n2 string
			if i > 1 {
				n1 = strconv.Itoa(i - 1)
			}
			n2 = strconv.Itoa(i)
			p1 := filepath.Join(w.logpath, w.name+n1+".log")
			p2 := filepath.Join(w.logpath, w.name+n2+".log")
			if _, err := os.Stat(p1); !os.IsNotExist(err) {
				_ = os.Rename(p1, p2)
			}
		}
		fullPath := filepath.Join(w.logpath, w.name+".log")
		reOpenFile(fullPath, &w.currFile, &w.openTime)
	}
}

func NewRollFileWriter(logpath, name string, num, sizeMB int) *RollFileWriter {
	w := &RollFileWriter{
		logpath: logpath,
		name:    name,
		num:     num,
		size:    int64(sizeMB) * 1024 * 1024,
	}
	fullPath := filepath.Join(logpath, name+".log")
	st, _ := os.Stat(fullPath)
	if st != nil {
		w.currSize = st.Size()
	}
	return w
}

func (w *RollFileWriter) NeedPrefix() bool {
	return true
}

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
	//isNewFile := false
	if w.currFile == nil || w.openTime+10 < currUnixTime {
		reOpenFile(fullPath, &w.currFile, &w.openTime)
		w.currDate = w.getFileDate()
	}
	if w.currFile == nil {
		return
	}

	currDate := w.getCurrDate()
	if w.currDate != currDate {
		// 文件改名
		sourceFile := fullPath
		destFile := filepath.Join(w.logpath, w.name+".log."+w.currDate)
		// 删除已有的目标文件
		w.currFile.Close()
		w.currFile = nil
		err := os.Remove(destFile)
		err = os.Rename(sourceFile, destFile)
		if err != nil {
			// 改名失败
		} else {
			// 改名成功
		}
		w.currDate = currDate

		_ = gzipFile(destFile)
		w.cleanOldLogs()
		reOpenFile(fullPath, &w.currFile, &w.openTime)
		// 清理旧文件 [wangfeng on 2018/12/25 12:39]
		_ = os.Remove(destFile)
	}
	_, _ = w.currFile.Write(v)
}

func (w *DateWriter) NeedPrefix() bool {
	return true
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

func (w *DateWriter) getFotmat() string {
	format := "20060102"
	if w.dateType == HOUR {
		format = "2006010215"
	}
	return format
}

func (w *DateWriter) cleanOldLogs() {
	format := "20060102"
	duration := -time.Hour * 24
	if w.dateType == HOUR {
		format = "2006010215"
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
		return currDateHour
	}
	return currDateDay // DAY
}

func (w *DateWriter) getFileDate() string {
	fi, err := w.currFile.Stat()
	if err == nil {
		return fi.ModTime().Format(w.getFotmat())
	} else {
		return ""
	}
}
