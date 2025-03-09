package logger

import (
	"testing"
	"time"
)

func TestRotate(t *testing.T) {
	tr := NewTimeRotate()
	go tr.AutoUpdate()
	time.Sleep(2 * time.Second)
	tr.Close()
}
