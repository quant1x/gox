package hooks

import (
	"context"
	"gitee.com/quant1x/gox/signal"
	"os"
	"sort"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// HookFunc Hook 函数类型
type HookFunc func(ctx context.Context) error

// HookPriority Hook 优先级排序
type HookPriority int

const (
	HighPriority HookPriority = iota
	DefaultPriority
	LowPriority
)

// Hook 结构体
type hook struct {
	fn       HookFunc
	priority HookPriority
	timeout  time.Duration
}

type HookManager struct {
	hooks      []hook
	mu         sync.Mutex
	once       sync.Once
	exitSignal chan struct{}
	exiting    atomic.Bool
	shutdown   chan os.Signal
}

func NewHookManager() *HookManager {
	hm := &HookManager{
		exitSignal: make(chan struct{}),
		shutdown:   nil,
		//shutdown:   make(chan os.Signal),
	}
	return hm
}

// Register 注册 Hook
func (hm *HookManager) Register(fn HookFunc, opts ...Option) {
	hm.mu.Lock()
	defer hm.mu.Unlock()

	h := hook{
		fn:       fn,
		priority: DefaultPriority,
		timeout:  5 * time.Second,
	}

	for _, opt := range opts {
		opt(&h)
	}

	hm.hooks = append(hm.hooks, h)
}

// Listen 启动信号监听
func (hm *HookManager) Listen() {
	hm.once.Do(func() {
		hm.shutdown = signal.Notify()
		go func() {
			s := <-hm.shutdown
			hm.Trigger(s)
		}()
	})
}

// Trigger 触发关闭流程
func (hm *HookManager) Trigger(s os.Signal) {
	if hm.exiting.Swap(true) {
		return // 避免重复触发
	}

	// 给主进程发结束信号
	// close(hm.exitSignal)

	// 按优先级排序
	hm.mu.Lock()
	sort.Slice(hm.hooks, func(i, j int) bool {
		return hm.hooks[i].priority < hm.hooks[j].priority
	})
	hm.mu.Unlock()

	var wg sync.WaitGroup

	for _, h := range hm.hooks {
		wg.Add(1)
		go func(h hook) {
			defer wg.Done()

			//ctx, cancel := context.WithTimeout(context.Background(), h.timeout)
			//defer cancel()
			ctx := context.WithValue(context.Background(), KeyExit, s)
			//ctx.Value("signal")
			_ = h.fn(ctx)
		}(h)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		//fmt.Println("All hooks completed")
	case <-time.After(10 * time.Second):
		//fmt.Println("Hooks timeout, force exiting")
	}
	//hm.originalExit(0)
	hm.Done()
	//os.Exit(0)
}

func (hm *HookManager) Done() {
	//fmt.Println("hook done")
	close(hm.exitSignal)
}
func (hm *HookManager) SendExitSignal() {
	hm.shutdown <- syscall.SIGINT
	//close(hm.shutdown)
}

//// 劫持 os.Exit
//func (hm *HookManager) exitHandler(code int) {
//	if !hm.exiting.Load() {
//		hm.Trigger()
//	}
//	hm.originalExit(code)
//}

// Wait 等待退出信号
func (hm *HookManager) Wait() {
	<-hm.exitSignal
}

// Option 配置选项
type Option func(*hook)

func WithPriority(p HookPriority) Option {
	return func(h *hook) {
		h.priority = p
	}
}

func WithTimeout(d time.Duration) Option {
	return func(h *hook) {
		h.timeout = d
	}
}
