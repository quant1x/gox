package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	intervalSnapshot := "@every 1s"
	intervalSnapshot = "* 15-59 8 * * *"
	//c := New()
	c := New(WithSeconds())
	c.Start()
	_, err := c.AddJobWithSkipIfStillRunning(intervalSnapshot, func() {
		time.Sleep(time.Second * 3)
		fmt.Printf("SkipIfStillRunningWithLogger: %v\n", time.Now())
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(100 * time.Second)

}
