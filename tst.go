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

type tstListBuf[K constraints.Ordered] struct {
	entries [][]K
	entPtr  int
	keyBuf  []K
}

type TernarySearchTree[K constraints.Ordered, V any] struct {
	root      *tsNode[K, V]
	count     int
	maxKeyLen int

	tstListBuf[K]
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
	n, ok := t.search(t.root, key)
	if !ok {
		return
	}
	return n.val, true
}

func (t *TernarySearchTree[K, V]) List() [][]K {
	if t.count > len(t.entries) {
		t.entries = make([][]K, t.count)
	}
	t.entPtr = 0
	if t.maxKeyLen > len(t.keyBuf) {
		t.keyBuf = make([]K, t.maxKeyLen)
	}
	t.list(0, t.root)
	entries := make([][]K, t.count)
	copy(entries, t.entries)
	return entries
}

func (t *TernarySearchTree[K, V]) Delete(key []K) (value V, found bool) {
	if len(key) == 0 {
		return
	}
	n, ok := t.search(t.root, key)
	if !ok {
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

func (t *TernarySearchTree[K, V]) search(node *tsNode[K, V], key []K) (*tsNode[K, V], bool) {
	switch {
	case node == nil:
		return nil, false
	case key[0] < node.split:
		return t.search(node.lt, key)
	case key[0] > (*node).split:
		return t.search(node.gt, key)
	default:
		if len(key) > 1 {
			return t.search(node.eq, key[1:])
		}
		if !node.end {
			return nil, false
		}
		return node, true
	}
}

func (t *TernarySearchTree[K, V]) list(bufPtr int, node *tsNode[K, V]) {
	if node == nil {
		return
	}

	t.list(bufPtr, node.lt)

	if node.end {
		t.keyBuf[bufPtr] = node.split
		t.entries[t.entPtr] = make([]K, bufPtr+1)
		copy(t.entries[t.entPtr], t.keyBuf[:bufPtr+1])
		t.entPtr++
	}
	if node.eq != nil {
		t.keyBuf[bufPtr] = node.split
		t.list(bufPtr+1, node.eq)
	}

	t.list(bufPtr, node.gt)
}
