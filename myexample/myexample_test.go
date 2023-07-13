// myexample_test.go
package myexample

import "fmt"

func ExampleHello() {
	Hello()
	// Output: hello
}

func ExampleHello_fail() {
	Hello()
	// Output: foo
}

func ExampleSalutations() {
	Salutations()
	// Output:
	// hello
	// goodbye
}

func ExampleSalutations_fail() {
	Salutations()
	// Output:
	// foo
	// goodbye
}

func ExampleShuffle() {
	for _, value := range Shuffle([]int{1, 2, 3, 4, 5}) {
		fmt.Println(value)
	}
	// Unordered output: 5
	// 3
	// 2
	// 4
	// 1
}
