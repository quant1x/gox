package pool

import "errors"

var (
	//ErrMaxActiveConnReached 连接池超限
	ErrMaxActiveConnReached = errors.New("max active conn reached")
	// ErrClosed 连接已关闭
	ErrClosed = errors.New("pool is closed")
	// ErrIsNil 连接无效
	ErrIsNil = errors.New("connection is nil. rejecting")
)

// Pool 基本方法
type Pool interface {
	Get() (any, error)
	Put(any) error
	Close(any) error
	CloseAll()
	Release()
	Len() int
}
