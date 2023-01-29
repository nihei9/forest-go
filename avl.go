package forest

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type avlNode[K constraints.Ordered, V any] struct {
	parent *avlNode[K, V]
	split  K
	left   *avlNode[K, V]
	right  *avlNode[K, V]
	val    V
}

func newAVLNode[K constraints.Ordered, V any](parent *avlNode[K, V], split K, value V) *avlNode[K, V] {
	return &avlNode[K, V]{
		parent: parent,
		split:  split,
		val:    value,
	}
}

func (n *avlNode[K, V]) insertAndBalance(key K, value V) (root *avlNode[K, V], maybeUnbalanced bool, err error) {
	switch {
	case key < n.split:
		if n.left == nil {
			n.left = newAVLNode(n, key, value)
			return n, true, nil
		}
		root, maybeUnbalanced, err := n.left.insertAndBalance(key, value)
		if err != nil || !maybeUnbalanced {
			return root, false, err
		}
		root, more := n.balance()
		return root, more, nil
	case key > n.split:
		if n.right == nil {
			n.right = newAVLNode(n, key, value)
			return n, true, nil
		}
		root, maybeUnbalanced, err := n.right.insertAndBalance(key, value)
		if err != nil || !maybeUnbalanced {
			return root, false, err
		}
		root, more := n.balance()
		return root, more, nil
	default:
		return nil, false, fmt.Errorf("key already exist: %v", key)
	}
}

func (n *avlNode[K, V]) balance() (root *avlNode[K, V], more bool) {
	bf := n.balanceFactor()
	if bf == 0 {
		return n, false
	} else if bf == -1 || bf == 1 {
		return n, true
	}

	// balance
	if bf < 0 {
		// left-heavy
		if n.left.leftHight() < n.left.rightHight() {
			n.left.rotateLeft()
			root = n.rotateRight()
		} else {
			root = n.rotateRight()
		}
	} else {
		// right-heavy
		if n.right.leftHight() > n.right.rightHight() {
			n.right.rotateRight()
			root = n.rotateLeft()
		} else {
			root = n.rotateLeft()
		}
	}
	return root, true
}

func (n *avlNode[K, V]) balanceFactor() int {
	return n.rightHight() - n.leftHight()
}

func (n *avlNode[K, V]) leftHight() int {
	if n.left == nil {
		return 0
	}
	return n.left.hight() + 1
}

func (n *avlNode[K, V]) rightHight() int {
	if n.right == nil {
		return 0
	}
	return n.right.hight() + 1
}

func (n *avlNode[K, V]) hight() int {
	var l, r int
	if n.left != nil {
		l = n.left.hight() + 1
	}
	if n.right != nil {
		r = n.right.hight() + 1
	}
	if l > r {
		return l
	}
	return r
}

// rotateLeft rotates a tree left.
func (n *avlNode[K, V]) rotateLeft() *avlNode[K, V] {
	// A          B
	//  \        / \
	//   B   => A   D
	//  / \      \
	// C   D      C

	var parent **avlNode[K, V]
	if n.parent != nil {
		if n.parent.left == n {
			parent = &n.parent.left
		} else {
			parent = &n.parent.right
		}
	}

	pivot := n.right
	pivot.parent = n.parent
	n.parent = pivot
	if pivot.left != nil {
		pivot.left.parent = n
	}
	n.right = pivot.left
	pivot.left = n

	if parent != nil {
		*parent = pivot
	}

	return pivot
}

// rotateRight rotates a tree right.
func (n *avlNode[K, V]) rotateRight() *avlNode[K, V] {
	//     A      B
	//    /      / \
	//   B   => C   A
	//  / \        /
	// C   D      D

	var parent **avlNode[K, V]
	if n.parent != nil {
		if n.parent.left == n {
			parent = &n.parent.left
		} else {
			parent = &n.parent.right
		}
	}

	pivot := n.left
	pivot.parent = n.parent
	n.parent = pivot
	if pivot.right != nil {
		pivot.right.parent = n
	}
	n.left = pivot.right
	pivot.right = n

	if parent != nil {
		*parent = pivot
	}

	return pivot
}

func (n *avlNode[K, V]) search(key K) (node *avlNode[K, V], found bool) {
	switch {
	case key < n.split:
		if n.left == nil {
			return
		}
		return n.left.search(key)
	case key > n.split:
		if n.right == nil {
			return
		}
		return n.right.search(key)
	default:
		return n, true
	}
}

