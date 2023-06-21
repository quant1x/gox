package cron

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	crontabInterval  = 10
	crontabSnapshot  = "0/1 * * * * ?"
	sleepMillisecond = 1
)

type Scheduler struct {
	crontabInterval float64 // 单位是秒
	crontabSequence int64
	m               sync.Mutex
	cron            *Cron
	service         func()
}

// NewScheduler 创建一个业务定时器
func NewScheduler(interval int, service func()) *Scheduler {
	crontab := Scheduler{
		crontabInterval: float64(interval),
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
		diffDuration := float64(sleepMillisecond) / 1000.00
		for {
			elapsedTime := time.Since(now).Seconds()
			if elapsedTime+diffDuration < this.crontabInterval {
				sleepStart := time.Now()
				time.Sleep(time.Millisecond * sleepMillisecond)
				diffDuration = float64(time.Since(sleepStart).Milliseconds()) / 1000
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
