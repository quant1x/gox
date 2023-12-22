package api

import (
	"fmt"
	"testing"
)

func TestCreateTime(t *testing.T) {
	ft := GetFileStat("./filestat.go")
	fmt.Println(ft)
}
