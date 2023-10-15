package concurrent

import "sync"

type v1HashMap[K comparable, V any] struct {
	mutex sync.RWMutex
	m     map[K]V
	once  sync.Once
}

func v1NewHashmap[K comparable, V any]() *v1HashMap[K, V] {
	hashmap := v1HashMap[K, V]{}
	hashmap.m = make(map[K]V)
	return &hashmap
}

func (this *v1HashMap[K, V]) Get(key K) (V, bool) {
	this.mutex.RLock()
	v, ok := this.m[key]
	this.mutex.RUnlock()
	return v, ok
}

func (this *v1HashMap[K, V]) Put(key K, value V) {
	this.mutex.Lock()
	this.m[key] = value
	this.mutex.Unlock()
}
