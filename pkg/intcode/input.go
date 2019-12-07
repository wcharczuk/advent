package intcode

import "fmt"

// ReadInt reads an int from stdin.
func ReadInt() int {
	fmt.Print("Input: ")
	var i int
	fmt.Scanf("%d", &i)
	return i
}
