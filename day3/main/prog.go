package main

import "fmt"

func rightCorner(n int) int {
	return (2*n - 1) * (2*n - 1)
}

// NOT RIGHT
func bottomMid(n int) int {
	rc := rightCorner(n)
	return rc - (n-1)*7
}

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println(i, rightCorner(i), bottomMid(i))
	}
}
