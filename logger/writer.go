package logger

const (
	DAY DateType = iota
	HOUR
)

type LogWriter interface {
	Write(v []byte)
	NeedPrefix() bool
}
