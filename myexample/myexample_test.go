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

func ExampleHelloBye() {
	HelloBye()
	// Output:
	// hello
	// bye
}

func ExampleHelloBye_fail() {
	HelloBye()
	// Output:
	// foo
	// bye
}

func ExampleShuffle() {
	nums := Shuffle([]int{1, 2, 3, 4, 5})
	for _, value := range nums {
		fmt.Println(value)
	}
	// Unordered output:
	// 5
	// 3
	// 2
	// 4
	// 1
}
