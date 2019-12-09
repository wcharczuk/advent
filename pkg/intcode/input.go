package intcode

import "fmt"

// ReadInt reads an int from stdin.
func ReadInt() int64 {
	fmt.Print("Input: ")
	var i int64
	fmt.Scanf("%d", &i)
	return i
}
