package mdc

import (
	"github.com/mymmsc/gox/gls"
)

const (
	APP_TRACEID = "app-traceid"
)

func Set(key string, value interface{}) {
	gls.Set(key, value)
}

func Get(key string) (interface{}) {
	return gls.Get(key)
}

func Remove(key string) {
	gls.Remove(key)
}
