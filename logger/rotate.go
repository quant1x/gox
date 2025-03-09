package logger

import (
	"sync"
	"time"
)

const (
	timeFmtHour = "2006010215"
	timeFmtDay  = "20060102"
)

type TimeRotate struct {
	unixTime int64
	dateHour string
	dateDay  string
	finished chan struct{}
	mutex    sync.Mutex
	ticker   *time.Ticker
}

func NewTimeRotate() *TimeRotate {
	tr := &TimeRotate{
		finished: make(chan struct{}),
		ticker:   time.NewTicker(time.Millisecond),
	}
	tr.update()
	return tr
}

func (t *TimeRotate) Close() {
	close(t.finished)
	t.ticker.Stop()
	for len(t.ticker.C) > 0 {
		<-t.ticker.C
	}
}

func (t *TimeRotate) GetUnixTime() int64 {
	return t.unixTime
}

func (t *TimeRotate) GetDateHour() string {
	return t.dateHour
}

func (t *TimeRotate) GetDateDay() string {
	return t.dateDay
}

func (t *TimeRotate) update() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	now := time.Now()
	t.unixTime = now.Unix()
	t.dateHour = now.Format(timeFmtHour)
	t.dateDay = now.Format(timeFmtDay)
}

func (t *TimeRotate) AutoUpdate() {
	for {
		select {
		case <-t.ticker.C:
			//fmt.Println("安全触发")
			t.update()
		case <-t.finished:
			//fmt.Println("Ticker已停止")
			return
		}
	}
}
