package forest

import "fmt"

func ExampleTernarySearchTree() {
	keys := [][]rune{
		[]rune("hello"),
		[]rune("world"),
		[]rune("heaven"),
		[]rune("hell"),
		[]rune("healthy"),
	}
	tst := NewTernarySearchTree[rune, int]()
	for i, key := range keys {
		err := tst.Insert(key, i+1)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("All entries:")
	for _, entry := range tst.Entries(nil) {
		fmt.Println(string(entry.Key), entry.Value)
	}

	fmt.Println("Entries with prefix `hea`:")
	for _, entry := range tst.Entries([]rune("hea")) {
		fmt.Println(string(entry.Key), entry.Value)
	}

	fmt.Println("Entries with prefix `hell`:")
	for _, entry := range tst.Entries([]rune("hell")) {
		fmt.Println(string(entry.Key), entry.Value)
	}

	// Output:
	// All entries:
	// healthy 5
	// heaven 3
	// hell 4
	// hello 1
	// world 2
	// Entries with prefix `hea`:
	// healthy 5
	// heaven 3
	// Entries with prefix `hell`:
	// hell 4
	// hello 1
}
