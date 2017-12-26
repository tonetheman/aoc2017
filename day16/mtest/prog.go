package main

import "fmt"

var programs map[int]string

func pp() {
	p := programs
	for i := 0; i < 5; i++ {
		fmt.Printf(p[i])
	}
	fmt.Println()
}

func shift1() {
	p := programs
	tmp := p[4] // this is the last one
	fmt.Println("tmp is saved", tmp)
	for i := 4; i > 0; i-- {
		fmt.Println("\ti=i-1", i, i-1)
		p[i] = p[i-1]
	}
	p[0] = tmp
}

func shift(count int) {
	for i := 0; i < count; i++ {
		shift1()
	}
}

func exchange(v1, v2 int) {
	programs[v1], programs[v2] = programs[v2], programs[v1]
}

func main() {
	programs = make(map[int]string)
	programs[0] = "a"
	programs[1] = "b"
	programs[2] = "c"
	programs[3] = "d"
	programs[4] = "e"

	pp()
	shift(3)
	pp()
}
