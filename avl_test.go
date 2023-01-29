package forest

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestAVLTree_Insert(t *testing.T) {
	node := func(split int, value string, left, right *avlNode[int, string]) *avlNode[int, string] {
		n := newAVLNode(nil, split, value)
		if left != nil {
			left.parent = n
			n.left = left
		}
		if right != nil {
			right.parent = n
			n.right = right
		}
		return n
	}

	t.Run("Do left rotation", func(t *testing.T) {
		avl := NewAVLTree[int, string]()
		err := avl.Insert(10, "10")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(11, "11")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(12, "12")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(13, "13")
		if err != nil {
			t.Fatal(err)
		}

		expected := &AVLTree[int, string]{
			root: node(11, "11",
				node(10, "10", nil, nil),
				node(12, "12",
					nil,
					node(13, "13", nil, nil),
				),
			),
		}

		if !reflect.DeepEqual(avl, expected) {
			t.Fatal("unexpected tree")
		}
	})

	t.Run("Do right rotation", func(t *testing.T) {
		avl := NewAVLTree[int, string]()
		err := avl.Insert(10, "10")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(9, "9")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(8, "8")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(7, "7")
		if err != nil {
			t.Fatal(err)
		}

		expected := &AVLTree[int, string]{
			root: node(9, "9",
				node(8, "8",
					node(7, "7", nil, nil),
					nil,
				),
				node(10, "10", nil, nil),
			),
		}

		if !reflect.DeepEqual(avl, expected) {
			t.Fatal("unexpected tree")
		}
	})

	t.Run("Do left-right rotation", func(t *testing.T) {
		avl := NewAVLTree[int, string]()
		err := avl.Insert(10, "10")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(11, "11")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(7, "7")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(6, "6")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(8, "8")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(9, "9")
		if err != nil {
			t.Fatal(err)
		}

		expected := &AVLTree[int, string]{
			root: node(8, "8",
				node(7, "7",
					node(6, "6", nil, nil),
					nil,
				),
				node(10, "10",
					node(9, "9", nil, nil),
					node(11, "11", nil, nil),
				),
			),
		}

		if !reflect.DeepEqual(avl, expected) {
			t.Fatal("unexpected tree")
		}
	})

	t.Run("Do right-left rotation", func(t *testing.T) {
		avl := NewAVLTree[int, string]()
		err := avl.Insert(10, "10")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(9, "9")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(13, "13")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(12, "12")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(14, "14")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(11, "11")
		if err != nil {
			t.Fatal(err)
		}

		expected := &AVLTree[int, string]{
			root: node(12, "12",
				node(10, "10",
					node(9, "9", nil, nil),
					node(11, "11", nil, nil),
				),
				node(13, "13",
					nil,
					node(14, "14", nil, nil),
				),
			),
		}

		if !reflect.DeepEqual(avl, expected) {
			t.Fatal("unexpected tree")
		}
	})

	t.Run("When keys are duplicated, an error occurs", func(t *testing.T) {
		avl := NewAVLTree[string, int]()
		err := avl.Insert("hello", 0)
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert("hello", 0)
		if err == nil {
			t.Fatal("error must occur")
		}
	})
}

func TestAVLTree_Search(t *testing.T) {
	t.Run("Do binary search", func(t *testing.T) {
		avl := NewAVLTree[int, string]()
		err := avl.Insert(10, "10")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(9, "9")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(13, "13")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(12, "12")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(14, "14")
		if err != nil {
			t.Fatal(err)
		}
		err = avl.Insert(11, "11")
		if err != nil {
			t.Fatal(err)
		}

		if val, ok := avl.Search(9); !ok || val != "9" {
			t.Fatalf("unexpected result. want: 9, true, got: %v, %v", val, ok)
		}
		if val, ok := avl.Search(10); !ok || val != "10" {
			t.Fatalf("unexpected result. want: 10, true, got: %v, %v", val, ok)
		}
		if val, ok := avl.Search(11); !ok || val != "11" {
			t.Fatalf("unexpected result. want: 11, true, got: %v, %v", val, ok)
		}
		if val, ok := avl.Search(12); !ok || val != "12" {
			t.Fatalf("unexpected result. want: 12, true, got: %v, %v", val, ok)
		}
		if val, ok := avl.Search(13); !ok || val != "13" {
			t.Fatalf("unexpected result. want: 13, true, got: %v, %v", val, ok)
		}
		if val, ok := avl.Search(14); !ok || val != "14" {
			t.Fatalf("unexpected result. want: 14, true, got: %v, %v", val, ok)
		}
	})

	t.Run("When a key is not exist, the Search method returns false", func(t *testing.T) {
		avl := NewAVLTree[int, string]()
		err := avl.Insert(10, "10")
		if err != nil {
			t.Fatal(err)
		}

		if val, ok := avl.Search(9); ok || val == "9" {
			t.Fatalf("unexpected result. want:\"\", false, got: %v, %v", val, ok)
		}
		if val, ok := avl.Search(11); ok || val == "11" {
			t.Fatalf("unexpected result. want:\"\", false, got: %v, %v", val, ok)
		}
	})
}

