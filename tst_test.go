package forest

import (
	"math"
	"reflect"
	"sort"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestTernarySearchTree_Insert(t *testing.T) {
	t.Run("The tree can contain different keys", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		err := tst.Insert([]rune("hello"), 0)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("world"), 0)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("The tree can contain keys with the same prefix", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		err := tst.Insert([]rune("hello"), 0)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("hell"), 0)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("hello😺"), 0)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("heaven"), 0)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("When keys are duplicated, an error occurs", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		err := tst.Insert([]rune("hello"), 0)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("hello"), 0)
		if err == nil {
			t.Fatal("error must occur")
		}
	})

	t.Run("A key must not be empty", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		err := tst.Insert([]rune(""), 0)
		if err == nil {
			t.Fatal("error must occur")
		}
		err = tst.Insert([]rune{}, 0)
		if err == nil {
			t.Fatal("error must occur")
		}
		err = tst.Insert(nil, 0)
		if err == nil {
			t.Fatal("error must occur")
		}
	})
}

func TestTernarySearchTree_Search(t *testing.T) {
	t.Run("The tree can contain different keys", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		if err := tst.Insert([]rune("hello"), 1); err != nil {
			t.Fatal(err)
		}
		if err := tst.Insert([]rune("world"), 2); err != nil {
			t.Fatal(err)
		}

		if val, ok := tst.Search([]rune("hello")); !ok || val != 1 {
			t.Fatalf("unexpected result. want: 1, true, got: %v, %v", val, ok)
		}
		if val, ok := tst.Search([]rune("world")); !ok || val != 2 {
			t.Fatalf("unexpected result. want: 2, true, got: %v, %v", val, ok)
		}
	})

	t.Run("The tree can contain keys with the same prefix", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		err := tst.Insert([]rune("hello"), 1)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("hell"), 2)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("hello😺"), 3)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("heaven"), 4)
		if err != nil {
			t.Fatal(err)
		}

		if val, ok := tst.Search([]rune("hello")); !ok || val != 1 {
			t.Fatalf("unexpected result. want: 1, true, got: %v, %v", val, ok)
		}
		if val, ok := tst.Search([]rune("hell")); !ok || val != 2 {
			t.Fatalf("unexpected result. want: 2, true, got: %v, %v", val, ok)
		}
		if val, ok := tst.Search([]rune("hello😺")); !ok || val != 3 {
			t.Fatalf("unexpected result. want: 3, true, got: %v, %v", val, ok)
		}
		if val, ok := tst.Search([]rune("heaven")); !ok || val != 4 {
			t.Fatalf("unexpected result. want: 4, true, got: %v, %v", val, ok)
		}
	})

	t.Run("Keys with the same suffix, but different prefixes, do not match", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		if err := tst.Insert([]rune("cat"), 1); err != nil {
			t.Fatal(err)
		}

		if val, ok := tst.Search([]rune("at")); ok || val == 1 {
			t.Fatalf("why found")
		}
	})

	t.Run("A key must not be empty", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		err := tst.Insert([]rune("hello"), 1)
		if err != nil {
			t.Fatal(err)
		}

		if val, ok := tst.Search([]rune("")); ok || val == 1 {
			t.Fatalf("unexpected result. want: 0, false, got: %v, %v", val, ok)
		}
		if val, ok := tst.Search([]rune{}); ok || val == 1 {
			t.Fatalf("unexpected result. want: 0, false, got: %v, %v", val, ok)
		}
		if val, ok := tst.Search(nil); ok || val == 1 {
			t.Fatalf("unexpected result. want: 0, false, got: %v, %v", val, ok)
		}
	})
}

func TestTernarySearchTree_List(t *testing.T) {
	prettier := func(runeSeq []rune) string {
		return string(runeSeq)
	}

	t.Run("The tree can contain different keys", func(t *testing.T) {
		keys := [][]rune{
			[]rune("hello"),
			[]rune("world"),
		}
		tst := NewTernarySearchTree[rune, int]()
		for i, key := range keys {
			if err := tst.Insert(key, i+1); err != nil {
				t.Fatal(err)
			}
		}

		list := tst.List(nil)
		testTSTListResult(t, list, keys, prettier)
	})

	t.Run("The tree can contain keys with the same prefix", func(t *testing.T) {
		keys := [][]rune{
			[]rune("healthy"),
			[]rune("hello"),
			[]rune("world"),
			[]rune("hell"),
			[]rune("hello😺"),
			[]rune("heaven"),
		}
		tst := NewTernarySearchTree[rune, int]()
		for i, key := range keys {
			if err := tst.Insert(key, i+1); err != nil {
				t.Fatal(err)
			}
		}

		{
			list := tst.List(nil)
			testTSTListResult(t, list, keys, prettier)
		}
		{
			expected := [][]rune{
				[]rune("healthy"),
				[]rune("heaven"),
			}
			list := tst.List([]rune("hea"))
			testTSTListResult(t, list, expected, prettier)
		}
		{
			expected := [][]rune{
				[]rune("hello"),
				[]rune("hell"),
				[]rune("hello😺"),
			}
			list := tst.List([]rune("hell"))
			testTSTListResult(t, list, expected, prettier)
		}
	})

	t.Run("When the tree is empty, the result of list operation is also empty", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()

		list := tst.List(nil)
		if len(list) != 0 {
			t.Fatalf("result must be empty")
		}
	})

	t.Run("A too-long prefix does not match any keys", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		if err := tst.Insert([]rune("foo"), 1); err != nil {
			t.Fatal(err)
		}

		list := tst.List([]rune("fooo"))
		if len(list) != 0 {
			t.Fatalf("result must be empty")
		}
	})
}

