package main

import (
	"container/list"
	"fmt"
)

// GROOOOOOSSSSSS
// DOES NOT APPEAR TO SUPPORT THE NEEDED
// functionality to be a stack.... grrrrr
func main() {
	l := list.New()
	fmt.Println(l)
	l.PushBack(10)
	fmt.Println("res from back", l.Back())
	fmt.Println("len", l.Len())
}