func TestAVLTree_Delete(t *testing.T) {
	node := func(split int, value string, left, right *avlNode[int, string]) *avlNode[int, string] {
		n := newAVLNode(nil, split, value)
		if left != nil {
			left.parent = n
			n.left = left
		}
		if right != nil {
			right.parent = n
			n.right = right
		}
		return n
	}

	tests := []struct {
		entries  []int
		delete   int
		expected *avlNode[int, string]
	}{
		// delete a root and don't rotate

		// 10
		{
			entries:  []int{10},
			delete:   10,
			expected: nil,
		},
		//   10
		//  /
		// 9
		{
			entries:  []int{10, 9},
			delete:   10,
			expected: node(9, "9", nil, nil),
		},
		// 10
		//   \
		//    11
		{
			entries:  []int{10, 11},
			delete:   10,
			expected: node(11, "11", nil, nil),
		},
		//   10
		//  /  \
		// 9    11
		{
			entries: []int{10, 9, 11},
			delete:  10,
			expected: node(9, "9",
				nil,
				node(11, "11", nil, nil)),
		},
		//   10
		//  /  \
		// 8    11
		//  \
		//   9
		{
			entries: []int{10, 8, 11, 9},
			delete:  10,
			expected: node(9, "9",
				node(8, "8", nil, nil),
				node(11, "11", nil, nil)),
		},
		//     10
		//    /  \
		//   8    11
		//  / \
		// 7   9
		{
			entries: []int{10, 8, 11, 7, 9},
			delete:  10,
			expected: node(9, "9",
				node(8, "8",
					node(7, "7", nil, nil),
					nil),
				node(11, "11", nil, nil)),
		},
		//     10
		//    /  \
		//   7    11
		//  / \    \
		// 6   9    12
		//    /
		//   8
		{
			entries: []int{10, 7, 11, 6, 9, 12, 8},
			delete:  10,
			expected: node(9, "9",
				node(7, "7",
					node(6, "6", nil, nil),
					node(8, "8", nil, nil)),
				node(11, "11",
					nil,
					node(12, "12", nil, nil))),
		},

		// delete a root's left child and don't rotate

		//   10
		//  /
		// 9
		{
			entries: []int{10, 9}, // (10 (9 nil nil) nil)
			delete:  9,
			expected: node(10, "10",
				nil,
				nil),
		},
		//     10
		//    /  \
		//   9    11
		//  /
		// 8
		{
			entries: []int{10, 9, 11, 8},
			delete:  9,
			expected: node(10, "10",
				node(8, "8", nil, nil),
				node(11, "11", nil, nil)),
		},
		//   10
		//  /  \
		// 8    11
		//  \
		//   9
		{
			entries: []int{10, 8, 11, 9},
			delete:  8,
			expected: node(10, "10",
				node(9, "9", nil, nil),
				node(11, "11", nil, nil)),
		},
		//     10
		//    /  \
		//   8    11
		//  / \
		// 7   9
		{
			entries: []int{10, 8, 11, 7, 9},
			delete:  8,
			expected: node(10, "10",
				node(7, "7",
					nil,
					node(9, "9", nil, nil)),
				node(11, "11", nil, nil)),
		},
		//     10
		//    /  \
		//   8    11
		//  / \    \
		// 6   9    12
		//  \
		//   7
		{
			entries: []int{10, 8, 11, 6, 9, 12, 7},
			delete:  8,
			expected: node(10, "10",
				node(7, "7",
					node(6, "6", nil, nil),
					node(9, "9", nil, nil)),
				node(11, "11",
					nil,
					node(12, "12", nil, nil))),
		},

		// delete a root's right child and don't rotate

		// 10
		//   \
		//    11
		{
			entries: []int{10, 11},
			delete:  11,
			expected: node(10, "10",
				nil,
				nil),
		},
		//   10
		//  /  \
		// 9    12
		//     /
		//   11
		{
			entries: []int{10, 9, 12, 11},
			delete:  12,
			expected: node(10, "10",
				node(9, "9", nil, nil),
				node(11, "11", nil, nil)),
		},
		//   10
		//  /  \
		// 9    11
		//       \
		//        12
		{
			entries: []int{10, 9, 11, 12},
			delete:  11,
			expected: node(10, "10",
				node(9, "9", nil, nil),
				node(12, "12", nil, nil)),
		},
		//   10
		//  /  \
		// 9    12
		//     /  \
		//   11    13
		{
			entries: []int{10, 9, 12, 11, 13},
			delete:  12,
			expected: node(10, "10",
				node(9, "9", nil, nil),
				node(11, "11",
					nil,
					node(13, "13", nil, nil))),
		},
		//     10
		//    /  \
		//   9    13
		//  /    /  \
		// 8   11    14
		//       \
		//        12
		{
			entries: []int{10, 9, 13, 8, 11, 14, 12},
			delete:  13,
			expected: node(10, "10",
				node(9, "9",
					node(8, "8", nil, nil),
					nil),
				node(12, "12",
					node(11, "11", nil, nil),
					node(14, "14", nil, nil))),
		},

		// delete a node with no children and rotate

		// left rotation
		//
		//   10
		//  /  \
		// 9    11
		//       \
		//        12
		{
			entries: []int{10, 9, 11, 12},
			delete:  9,
			expected: node(11, "11",
				node(10, "10", nil, nil),
				node(12, "12", nil, nil)),
		},
		// right rotation
		//
		//     10
		//    /  \
		//   9    11
		//  /
		// 8
		{
			entries: []int{10, 9, 11, 8},
			delete:  11,
			expected: node(9, "9",
				node(8, "8", nil, nil),
				node(10, "10", nil, nil)),
		},
		// right-left rotation
		//
		//   10
		//  /  \
		// 9    12
		//     /
		//   11
		{
			entries: []int{10, 9, 12, 11},
			delete:  9,
			expected: node(11, "11",
				node(10, "10", nil, nil),
				node(12, "12", nil, nil)),
		},
		// left-right rotation
		//
		//   10
		//  /  \
		// 8    11
		//  \
		//   9
		{
			entries: []int{10, 8, 11, 9}, // (10 (8 nil (9 nil nil)) (11 nil nil))
			delete:  11,
			expected: node(9, "9",
				node(8, "8", nil, nil),
				node(10, "10", nil, nil)),
		},

		// delete a node with a left child and rotate

		// left rotation
		//
		//     10
		//    /  \
		//   9    12
		//  /    /  \
		// 8   11    13
		//            \
		//             14
		{
			entries: []int{10, 9, 12, 8, 11, 13, 14},
			delete:  9,
			expected: node(12, "12",
				node(10, "10",
					node(8, "8", nil, nil),
					node(11, "11", nil, nil)),
				node(13, "13",
					nil,
					node(14, "14", nil, nil))),
		},
		// right rotation
		//
		//       10
		//      /  \
		//     8    12
		//    / \   /
		//   7   9 11
		//  /
		// 6
		{
			entries: []int{10, 8, 12, 7, 9, 11, 6},
			delete:  12,
			expected: node(8, "8",
				node(7, "7",
					node(6, "6", nil, nil),
					nil),
				node(10, "10",
					node(9, "9", nil, nil),
					node(11, "11", nil, nil))),
		},
		// right-left rotation
		//
		//     10
		//    /  \
		//   9    13
		//  /    /  \
		// 8   11    14
		//       \
		//        12
		{
			entries: []int{10, 9, 13, 8, 11, 14, 12},
			delete:  9,
			expected: node(11, "11",
				node(10, "10",
					node(8, "8", nil, nil),
					nil),
				node(13, "13",
					node(12, "12", nil, nil),
					node(14, "14", nil, nil))),
		},
		// left-right rotation
		//
		//       10
		//      /  \
		//     7    12
		//    / \   /
		//   6   9 11
		//      /
		//     8
		{
			entries: []int{10, 7, 12, 6, 9, 11, 8},
			delete:  12,
			expected: node(9, "9",
				node(7, "7",
					node(6, "6", nil, nil),
					node(8, "8", nil, nil)),
				node(10, "10",
					nil,
					node(11, "11", nil, nil))),
		},

		// delete a node with a right child and rotate

		// left rotation
		//
		//     10
		//    /  \
		//   8    12
		//    \   / \
		//     9 11  13
		//            \
		//             14
		{
			entries: []int{10, 8, 12, 9, 11, 13, 14},
			delete:  8,
			expected: node(12, "12",
				node(10, "10",
					node(9, "9", nil, nil),
					node(11, "11", nil, nil)),
				node(13, "13",
					nil,
					node(14, "14", nil, nil))),
		},
		// right rotation
		//
		//       10
		//      /  \
		//     8    11
		//    / \    \
		//   7   9    12
		//  /
		// 6
		{
			entries: []int{10, 8, 11, 7, 9, 12, 6},
			delete:  11,
			expected: node(8, "8",
				node(7, "7",
					node(6, "6", nil, nil),
					nil),
				node(10, "10",
					node(9, "9", nil, nil),
					node(12, "12", nil, nil))),
		},
		// right-left rotation
		//
		//   10
		//  /  \
		// 8    13
		//  \   / \
		//   9 11  14
		//      \
		//       12
		{
			entries: []int{10, 8, 13, 9, 11, 14, 12},
			delete:  8,
			expected: node(11, "11",
				node(10, "10",
					node(9, "9", nil, nil),
					nil),
				node(13, "13",
					node(12, "12", nil, nil),
					node(14, "14", nil, nil))),
		},
		// left-right rotation
		//
		//     10
		//    /  \
		//   7    11
		//  / \    \
		// 6   9    12
		//    /
		//   8
		{
			entries: []int{10, 7, 11, 6, 9, 12, 8},
			delete:  11,
			expected: node(9, "9",
				node(7, "7",
					node(6, "6", nil, nil),
					node(8, "8", nil, nil)),
				node(10, "10",
					nil,
					node(12, "12", nil, nil))),
		},

		// rotate a left-sub tree of deleted node (rotation on `extractMax`)

		// right rotation
		//
		//       10
		//      /  \
		//     8    11
		//    / \    \
		//   7   9    12
		//  /
		// 6
		{
			entries: []int{10, 8, 11, 7, 9, 12, 6},
			delete:  10,
			expected: node(9, "9",
				node(7, "7",
					node(6, "6", nil, nil),
					node(8, "8", nil, nil)),
				node(11, "11",
					nil,
					node(12, "12", nil, nil))),
		},
		// left-right rotation
		//
		//     10
		//    /  \
		//   8    11
		//  / \    \
		// 6   9    12
		//  \
		//   7
		{
			entries: []int{10, 8, 11, 6, 9, 12, 7},
			delete:  10,
			expected: node(9, "9",
				node(7, "7",
					node(6, "6", nil, nil),
					node(8, "8", nil, nil)),
				node(11, "11",
					nil,
					node(12, "12", nil, nil))),
		},
		// right rotation
		//
		//          10
		//        /    \
		//       7      12
		//      / \     / \
		//     5   8   11  13
		//    / \   \       \
		//   4   6   9       14
		//  /
		// 3
		{
			entries: []int{10, 7, 12, 5, 8, 11, 13, 4, 6, 9, 14, 3},
			delete:  10,
			expected: node(9, "9",
				node(5, "5",
					node(4, "4",
						node(3, "3", nil, nil),
						nil),
					node(7, "7",
						node(6, "6", nil, nil),
						node(8, "8", nil, nil))),
				node(12, "12",
					node(11, "11", nil, nil),
					node(13, "13",
						nil,
						node(14, "14", nil, nil)))),
		},
		// left-right rotation
		//
		//        10
		//      /    \
		//     7      12
		//    / \     / \
		//   4   8   11  13
		//  / \   \       \
		// 3   6   9       14
		//  \
		//   5
		{
			entries: []int{10, 7, 12, 4, 8, 11, 13, 3, 6, 9, 14, 5},
			delete:  10,
			expected: node(9, "9",
				node(6, "6",
					node(4, "4",
						node(3, "3", nil, nil),
						node(5, "5", nil, nil)),
					node(7, "7",
						nil,
						node(8, "8", nil, nil))),
				node(12, "12",
					node(11, "11", nil, nil),
					node(13, "13",
						nil,
						node(14, "14", nil, nil)))),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%v", i), func(t *testing.T) {
			avl := NewAVLTree[int, string]()
			for _, e := range tt.entries {
				err := avl.Insert(e, strconv.Itoa(e))
				if err != nil {
					t.Fatal(err)
				}
			}
			v := strconv.Itoa(tt.delete)
			if val, ok := avl.Delete(tt.delete); !ok || val != v {
				t.Fatalf("unexpected result. want: %+v, true, got: %v, %v", v, val, ok)
			}
			e := &AVLTree[int, string]{
				root: tt.expected,
			}
			if !reflect.DeepEqual(avl, e) {
				t.Fatal("unexpected tree")
			}
		})
	}
}
