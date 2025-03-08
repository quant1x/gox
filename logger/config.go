package logger

import (
	"fmt"
	"gitee.com/quant1x/gox/signal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Config 日志配置
type Config struct {
	Level         zapcore.Level // 日志级别
	Path          string        // 路径
	EnableConsole bool          // 控制台开关
	MaxAge        time.Duration // 最大保留时间
	RotationTime  time.Duration // 日志切割时间
	BufferSize    int           // 缓冲区大小, 单位KB
	FlushInterval time.Duration // 定时刷新间隔, 单位秒
}

var (
	// --------------------------------------------
	// 1. 定义纯文本编码器
	// --------------------------------------------
	encoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	textEncoder = zapcore.NewConsoleEncoder(encoderConfig)
)

type LogLevel uint8

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	OFF
	FATAL
)

var (
	defaultLevel = DEBUG
	cfg          = Config{
		Level:         zapcore.DebugLevel,
		MaxAge:        7 * 24 * time.Hour,
		RotationTime:  24 * time.Hour,
		BufferSize:    256,
		FlushInterval: 5,
	}
	logger *zap.SugaredLogger = nil
)

func init() {
	tempPath := os.TempDir()
	//cfg.Path = getLogRoot(tempPath)
	//zapLogger := NewTextLoggerWithCompression(cfg)
	//logger = zapLogger.Sugar()
	InitLogger(tempPath, defaultLevel)
	chSignal := signal.Notify()
	go func() {
		s := <-chSignal
		Infof("exit sign, [%+v]", s)
		fmt.Println("exit", s)
		os.Exit(0)
	}()
}

// SetLevel 在临时路径记录日志
func SetLevel(level LogLevel) {
	InitLogger("", level)
}

// IsDebug 是否debug日志模式
func IsDebug() bool {
	return cfg.Level == zapcore.DebugLevel
}

// InitLogger 初始化全局日志模块
func InitLogger(path string, level LogLevel) {
	path = strings.TrimSpace(path)
	if path == "" {
		path = os.TempDir()
	}
	defaultLevel = level
	switch level {
	case DEBUG:
		cfg.Level = zapcore.DebugLevel
	case INFO:
		cfg.Level = zapcore.InfoLevel
	case ERROR:
		cfg.Level = zapcore.ErrorLevel
	case WARN:
		cfg.Level = zapcore.WarnLevel
	default:
		cfg.Level = zapcore.FatalLevel
	}
	cfg.Path = getLogRoot(path)
	fmt.Println(cfg)
	zapLogger := NewTextLoggerWithCompression(cfg)
	logger = zapLogger.Sugar()

}

func getLogRoot(path string) string {
	applicationName := getApplicationName()
	return filepath.Join(path, applicationName)
}

// getApplicationName 获取执行文件名
func getApplicationName() string {
	path, _ := os.Executable()
	_, exec := filepath.Split(path)
	arr := strings.Split(exec, ".")
	__applicationName := arr[0]
	return __applicationName
}