func TestTernarySearchTree_Delete(t *testing.T) {
	t.Run("The tree can contain different keys", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		if err := tst.Insert([]rune("hello"), 1); err != nil {
			t.Fatal(err)
		}
		if err := tst.Insert([]rune("world"), 2); err != nil {
			t.Fatal(err)
		}

		if val, ok := tst.Delete([]rune("hello")); !ok || val != 1 {
			t.Fatalf("unexpected result. want: 1, true, got: %v, %v", val, ok)
		}
		if val, ok := tst.Delete([]rune("world")); !ok || val != 2 {
			t.Fatalf("unexpected result. want: 2, true, got: %v, %v", val, ok)
		}
	})

	t.Run("The tree can contain keys with the same prefix", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		err := tst.Insert([]rune("hello"), 1)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("hell"), 2)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("hello😺"), 3)
		if err != nil {
			t.Fatal(err)
		}
		err = tst.Insert([]rune("heaven"), 4)
		if err != nil {
			t.Fatal(err)
		}

		if val, ok := tst.Delete([]rune("hello")); !ok || val != 1 {
			t.Fatalf("unexpected result. want: 1, true, got: %v, %v", val, ok)
		}
		if val, ok := tst.Delete([]rune("hell")); !ok || val != 2 {
			t.Fatalf("unexpected result. want: 2, true, got: %v, %v", val, ok)
		}
		if val, ok := tst.Delete([]rune("hello😺")); !ok || val != 3 {
			t.Fatalf("unexpected result. want: 3, true, got: %v, %v", val, ok)
		}
		if val, ok := tst.Delete([]rune("heaven")); !ok || val != 4 {
			t.Fatalf("unexpected result. want: 4, true, got: %v, %v", val, ok)
		}
	})

	t.Run("Keys with the same suffix, but different prefixes, do not match", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		if err := tst.Insert([]rune("cat"), 1); err != nil {
			t.Fatal(err)
		}

		if val, ok := tst.Delete([]rune("at")); ok || val == 1 {
			t.Fatalf("why found")
		}
	})

	t.Run("A key must not be empty", func(t *testing.T) {
		tst := NewTernarySearchTree[rune, int]()
		err := tst.Insert([]rune("hello"), 1)
		if err != nil {
			t.Fatal(err)
		}

		if val, ok := tst.Delete([]rune("")); ok || val == 1 {
			t.Fatalf("unexpected result. want: 0, false, got: %v, %v", val, ok)
		}
		if val, ok := tst.Delete([]rune{}); ok || val == 1 {
			t.Fatalf("unexpected result. want: 0, false, got: %v, %v", val, ok)
		}
		if val, ok := tst.Delete(nil); ok || val == 1 {
			t.Fatalf("unexpected result. want: 0, false, got: %v, %v", val, ok)
		}
	})
}

func testTSTListResult[K constraints.Ordered, P any](t *testing.T, actual, expected [][]K, prettier func([]K) P) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Fatalf("unexpected list length. want: %v, got: %v", len(expected), len(actual))
	}
	sortTSTList(actual)
	sortTSTList(expected)
	for i, a := range actual {
		if !reflect.DeepEqual(a, expected[i]) {
			t.Fatalf("unexpected result. want: %v, got: %v", prettier(expected[i]), prettier(a))
		}
	}
}

func sortTSTList[K constraints.Ordered](list [][]K) {
	sort.Slice(list, func(i, j int) bool {
		minLen := int(math.Min(float64(len(list[i])), float64(len(list[j]))))
		for k := 0; k < minLen; k++ {
			if list[i][k] != list[j][k] {
				return list[i][k] < list[j][k]
			}
		}
		return len(list[i]) < len(list[j])
	})
}
