package slf4g

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	OFF
	FATAL
)

// 日志默认以天为单位
const (
	// 保持7天 [wangfeng on 2018/12/25 12:38]
	__logger_roller_days int = 7
	__logget_global_skip = 3
	__logget_local_skip = 2
)

var (
	__logger_path string
)

var (
	logLevel LogLevel = DEBUG

	logQueue  = make(chan *logValue, 10000)
	loggerMap = make(map[string]*Logger)
	writeDone = make(chan bool)

	currUnixTime int64
	currDateTime string
	currDateHour string
	currDateDay  string
)

type Logger struct {
	name   string
	writer LogWriter
}
type LogLevel uint8

type logValue struct {
	level  LogLevel
	value  []byte
	fileNo string
	writer LogWriter
}

func init() {
	now := time.Now()
	currUnixTime = now.Unix()
	currDateTime = now.Format("2006-01-02 15:04:05.999")
	currDateHour = now.Format("2006010215")
	currDateDay = now.Format("20060102")
	go func() {
		tm := time.NewTimer(time.Second)
		if err := recover(); err != nil { // avoid timer panic
		}
		for {
			now := time.Now()
			d := time.Second - time.Duration(now.Nanosecond())
			tm.Reset(d)
			<-tm.C
			now = time.Now()
			currUnixTime = now.Unix()
			currDateTime = now.Format("2006-01-02 15:04:05.999")
			currDateHour = now.Format("2006010215")
			currDateDay = now.Format("20060102")
		}
	}()
	go flushLog(true)
}

func (lv *LogLevel) String() string {
	switch *lv {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func SetLogPath(path string) {
	__logger_path = path
	SetLevel(DEBUG)
}

// GetLogger return an slf4g instance
func GetLogger(name string) *Logger {
	if lg, ok := loggerMap[name]; ok {
		return lg
	}
	lg := &Logger{
		name:   name,
		writer: &ConsoleWriter{},
	}
	lg.SetDayRoller(__logger_path, __logger_roller_days)
	loggerMap[name] = lg
	return lg
}

func SetLevel(level LogLevel) {
	logLevel = level
}

func StringToLevel(level string) LogLevel {
	switch level {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	case "FATAL":
		return FATAL
	default:
		return DEBUG
	}
}

func (l *Logger) SetLogName(name string) {
	l.name = name
}

func (l *Logger) SetFileRoller(logpath string, num int, sizeMB int) error {
	if err := os.MkdirAll(logpath, 0755); err != nil {
		panic(err)
	}
	w := NewRollFileWriter(logpath, l.name, num, sizeMB)
	l.writer = w
	return nil
}

func (l *Logger) IsConsoleWriter() bool {
	if reflect.TypeOf(l.writer) == reflect.TypeOf(&ConsoleWriter{}) {
		return true
	}
	return false
}

func (l *Logger) SetWriter(w LogWriter) {
	l.writer = w
}

func (l *Logger) SetDayRoller(logpath string, num int) error {
	if err := os.MkdirAll(logpath, 0755); err != nil {
		return err
	}
	w := NewDateWriter(logpath, l.name, DAY, num)
	l.writer = w
	return nil
}

func (l *Logger) SetHourRoller(logpath string, num int) error {
	if err := os.MkdirAll(logpath, 0755); err != nil {
		return err
	}
	w := NewDateWriter(logpath, l.name, HOUR, num)
	l.writer = w
	return nil
}

func (l *Logger) SetConsole() {
	l.writer = &ConsoleWriter{}
}

func (l *Logger) Debug(v ...interface{}) {
	l.writef(__logget_local_skip, DEBUG, "", v)
}

func (l *Logger) Info(v ...interface{}) {
	l.writef(__logget_local_skip, INFO, "", v)
}

func (l *Logger) Warn(v ...interface{}) {
	l.writef(__logget_local_skip, WARN, "", v)
}

func (l *Logger) Error(v ...interface{}) {
	l.writef(__logget_local_skip, ERROR, "", v)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.writef(__logget_local_skip, DEBUG, format, v)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.writef(__logget_local_skip, INFO, format, v)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.writef(__logget_local_skip, WARN, format, v)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.writef(__logget_local_skip, ERROR, format, v)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.writef(__logget_local_skip, FATAL, "", v)
	os.Exit(-1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.writef(__logget_local_skip, FATAL, format, v)
	os.Exit(-1)
}

func (l *Logger) writef(skip int, level LogLevel, format string, v []interface{}) {
	if level < logLevel {
		return
	}

	t := time.Now()
	ms := (t.UnixNano() / int64(time.Millisecond)) % 1000
	buf := bytes.NewBuffer(nil)
	if l.writer.NeedPrefix() {
		fmt.Fprintf(buf, "%s.%03d|", t.Format("2006-01-02T15:04:05"), ms)
		if logLevel == DEBUG {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				file = "???"
				line = 0
			} else {
				file = filepath.Base(file)
			}
			fmt.Fprintf(buf, "%s:%d|", file, line)
		}
	}
	buf.WriteString(level.String())
	buf.WriteByte('|')

	if format == "" {
		fmt.Fprint(buf, v...)
	} else {
		fmt.Fprintf(buf, format, v...)
	}
	if l.writer.NeedPrefix() {
		buf.WriteByte('\n')
	}
	logQueue <- &logValue{value: buf.Bytes(), writer: l.writer}
}

func FlushLogger() {
	flushLog(false)
}

func flushLog(sync bool) {
	if sync {
		for v := range logQueue {
			v.writer.Write(v.value)
		}
	} else {
		for {
			select {
			case v := <-logQueue:
				v.writer.Write(v.value)
				continue
			default:
				return
			}
		}
	}
}

func Info(v ...interface{}) {
	logger := GetLogger("runtime")
	logger.writef(__logget_local_skip, INFO, "", v)
}

func Infof(format string, v ...interface{}) {
	logger := GetLogger("runtime")
	logger.writef(__logget_local_skip, INFO, format, v)
}

func Debug(v ...interface{}) {
	logger := GetLogger("debug")
	logger.writef(__logget_local_skip, DEBUG, "", v)
}

func Debugf(format string, v ...interface{}) {
	logger := GetLogger("debug")
	logger.writef(__logget_local_skip, DEBUG, format, v)
}

func Warn(v ...interface{}) {
	logger := GetLogger("warn")
	logger.writef(__logget_local_skip, WARN, "", v)
}

func Warnf(format string, v ...interface{}) {
	logger := GetLogger("warn")
	logger.writef(__logget_local_skip, WARN, format, v)
}

func Error(v ...interface{}) {
	logger := GetLogger("error")
	logger.writef(__logget_local_skip, ERROR, "", v)
}

func Errorf(format string, v ...interface{}) {
	logger := GetLogger("error")
	logger.writef(__logget_local_skip, ERROR, format, v)
}

func Fatal(v ...interface{}) {
	logger := GetLogger("error")
	logger.writef(__logget_local_skip, FATAL, "", v)
	os.Exit(-1)
}

func Fatalf(format string, v ...interface{}) {
	logger := GetLogger("error")
	logger.writef(__logget_local_skip, FATAL, format, v)
	os.Exit(-1)
}