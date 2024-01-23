package cache

import "sync"

// Pool 二次封装的泛型sync.Pool
type Pool[E any] struct {
	once sync.Once // 初始化sync.Pool的New接口
	pool sync.Pool // sync.Pool
	zero E         // 零值
}

// 初始化sync.Pool.New
func (this *Pool[E]) init() {
	this.pool = sync.Pool{New: func() any {
		var e E
		return &e
	}}
}

// Acquire 申请内存
func (this *Pool[E]) Acquire() *E {
	this.once.Do(this.init)
	obj := this.pool.Get().(*E)
	*obj = this.zero
	return obj
}

// Release 释放内存
func (this *Pool[E]) Release(obj *E) {
	this.once.Do(this.init)
	if obj == nil {
		return
	}
	this.pool.Put(obj)
}
