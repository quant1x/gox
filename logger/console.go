package logger

import "os"

type ConsoleWriter struct {
}

func (w *ConsoleWriter) Write(v []byte) {
	_, _ = os.Stdout.Write(v)
}

func (w *ConsoleWriter) NeedPrefix() bool {
	return true
}
