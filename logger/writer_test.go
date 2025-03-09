package logger

import "testing"

func Test_gzipFile(t *testing.T) {
	_ = gzipFile("writer.go")
}
