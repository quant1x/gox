//go:build windows

package logger

import (
	"os"
	"syscall"
)

var stopSignals = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
