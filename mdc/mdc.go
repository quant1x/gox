package mdc

import (
	"github.com/mymmsc/gox/gls"
	uuid "github.com/satori/go.uuid"
)

const (
	APP_TRACEID = "app-traceid"
)

func init() {
	RefreshTraceId()
}

func RefreshTraceId()  {
	u1 := uuid.NewV4()
	Set(APP_TRACEID, u1.String())
}

func Set(key string, value interface{}) {
	gls.Set(key, value)
}

func Get(key string) (interface{}) {
	return gls.Get(key)
}

func Remove(key string) {
	gls.Remove(key)
}
