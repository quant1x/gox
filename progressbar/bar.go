package progressbar

import (
	"fmt"
	"sync"
	"time"
)

// Bar 进度条
type Bar struct {
	mu       sync.Mutex
	line     int
	prefix   string
	total    int
	width    int
	bar1     string
	bar2     string
	advance  chan struct{} // for data
	done     chan struct{} // for updateCost
	finished chan struct{} // 等待结束信号
	currents map[string]int
	current  int
	before   int
	rate     int
	speed    int
	cost     int
	estimate int
	fast     int
	slow     int
	srcUnit  string
	dstUnit  string
	change   int
	start    time.Time
}

const (
	defaultFast       = 20
	defaultSlow       = 5
	defaultTickTimes  = time.Millisecond * 500 // 默认定时刷新进度条的间隔时间
	defaultSleepTimes = time.Millisecond * 100 // 默认sleep时间
)

func NewBar(line int, prefix string, total int) *Bar {
	if total <= 0 {
		return nil
	}

	line = adjustLine(line)

	bar := &Bar{
		line:     line,
		prefix:   prefix,
		total:    total,
		fast:     defaultFast,
		slow:     defaultSlow,
		width:    100,
		advance:  make(chan struct{}, total),
		done:     make(chan struct{}, total),
		finished: make(chan struct{}),
		currents: make(map[string]int),
		change:   1,
		start:    time.Now(),
	}
	bar.initBar(bar.width)
	go bar.updateCost()
	go bar.run()

	return bar
}

// SetPrefix 设置进度条前端文字
func (b *Bar) SetPrefix(prefix string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.prefix = prefix
}

func (b *Bar) initBar(width int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for i := 0; i < width; i++ {
		b.bar1 += "="
		b.bar2 += "-"
	}
}

func (b *Bar) SetUnit(src string, dst string, change int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.srcUnit = src
	b.dstUnit = dst
	b.change = change

	if b.change == 0 {
		b.change = 1
	}
}

func (b *Bar) SetSpeedSection(fast, slow int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if fast > slow {
		b.fast, b.slow = fast, slow
	} else {
		b.fast, b.slow = slow, fast
	}
}

// Add 进度条前进, 同时也会触发其它信号
func (b *Bar) Add(n ...int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	step := 1
	if len(n) > 0 {
		step = n[0]
	}
	b.current += step
	// 计算速度
	b.computeAndUpdate()
	// 进度条达到100%, 通知工作协程结束并关闭channel
	if b.current >= b.total {
		// 通知updateCost协程结束
		b.done <- struct{}{}
		close(b.done)
	}
}

func (b *Bar) computeAndUpdate() {
	now := time.Now()
	nowKey := now.Format("20060102150405")
	befKey := now.Add(time.Minute * -1).Format("20060102150405")
	b.currents[nowKey] = b.current
	if v, ok := b.currents[befKey]; ok {
		b.before = v
	}
	delete(b.currents, befKey)

	b.rate = b.current * 100 / b.total
	b.cost = int(time.Since(b.start) / time.Second)
	if b.cost == 0 {
		b.speed = b.current * 100
	} else if b.before == 0 {
		b.speed = b.current * 100 / b.cost
	} else {
		b.speed = (b.current - b.before) * 100 / 60
	}

	if b.speed != 0 {
		b.estimate = (b.total - b.current) * 100 / b.speed
	}
	b.advance <- struct{}{} // 发送更新进度条信号
}

// 发送bar更新信号
func (b *Bar) sendBarUpdateSignal() {
	b.mu.Lock()
	defer b.mu.Unlock()
	// 统计数据
	b.computeAndUpdate()
}

func (b *Bar) updateCost() {
	for {
		select {
		case <-time.After(defaultTickTimes):
			b.sendBarUpdateSignal()
		case <-b.done:
			// 补全不满100%进度的信号
			for b.rate <= 100 {
				b.sendBarUpdateSignal()
				if b.rate >= 100 {
					break
				}
				time.Sleep(defaultSleepTimes)
			}
			close(b.advance)
			return
		}
	}
}

// Wait 等待进度条刷新run协程结束
func (b *Bar) Wait() {
	<-b.finished
}

func (b *Bar) run() {
	defer func() {
		b.finished <- struct{}{}
	}()
	// 只有关闭channel才会结束循环, 且不能对channel加速
	for range b.advance {
		text := b.barMsg()
		barPrintf(b.line, "\r%s", text)
	}
}

// 重置进度条消息
func (b *Bar) barMsg() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	unit := ""
	change := 1
	if b.srcUnit != "" {
		unit = b.srcUnit
	}

	if b.dstUnit != "" {
		unit = b.dstUnit
		change = b.change
	}

	prefix := fmt.Sprintf("%s", b.prefix)
	rate := fmt.Sprintf("%3d%%", b.rate)
	speed := fmt.Sprintf("%3.2f %s ps", 0.01*float64(b.speed)/float64(change), unit)
	cost := b.timeFmt(b.cost)
	estimate := b.timeFmt(b.estimate)
	ct := fmt.Sprintf(" (%d/%d)", b.current, b.total)
	barLen := b.width - len(prefix) - len(rate) - len(speed) - len(cost) - len(estimate) - len(ct) - 10
	bar1Len := barLen * b.rate / 100
	bar2Len := barLen - bar1Len

	realBar1 := b.bar1[:bar1Len]
	var realBar2 string
	if bar2Len > 0 {
		realBar2 = ">" + b.bar2[:bar2Len-1]
	}

	msg := fmt.Sprintf(`%s %s%s [%s%s] %s %s in: %s`, prefix, rate, ct, realBar1, realBar2, speed, cost, estimate)
	switch {
	case b.speed <= b.slow*100:
		return "\033[0;31m" + msg + "\033[0m"
	case b.speed > b.slow*100 && b.speed < b.fast*100:
		return "\033[0;33m" + msg + "\033[0m"
	default:
		return "\033[0;32m" + msg + "\033[0m"
	}
}

func (b *Bar) timeFmt(cost int) string {
	var h, m, s int
	h = cost / 3600
	m = (cost - h*3600) / 60
	s = cost - h*3600 - m*60

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
