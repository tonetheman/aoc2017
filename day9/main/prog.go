package main

import "fmt"

func parse(s string) {
	var inGarbage = false
	for index, runeValue := range s {
		fmt.Println(index, runeValue)
		switch runeValue {
		case 60:
			fmt.Println("got a < moving to inGarbage state")
			inGarbage = true
		case 62:
			fmt.Println("got a > moving to a not inGarbage state")
			inGarbage = false
		}
	}
	inGarbage = true
}

func tests() {
	parse("<>")
	//isGarbage("<random characters>")
	//isGarbage("<<<<>")
}

func main() {
	tests()
}
