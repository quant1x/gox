package runtime

import "runtime"

// Caller 获取正在运行的函数名, 文件名以及行号
func Caller() (function, file string, line int) {
	pcs := make([]uintptr, 1)
	runtime.Callers(2, pcs)
	pc := pcs[0]
	f := runtime.FuncForPC(pc)
	function = f.Name()
	file, line = f.FileLine(pc)
	return
}
