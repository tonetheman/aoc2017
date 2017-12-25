package main

import "fmt"

const A = 98
const B = 99
const tape_size = 10

type tape [tape_size]int

func part1() {
	var pos int = tape_size / 2
	var state int = A
	var t tape

	count := 0
	for {
		fmt.Println("before step", t, pos, state)
		if state == A {
			if t[pos] == 0 {
				t[pos] = 1
				pos++
				state = B
			} else if t[pos] == 1 {
				t[pos] = 0
				pos--
				state = B
			}
		} else if state == B {

		} else {
			fmt.Println("ERR: illegal state")
			break
		}

		fmt.Println("after step", t, pos, state)
		count++
		if count == 1 {
			break
		}
	}
}

func main() {
	part1()
}
