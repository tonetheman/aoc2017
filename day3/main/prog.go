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
	//fmt.Println("bm", rc, n)
	return rc - (n - 1)
}

func rightMid(n int) int {
	if n == 1 {
		return 1
	}
	rc := rightCorner(n)
	return rc - (n-1)*7
}

func topMid(n int) int {
	if n == 1 {
		return 1
	}
	rc := rightCorner(n)
	return rc - (n-1)*5
}

func leftMid(n int) int {
	if n == 1 {
		return 1
	}
	rc := rightCorner(n)
	return rc - (n-1)*3
}

func main() {
	i := 1
	const target int = 325489
	for {
		fmt.Println(i, rightCorner(i), leftMid(i), bottomMid(i), topMid(i), rightMid(i))
		if target < rightCorner(i) {
			fmt.Println("stopping...")
			break
		}
		i++
	}

}
