package main

import "fmt"

func parse(s string) {
	var inGarbage = false
	var ignoreNextRune = false

	for index, runeValue := range s {
		fmt.Println(index, runeValue)
		if inGarbage && ignoreNextRune {
			fmt.Println("ignoring from !")
			ignoreNextRune = false
			continue
		}
		switch runeValue {
		case 60:
			if inGarbage {
				fmt.Println("got a < already in garbage: IGNORE")
				continue
			}
			fmt.Println("got a < moving to inGarbage state", inGarbage)
			inGarbage = true
		case 62:
			fmt.Println("got a > moving to a not inGarbage state", inGarbage)
			inGarbage = false
		case 33:
			if inGarbage {
				ignoreNextRune = true
			}
		default:
			if inGarbage {
				fmt.Println("in default", runeValue, "ignoring for inGarbage")
				continue
			}
		}
	}
	inGarbage = true
}

func tests() {
	//parse("<>")
	//parse("<random characters>")
	//parse("<<<<>")
	//parse("<{!>}>")
	//parse("<!!>")
	//parse("<!!!>>")
	parse("<{o\"i!a,<{i<a>")
}

func main() {
	tests()
}
