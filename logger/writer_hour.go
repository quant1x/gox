package logger

import (
	"os"
	"path/filepath"
	"strconv"
)

type RollFileWriter struct {
	logpath  string
	name     string
	num      int
	size     int64
	currSize int64
	currFile *os.File
	openTime int64
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

func (w *RollFileWriter) Write(v []byte) {
	if w.currFile == nil || w.openTime+10 < timeRotate.GetUnixTime() {
		fullPath := filepath.Join(w.logpath, w.name+".log")
		reopenFile(fullPath, &w.currFile, &w.openTime)
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
		reopenFile(fullPath, &w.currFile, &w.openTime)
	}
}

func (w *RollFileWriter) NeedPrefix() bool {
	return true
}
