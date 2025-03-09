package logger

import (
	"bytes"
	"context"
	"fmt"
	"gitee.com/quant1x/gox/api"
	"gitee.com/quant1x/gox/cache"
	"gitee.com/quant1x/gox/mdc"
	"gitee.com/quant1x/gox/signal"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
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

const (
	loggerRollerDays int = 7 // 保持7天
	loggerLocalSkip      = 2
	timeFmtTimestamp     = "2006-01-02T15:04:05.000"
	loggerTraceId        = mdc.APP_TRACEID
)

var (
	loggerPath string
	logLevel   = DEBUG
	logQueue   = make(chan *logValue, 10000)
	loggerMap  sync.Map
	timeRotate = NewTimeRotate()
	pool       cache.Pool[logValue]
	finished   chan struct{}
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
	fatal  bool
}

func init() {
	finished = make(chan struct{})
	go timeRotate.AutoUpdate()
	go flushLog(true)

	// 创建监听退出chan
	sigs := signal.Notify()

	_, cancel := context.WithCancel(context.Background())

	go func() {
		s := <-sigs
		Infof("exit sign, [%+v]", s)
		FlushLogger()
		fmt.Println("exit", s)
		cancel()
		timeRotate.Close()
		os.Exit(0)
	}()

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

// SetLogPath 设置日志路径, 默认是INFO级别日志
//
//	Deprecated: 推荐使用 InitLogger
func SetLogPath(path string) {
	InitLogger(path, INFO)
}

// getApplicationName 获取执行文件名
func getApplicationName() string {
	path, _ := os.Executable()
	_, exec := filepath.Split(path)
	arr := strings.Split(exec, ".")
	__applicationName := arr[0]
	return __applicationName
}

// InitLogger 初始化
func InitLogger(path string, level ...LogLevel) {
	// 日志路径非空, 赋值
	if !api.IsEmpty(path) {
		loggerPath = path
	}
	name := getApplicationName()
	loggerPath = filepath.Join(loggerPath, name)

	// 日志级别默认是INFO
	optLevel := INFO
	if len(level) > 0 {
		optLevel = level[0]
	}
	SetLevel(optLevel)
}

// GetLogger return an logger instance
func GetLogger(name string) *Logger {
	v, found := loggerMap.Load(name)
	if found {
		return v.(*Logger)
	}
	lg := &Logger{
		name:   name,
		writer: &ConsoleWriter{},
	}
	_ = lg.SetDayRoller(loggerPath, loggerRollerDays)
	loggerMap.Store(name, lg)
	return lg
}

func SetLevel(level LogLevel) {
	logLevel = level
}

// IsDebug 是否DEBUG模式
func IsDebug() bool {
	if DEBUG < logLevel {
		return false
	}
	return true
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

func (l *Logger) Debug(v ...any) {
	l.writef(loggerLocalSkip, DEBUG, "", v)
}

func (l *Logger) Info(v ...any) {
	l.writef(loggerLocalSkip, INFO, "", v)
}

func (l *Logger) Warn(v ...any) {
	l.writef(loggerLocalSkip, WARN, "", v)
}

func (l *Logger) Error(v ...any) {
	l.writef(loggerLocalSkip, ERROR, "", v)
}

func (l *Logger) Debugf(format string, v ...any) {
	l.writef(loggerLocalSkip, DEBUG, format, v)
}

func (l *Logger) Infof(format string, v ...any) {
	l.writef(loggerLocalSkip, INFO, format, v)
}

func (l *Logger) Warnf(format string, v ...any) {
	l.writef(loggerLocalSkip, WARN, format, v)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.writef(loggerLocalSkip, ERROR, format, v)
}

func (l *Logger) Fatal(v ...any) {
	l.writef(loggerLocalSkip, FATAL, "", v)
	waitForExit()
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.writef(loggerLocalSkip, FATAL, format, v)
	waitForExit()
}

func getTraceId() string {
	traceId := mdc.Get(loggerTraceId)
	t := reflect.ValueOf(traceId)
	if t.Kind() == reflect.String {
		return t.String()
	}
	return ""
}

func (l *Logger) writef(skip int, level LogLevel, format string, v []any) {
	if level < logLevel {
		return
	}

	t := time.Now()
	buf := bytes.NewBuffer(nil)
	if l.writer.NeedPrefix() {
		traceId := getTraceId()
		_, _ = fmt.Fprintf(buf, "%s|%s|", t.Format(timeFmtTimestamp), traceId)
		if logLevel == DEBUG {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				file = "???"
				line = 0
			} else {
				file = filepath.Base(file)
			}
			_, _ = fmt.Fprintf(buf, "%s:%d|", file, line)
		}
	}
	buf.WriteString(level.String())
	buf.WriteByte('|')

	if format == "" {
		_, _ = fmt.Fprint(buf, v...)
	} else {
		_, _ = fmt.Fprintf(buf, format, v...)
	}
	if l.writer.NeedPrefix() {
		buf.WriteByte('\n')
	}
	lv := pool.Acquire()
	lv.value = buf.Bytes()
	lv.writer = l.writer
	lv.fatal = level == FATAL
	logQueue <- lv
}

func FlushLogger() {
	flushLog(false)
}

// 等待结束信号并退出
func waitForExit() {
	<-finished
	os.Exit(-1)
}

func refreshLogFile(v *logValue) {
	if v == nil {
		return
	}
	v.writer.Write(v.value)
	defer pool.Release(v)
	if !v.fatal {
		return
	}
	// 致命的日志, 同时输出到控制台
	fmt.Println(api.Bytes2String(v.value))
	// 发送结束信号
	finished <- struct{}{}
}

func flushLog(sync bool) {
	if sync {
		for v := range logQueue {
			refreshLogFile(v)
		}
	} else {
		for {
			select {
			case v := <-logQueue:
				refreshLogFile(v)
			default:
				return
			}
		}
	}
}

func Info(v ...any) {
	logger := GetLogger("runtime")
	logger.writef(loggerLocalSkip, INFO, "", v)
}

func Infof(format string, v ...any) {
	logger := GetLogger("runtime")
	logger.writef(loggerLocalSkip, INFO, format, v)
}

func Debug(v ...any) {
	logger := GetLogger("debug")
	logger.writef(loggerLocalSkip, DEBUG, "", v)
}

func Debugf(format string, v ...any) {
	logger := GetLogger("debug")
	logger.writef(loggerLocalSkip, DEBUG, format, v)
}

func Warn(v ...any) {
	logger := GetLogger("warn")
	logger.writef(loggerLocalSkip, WARN, "", v)
}

func Warnf(format string, v ...any) {
	logger := GetLogger("warn")
	logger.writef(loggerLocalSkip, WARN, format, v)
}

func Error(v ...any) {
	logger := GetLogger("error")
	logger.writef(loggerLocalSkip, ERROR, "", v)
}

func Errorf(format string, v ...any) {
	logger := GetLogger("error")
	logger.writef(loggerLocalSkip, ERROR, format, v)
}

func Fatal(v ...any) {
	logger := GetLogger("error")
	logger.writef(loggerLocalSkip, FATAL, "", v)
	waitForExit()
}

func Fatalf(format string, v ...any) {
	logger := GetLogger("error")
	logger.writef(loggerLocalSkip, FATAL, format, v)
	waitForExit()
}
