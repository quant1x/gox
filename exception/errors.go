package exception

import (
	"fmt"
)

type Throwable interface {
	error
	Code() int
}

type Exception struct {
	Throwable
	code    int
	message string
}

func New(code int, message string) *Exception {
	return &Exception{
		code:    code,
		message: message,
	}
}

func (this Exception) Error() string {
	return fmt.Sprintf("#%d, message=%s", this.code, this.message)
}

func (this Exception) Code() int {
	return this.code
}

// PanicIgnore 忽略panic, 继续执行
func PanicIgnore() {
	// 解析失败以后输出日志, 以备检查
	if err := recover(); err != nil {
		//logger.Errorf("panic ignore, error=%+v\n", err)
	}
}