func (n *avlNode[K, V]) deleteAndBalance(key K) (root *avlNode[K, V], value V, found bool) {
	switch {
	case key < n.split:
		if n.left == nil {
			return
		}
		root, value, found = n.left.deleteAndBalance(key)
		if !found {
			return
		}
		root, _ = n.balanceOnDeletion()
		return
	case key > n.split:
		if n.right == nil {
			return
		}
		root, value, found = n.right.deleteAndBalance(key)
		if !found {
			return
		}
		root, _ = n.balanceOnDeletion()
		return
	}

	d := n

	val := d.val
	switch {
	case d.left == nil && d.right == nil:
		if d.parent != nil {
			if d.parent.left == d {
				d.parent.left = nil
			} else {
				d.parent.right = nil
			}
		}

		if d == n {
			root = nil
		} else {
			root = n
		}
	case d.left != nil && d.right == nil:
		alt := d.left
		alt.parent = d.parent
		*d = *alt

		if d == n {
			root = alt
		} else {
			root = n
		}
	case d.left == nil && d.right != nil:
		alt := d.right
		alt.parent = d.parent
		*d = *alt

		if d == n {
			root = alt
		} else {
			root = n
		}
	default:
		alt, leftRoot, shrinked := d.left.extractMax()
		d.split = alt.split
		d.val = alt.val
		d.left = leftRoot

		var r *avlNode[K, V]
		if shrinked {
			r, _ = d.balanceOnDeletion()
		} else {
			r = d
		}

		if d == n {
			root = r
		} else {
			root = n
		}
	}
	return root, val, true
}

func (n *avlNode[K, V]) extractMax() (max *avlNode[K, V], root *avlNode[K, V], shrinked bool) {
	if n.right == nil {
		max = &avlNode[K, V]{
			split: n.split,
			val:   n.val,
		}
		if n.left != nil {
			n.left.parent = n.parent
			*n = *n.left
		} else {
			if n.parent != nil {
				if n.parent.left == n {
					n.parent.left = nil
				} else {
					n.parent.right = nil
				}
			}
		}
		return max, nil, true
	}
	max, _, _ = n.right.extractMax()
	root, shrinked = n.balanceOnDeletion()
	return max, root, shrinked
}

func (n *avlNode[K, V]) balanceOnDeletion() (root *avlNode[K, V], shrinked bool) {
	bf := n.balanceFactor()
	if bf == 0 {
		return n, false
	} else if bf == -1 || bf == 1 {
		return n, false
	}

	h := n.hight()

	// balance
	if bf < 0 {
		// left-heavy
		if n.left.leftHight() < n.left.rightHight() {
			n.left.rotateLeft()
			root = n.rotateRight()
		} else {
			root = n.rotateRight()
		}
	} else {
		// right-heavy
		if n.right.leftHight() > n.right.rightHight() {
			n.right.rotateRight()
			root = n.rotateLeft()
		} else {
			root = n.rotateLeft()
		}
	}

	if root.hight() < h {
		shrinked = true
	}

	return root, shrinked
}

type AVLTree[K constraints.Ordered, V any] struct {
	root *avlNode[K, V]
}

// NewAVLTree returns a new AVL tree that can contain entries mapping `K` to `V`.
func NewAVLTree[K constraints.Ordered, V any]() *AVLTree[K, V] {
	return &AVLTree[K, V]{}
}

// Insert inserts an entry. When the key already exists, this function return an error.
func (t *AVLTree[K, V]) Insert(key K, value V) error {
	if t.root == nil {
		t.root = newAVLNode(nil, key, value)
		return nil
	}
	root, _, err := t.root.insertAndBalance(key, value)
	if err != nil {
		return err
	}
	if root.parent == nil {
		t.root = root
	}
	return nil
}

// Search earches for an entry having a key that exactly matches a specified key and returns its value.
func (t *AVLTree[K, V]) Search(key K) (value V, found bool) {
	if t.root == nil {
		return
	}
	n, ok := t.root.search(key)
	if !ok {
		return
	}
	return n.val, true
}

// Delete deletes an entry and returns its value.
func (t *AVLTree[K, V]) Delete(key K) (value V, found bool) {
	if t.root == nil {
		return
	}
	root, v, ok := t.root.deleteAndBalance(key)
	if !ok {
		return
	}
	t.root = root
	return v, true
}
