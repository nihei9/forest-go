package forest

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type tsNode[K constraints.Ordered, V any] struct {
	split K
	lt    *tsNode[K, V]
	eq    *tsNode[K, V]
	gt    *tsNode[K, V]
	end   bool
	val   V
}

type TernarySearchTree[K constraints.Ordered, V any] struct {
	root      *tsNode[K, V]
	count     int
	maxKeyLen int
}

func NewTernarySearchTree[K constraints.Ordered, V any]() *TernarySearchTree[K, V] {
	return &TernarySearchTree[K, V]{}
}

func (t *TernarySearchTree[K, V]) Insert(key []K, value V) error {
	if len(key) == 0 {
		return fmt.Errorf("key must not be empty")
	}
	ok := t.insertTo(&t.root, key, value)
	if !ok {
		return fmt.Errorf("key already exist: %v", key)
	}
	t.count++
	if len(key) > t.maxKeyLen {
		t.maxKeyLen = len(key)
	}
	return nil
}

func (t *TernarySearchTree[K, V]) Search(key []K) (value V, found bool) {
	if len(key) == 0 {
		return
	}
	n := t.search(t.root, key)
	if n != nil && n.end {
		return n.val, true
	}
	return
}

func (t *TernarySearchTree[K, V]) List(prefix []K) [][]K {
	return ApplyToTernarySearchTree(t, prefix, func(key []K, val V) []K {
		return key
	})
}

func (t *TernarySearchTree[K, V]) Delete(key []K) (value V, found bool) {
	if len(key) == 0 {
		return
	}
	n := t.search(t.root, key)
	if n == nil {
		return
	}
	n.end = false
	var zero V
	val := n.val
	n.val = zero
	return val, true
}

func (t *TernarySearchTree[K, V]) insertTo(node **tsNode[K, V], key []K, value V) bool {
	if *node == nil {
		*node = &tsNode[K, V]{
			split: key[0],
		}
	}
	switch {
	case key[0] < (*node).split:
		return t.insertTo(&(*node).lt, key, value)
	case key[0] > (*node).split:
		return t.insertTo(&(*node).gt, key, value)
	default:
		if len(key) > 1 {
			return t.insertTo(&(*node).eq, key[1:], value)
		}
		if (*node).end {
			return false
		}
		(*node).end = true
		(*node).val = value
		return true
	}
}

func (t *TernarySearchTree[K, V]) search(node *tsNode[K, V], prefix []K) *tsNode[K, V] {
	switch {
	case node == nil:
		return nil
	case prefix[0] < node.split:
		return t.search(node.lt, prefix)
	case prefix[0] > (*node).split:
		return t.search(node.gt, prefix)
	default:
		if len(prefix) > 1 {
			return t.search(node.eq, prefix[1:])
		}
		return node
	}
}

func ApplyToTernarySearchTree[K constraints.Ordered, V any, R any](t *TernarySearchTree[K, V], prefix []K, callback func([]K, V) R) []R {
	if len(prefix) > t.maxKeyLen {
		return nil
	}

	w := &tstWalker[K, V, R]{
		keyBuf:   make([]K, t.maxKeyLen),
		results:  make([]R, t.count),
		callback: callback,
	}

	root := t.root
	if len(prefix) > 0 {
		n := t.search(t.root, prefix)
		if n == nil {
			return nil
		}
		for i, p := range prefix {
			w.keyBuf[i] = p
		}
		if n.end {
			w.apply(len(prefix), n.val)
		}
		root = n.eq
	}
	w.walk(root, len(prefix))
	return w.results[:w.resultPtr]
}

type tstWalker[K constraints.Ordered, V any, R any] struct {
	keyBuf    []K
	results   []R
	resultPtr int
	callback  func([]K, V) R
}

func (w *tstWalker[K, V, R]) walk(node *tsNode[K, V], bufPtr int) {
	if node == nil {
		return
	}

	w.walk(node.lt, bufPtr)

	if node.end {
		w.keyBuf[bufPtr] = node.split
		w.apply(bufPtr+1, node.val)
	}
	if node.eq != nil {
		w.keyBuf[bufPtr] = node.split
		w.walk(node.eq, bufPtr+1)
	}

	w.walk(node.gt, bufPtr)
}

func (w *tstWalker[K, V, R]) apply(end int, val V) {
	key := make([]K, end)
	copy(key, w.keyBuf[:end])
	w.results[w.resultPtr] = w.callback(key, val)
	w.resultPtr++
}
