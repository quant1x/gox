package hooks

var (
	hm *HookManager = nil
)

const (
	KeyExit = "signal"
)

func init() {
	hm = NewHookManager()
}

// Register 注册hook
func Register(fn HookFunc, opts ...Option) {
	hm.Register(fn, opts...)
}

// Done 主动结束进程
func Done() {
	hm.exitSignal <- struct{}{}
}

// WaitForShutdown 等待进程结束信号
func WaitForShutdown() {
	hm.Listen()
	hm.Wait()
}
