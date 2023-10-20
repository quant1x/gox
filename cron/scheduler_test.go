package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	c := New()
	//c := New(WithSeconds())
	c.Start()
	_, err := c.AddFuncWithSkipIfStillRunning("@every 1s", func() {
		time.Sleep(time.Second * 3)
		fmt.Printf("SkipIfStillRunningWithLogger: %v\n", time.Now())
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(100 * time.Second)

}
