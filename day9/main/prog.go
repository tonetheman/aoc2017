package main

import "fmt"

func isGarbage(s string) bool {
	var inG bool = false
	for c, runeValue := range s {
		fmt.Println(c, runeValue)
		if runeValue == 60 {
			fmt.Println("inG set to true, in garbage")
			inG = true
			continue
		}
		if runeValue == 62 {
			fmt.Println("inG is false out of gargage")
			inG = false
		}
		if inG {
			continue
		}
	}
	return false
}

func tests() {
	//isGarbage("<>")
	//isGarbage("<random characters>")
	isGarbage("<<<<>")
}

func main() {
	tests()
}
