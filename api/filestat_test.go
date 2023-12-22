package api

import (
	"fmt"
	"testing"
)

func TestCreateTime(t *testing.T) {
	ft, err := GetFileStat("./filestat.go")
	fmt.Println(ft, err)
}
