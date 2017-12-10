package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}
func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}
func (s stack) Peek() int {
	l := len(s)
	if l == 0 {
		return 0
	}
	return s[l-1]
}

func parse(s string) {
	var inGarbage = false
	var ignoreNextRune = false
	var groupCount = 0

	// make a stack needed to keep up with groups
	st := make(stack, 0)

	for index, runeValue := range s {
		fmt.Println("NEXTCHAR:", index, runeValue, string(runeValue))

		if inGarbage && ignoreNextRune {
			fmt.Println("inGarbage and ignoreNextRune is true ignoring ! and setting ignoreNext false")
			ignoreNextRune = false
			continue
		}

		// GARBAGE NOW
		if inGarbage {
			switch runeValue {
			case 62:
				inGarbage = false
				fmt.Println("got a > moving to a not inGarbage state", inGarbage)
			case 33:
				fmt.Println("got a ! setting ignoreNextRune to true")
				ignoreNextRune = true
			default:
				fmt.Println("in default", runeValue, "ignoring for inGarbage")
				continue

			}
		}

		// NOT in GARBAGE DOWN HERE
		switch runeValue {
		case 123:
			// this is the start of a group
			currentTop := st.Peek()
			if currentTop == 0 {
				st = st.Push(1)
			} else {
				st = st.Push(currentTop + 1)
			}
		case 125:
			var val int
			// this is the end of a group
			st, val = st.Pop()
			groupCount += val
		case 60:
			fmt.Println("got a < moving to inGarbage state", inGarbage)
			inGarbage = true
		default:
			fmt.Println("in default")
		}
	}

	fmt.Println("end of parse groupCount", groupCount)
}

func tests() {
	//parse("<>")
	//parse("<random characters>")
	//parse("<<<<>")
	//parse("<{!>}>")
	//parse("<!!>")
	//parse("<!!!>>")
	//parse("<{o\"i!a,<{i<a>")
	//parse("{}") // 1
	//parse("{{{}}}") // 6
	//parse("{{},{}}") // 5
	//parse("{{{},{},{{}}}}") //16
	//parse("{<a>,<a>,<a>,<a>}") //1
	//parse("{{<ab>},{<ab>},{<ab>},{<ab>}}") //9
	//parse("{{<!!>},{<!!>},{<!!>},{<!!>}}") // 9
	//parse("{{<a!>},{<a!>},{<a!>},{<ab>}}") //3
}

func part1(filename string) {
	// need to read from file gross parsing avoided
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	// make a scanner to read in line by line
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(len(text))
		fmt.Println("starting parse...")
		parse(text)
	}
}

func main() {
	//tests()
	part1("main\\part1.input")
}
