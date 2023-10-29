package cron

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	crontabInterval  = 10
	crontabSnapshot  = "*/1 * * * * ?"
	sleepMillisecond = int64(100)
)

type Scheduler struct {
	crontabInterval int64 // 单位是秒
	crontabSequence int64
	m               sync.Mutex
	cron            *Cron
	service         func()
}

// NewScheduler 创建一个业务定时器
func NewScheduler(interval int, service func()) *Scheduler {
	crontab := Scheduler{
		crontabInterval: int64(interval) * 1000,
		crontabSequence: 0,
		service:         service,
	}
	crontab.cron = New(WithSeconds())
	return &crontab
}

func (this *Scheduler) Run() error {
	_, err := this.cron.AddFunc(crontabSnapshot, func() {
		for atomic.LoadInt64(&this.crontabSequence) != 0 {
			return
		}
		this.m.Lock()
		defer this.m.Unlock()
		if this.crontabSequence != 0 {
			return
		}
		atomic.StoreInt64(&this.crontabSequence, 1)
		now := time.Now()
		// 执行业务
		this.service()
		atomic.StoreInt64(&this.crontabSequence, 0)
		elapsedTime := sleepMillisecond
		for {
			elapsedTime = time.Since(now).Milliseconds()
			sleepTimes := this.crontabInterval - elapsedTime
			if sleepTimes > 0 {
				time.Sleep(time.Millisecond * time.Duration(sleepTimes))
			} else {
				break
			}
		}
	})
	if err != nil {
		return err
	}
	this.cron.Start()
	return nil
}
