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

// New 创建一个新的错误信息, 包含一个状态码和信息
func New(code int, message string, a ...any) *Exception {
	return &Exception{
		code:    code,
		message: fmt.Sprintf(message, a...),
	}
}

// 格式化输出错误信息
func (this Exception) Error() string {
	return fmt.Sprintf("#%d, message=%s", this.code, this.message)
}

func (this Exception) Code() int {
	return this.code
}
