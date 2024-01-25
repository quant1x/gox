package progressbar

import (
	"fmt"
	"sync"
	"sync/atomic"
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
	closed   atomic.Uint32
	start    time.Time
}

var (
// bar1 string
// bar2 string
)

const (
	defaultFast      = 20
	defaultSlow      = 5
	defaultTickTimes = time.Millisecond * 1
)

//func initBar(width int) {
//	for i := 0; i < width; i++ {
//		bar1 += "="
//		bar2 += "-"
//	}
//}

func NewBar(line int, prefix string, total int) *Bar {
	if total <= 0 {
		return nil
	}

	if line <= 0 {
		gMaxLine++
		line = gMaxLine
	}
	if line > gMaxLine {
		gMaxLine = line
	}

	bar := &Bar{
		line:     line,
		prefix:   prefix,
		total:    total,
		fast:     defaultFast,
		slow:     defaultSlow,
		width:    100,
		advance:  make(chan struct{}),
		done:     make(chan struct{}),
		finished: make(chan struct{}),
		currents: make(map[string]int),
		change:   1,
		start:    time.Now(),
	}
	//initBar(bar.width)
	bar.initBar(bar.width)
	go bar.updateCost()
	go bar.run()

	return bar
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
	step := 1
	if len(n) > 0 {
		step = n[0]
	}

	b.current += step

	lastRate := b.rate
	lastSpeed := b.speed
	// 计算速度
	b.count()
	filled := b.total == b.current
	// 速度是否有改变的条件
	speedChanged := lastRate != b.rate || lastSpeed != b.speed || !filled

	if b.closed.Load() == 0 && speedChanged {
		// 如果速度有改变, 继续发送更新进度条的信号
		b.mu.Unlock()
		b.advance <- struct{}{}
		b.mu.Lock()
	}
	// 进度条达到100%, 通知工作协程结束并关闭channel
	if b.rate >= 100 || filled {
		// 通知updateCost协程结束
		b.mu.Unlock()
		b.done <- struct{}{}
		b.mu.Lock()
		close(b.done)
		// 阻塞, 等待updateCost协程设置关闭状态
		for b.closed.Load() == 0 {
			fmt.Println(b.prefix, "4:add")
			time.Sleep(defaultTickTimes)
		}
		close(b.advance)
	}
	b.mu.Unlock()
}

func (b *Bar) count() {
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
}

func (b *Bar) updateCost() {
	//defer runtime.IgnorePanic()
	for {
		select {
		case <-time.After(defaultTickTimes):
			b.mu.Lock()
			// 统计数据
			b.count()
			b.mu.Unlock()
			if b.closed.Load() == 0 {
				// 这里是为了增加刷新频次
				b.advance <- struct{}{}
			} else {
				fmt.Println(b.prefix, "1:updateCost")
				return
			}
		case <-b.done:
			// 收到结束信号, 设置关闭状态, 返回
			b.closed.Store(1)
			fmt.Println(b.prefix, "2:updateCost")
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
		defer b.closed.Store(1)
		fmt.Println(b.prefix, "3:run")
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

	//realBar1 := bar1[:bar1Len]
	realBar1 := b.bar1[:bar1Len]
	var realBar2 string
	if bar2Len > 0 {
		//realBar2 = ">" + bar2[:bar2Len-1]
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
