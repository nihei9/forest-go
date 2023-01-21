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

func (b *tstListBuf[K]) setup(entryCount int, maxKeyLen int) {
	if entryCount > len(b.entries) {
		b.entries = make([][]K, entryCount)
	}
	b.entPtr = 0
	if maxKeyLen > len(b.keyBuf) {
		b.keyBuf = make([]K, maxKeyLen)
	}
}

func (b *tstListBuf[K]) put(i int, k K) {
	b.keyBuf[i] = k
}

func (b *tstListBuf[K]) append(end int) {
	key := make([]K, end+1)
	copy(key, b.keyBuf[:end+1])
	b.entries[b.entPtr] = key
	b.entPtr++
}

func (b *tstListBuf[K]) list() [][]K {
	entries := make([][]K, b.entPtr)
	copy(entries, b.entries[:b.entPtr])
	return entries
}

type TernarySearchTree[K constraints.Ordered, V any] struct {
	root      *tsNode[K, V]
	count     int
	maxKeyLen int

	listBuf *tstListBuf[K]
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
	if len(prefix) > t.maxKeyLen {
		return nil
	}

	if t.listBuf == nil {
		t.listBuf = &tstListBuf[K]{}
	}
	t.listBuf.setup(t.count, t.maxKeyLen)

	root := t.root
	if len(prefix) > 0 {
		n := t.search(t.root, prefix)
		if n == nil {
			return nil
		}
		for i, p := range prefix {
			t.listBuf.put(i, p)
		}
		if n.end {
			t.listBuf.append(len(prefix) - 1)
		}
		root = n.eq
	}
	t.list(len(prefix), root)
	return t.listBuf.list()
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

func (t *TernarySearchTree[K, V]) list(bufPtr int, node *tsNode[K, V]) {
	if node == nil {
		return
	}

	t.list(bufPtr, node.lt)

	if node.end {
		t.listBuf.put(bufPtr, node.split)
		t.listBuf.append(bufPtr)
	}
	if node.eq != nil {
		t.listBuf.put(bufPtr, node.split)
		t.list(bufPtr+1, node.eq)
	}

	t.list(bufPtr, node.gt)
}
