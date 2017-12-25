package main

import "fmt"

const A = 98
const B = 99
const C = 100
const D = 101
const E = 102
const F = 103

const tape_size = 12800 * 2

type tape [tape_size]int

func test() {
	var pos int = tape_size / 2
	var state int = A
	var t tape

	count := 0
	for {
		fmt.Println(count, "before step", t, pos, state)
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
			if t[pos] == 0 {
				t[pos] = 1
				pos--
				state = A
			} else if t[pos] == 1 {
				t[pos] = 1
				pos++
				state = A
			}
		} else {
			fmt.Println("ERR: illegal state")
			break
		}

		fmt.Println(count, "after step", t, pos, state)
		count++
		if count == 6 {
			break
		}
	}

	total := 0
	for i := 0; i < len(t); i++ {
		total += t[i]
	}
	fmt.Println("chksum", total)
}

func part1() {
	var pos int = tape_size / 2
	var state int = A
	var t tape

	count := 0
	for {
		//fmt.Println(count, "before step", t, pos, state)
		if state == A {
			if t[pos] == 0 {
				t[pos] = 1
				pos++
				state = B
			} else if t[pos] == 1 {
				t[pos] = 1
				pos--
				state = E
			}
		} else if state == B {
			if t[pos] == 0 {
				t[pos] = 1
				pos++
				state = C
			} else if t[pos] == 1 {
				t[pos] = 1
				pos++
				state = F
			}
		} else if state == C {
			if t[pos] == 0 {
				t[pos] = 1
				pos--
				state = D
			} else if t[pos] == 1 {
				t[pos] = 0
				pos++
				state = B
			}
		} else if state == D {
			if t[pos] == 0 {
				t[pos] = 1
				pos++
				state = E
			} else if t[pos] == 1 {
				t[pos] = 0
				pos--
				state = C
			}
		} else if state == E {
			if t[pos] == 0 {
				t[pos] = 1
				pos--
				state = A
			} else if t[pos] == 1 {
				t[pos] = 0
				pos++
				state = D
			}
		} else if state == F {
			if t[pos] == 0 {
				t[pos] = 1
				pos++
				state = A
			} else if t[pos] == 1 {
				t[pos] = 1
				pos++
				state = C
			}
		} else {
			fmt.Println("ERR: illegal state")
			break
		}

		//fmt.Println(count, "after step", t, pos, state)
		count++
		if count%10000 == 0 {
			fmt.Printf("*")
		}
		if count == 12523873 {
			break
		}
	}

	total := 0
	for i := 0; i < len(t); i++ {
		total += t[i]
	}
	fmt.Println("chksum", total)
}

func main() {
	//test()
	part1()
}
