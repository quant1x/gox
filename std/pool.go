package std

// ObjectPool 对象池
type ObjectPool[T any] interface {
	Acquire() *T
	Release(obj *T)
}
