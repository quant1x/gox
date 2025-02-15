package gls

import (
	"github.com/petermattis/goid"
)

func GoID() int64 {
	return goid.Get()
}
