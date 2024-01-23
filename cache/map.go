package cache

import "sync"

// Map 二次封装的泛型sync.Map
type Map[K any, V any] struct {
	syncMap sync.Map
}

func (this *Map[K, V]) Get(k K) (V, bool) {
	obj, ok := this.syncMap.Load(k)
	var v V
	if ok {
		v, ok = obj.(V)
	}
	return v, ok
}

func (this *Map[K, V]) Put(k K, v V) {
	this.syncMap.Store(k, v)
}
