package main

import "fmt"

func rightCorner(n int) int {
	return (2*n - 1) * (2*n - 1)
}

func bottomMid(n int) int {
	if n == 1 {
		return 1
	}
	rc := rightCorner(n)
	fmt.Println("bm", rc, n)
	return rc - (n - 1)
}

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println(i, rightCorner(i), bottomMid(i))
	}
}
