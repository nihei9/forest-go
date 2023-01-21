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

// NewTernarySearchTree returns a new ternary search tree that can contain entries mapping `[]K` to `V`.
func NewTernarySearchTree[K constraints.Ordered, V any]() *TernarySearchTree[K, V] {
	return &TernarySearchTree[K, V]{}
}

// Insert inserts an entry. When the key already exists, this function return an error.
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

// Search earches for an entry having a key that exactly matches a specified key and returns its value.
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

type TernarySearchTreeEntry[K constraints.Ordered, V any] struct {
	Key   []K
	Value V
}

// Entries returns entries. When a prefix isn't empty, this function returns entries whose key has the prefix.
func (t *TernarySearchTree[K, V]) Entries(prefix []K) []*TernarySearchTreeEntry[K, V] {
	return ApplyToTernarySearchTree(t, prefix, func(key []K, val V) *TernarySearchTreeEntry[K, V] {
		return &TernarySearchTreeEntry[K, V]{
			Key:   key,
			Value: val,
		}
	})
}

// Keys returns keys. When a prefix isn't empty, this function returns keys having the prefix.
func (t *TernarySearchTree[K, V]) Keys(prefix []K) [][]K {
	return ApplyToTernarySearchTree(t, prefix, func(key []K, val V) []K {
		return key
	})
}

// Values returns values. When a prefix isn't empty, this function returns values whose key has the prefix.
func (t *TernarySearchTree[K, V]) Values(prefix []K) []V {
	return ApplyToTernarySearchTree(t, prefix, func(key []K, val V) V {
		return val
	})
}

// Delete deletes an entry and returns its value.
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

// ApplyToTernarySearchTree applies a user-defined function to each entry whose key has a specified prefix.
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
		copy(w.keyBuf, prefix)
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
