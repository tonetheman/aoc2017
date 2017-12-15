package main

import "fmt"

func main() {

	const aFactor = 16807
	const bFactor = 48271

	const aStart = 65
	const bStart = 8921

	fmt.Println(aFactor * aStart, bFactor * bStart)

}
