package main

import "fmt"

func parse(s string) {
	var inGarbage = false
	var ignoreNextRune = false

	for index, runeValue := range s {
		fmt.Println(index, runeValue)
		if inGarbage && ignoreNextRune {
			fmt.Println("inGarbage and ignoreNextRune is true")
			fmt.Println("ignoring !")
			ignoreNextRune = false
			continue
		}

		if inGarbage {
			switch runeValue {
			case 62:
				inGarbage = false
				fmt.Println("got a > moving to a not inGarbage state", inGarbage)
			case 33:
				ignoreNextRune = true
			default:
				fmt.Println("in default", runeValue, "ignoring for inGarbage")
				continue

			}
		}

		// NOT in GARBAGE DOWN HERE
		switch runeValue {
		case 60:
			fmt.Println("got a < moving to inGarbage state", inGarbage)
			inGarbage = true
		default:
			fmt.Println("in default")
		}
	}
	inGarbage = true
}

func tests() {
	parse("<>")
	//parse("<random characters>")
	//parse("<<<<>")
	//parse("<{!>}>")
	//parse("<!!>")
	//parse("<!!!>>")
	//parse("<{o\"i!a,<{i<a>")
}

func main() {
	tests()
}
