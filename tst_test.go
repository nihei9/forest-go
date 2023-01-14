package forest

import (
	"testing"
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
		err = tst.Insert([]rune("hello😺"),  3)
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
		if val, ok := tst.Search([]rune("hello😺")); !ok || val !=  3 {
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
