package slices

import "fmt"

func ExampleContains() {
	ok := Contains([]int{10, 20, 30}, 10)
	fmt.Println(ok)

	ok = Contains([]int{20, 30, 40}, 10)
	fmt.Println(ok)

	// Output:
	// true
	// false
}
