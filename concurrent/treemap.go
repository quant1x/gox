package concurrent

import (
	"cmp"

	rbt "github.com/quant1x/gox/util/redblacktree"
)

type TreeMap[K comparable, V any] struct {
	//treemap.Map
	tree *rbt.Tree
}

func NewTreeMap[K cmp.Ordered, V any]() *TreeMap[K, V] {
	tree := rbt.Tree{Comparator: func(a, b any) int {
		A := a.(K)
		B := b.(K)
		return cmp.Compare(A, B)
	}}
	return &TreeMap[K, V]{tree: &tree}
}

func (m *TreeMap[K, V]) Size() int {
	return m.tree.Size()
}

func (m *TreeMap[K, V]) Put(k K, v V) {
	m.tree.Put(k, v)
}

func (m *TreeMap[K, V]) Get(k K) (v V, found bool) {
	tmp, found := m.tree.Get(k)
	if !found {
		return
	}
	return tmp.(V), true
}

func (m *TreeMap[K, V]) Each(f func(key K, value V)) {
	iterator := m.tree.Iterator()
	for iterator.Next() {
		f(iterator.Key().(K), iterator.Value().(V))
	}
}

func (m *TreeMap[K, V]) Clear() {
	m.tree.Clear()
}
