package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(filename string) []int {
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	vals := make([]int, 0)
	// make a scanner to read in line by line
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		val, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("err in convert", err)
			continue
		}
		vals = append(vals, val)
		//fmt.Println(val, err)
	}
	//fmt.Println("final vals", vals)
	return vals
}

func runProgram(vals []int) {
	debug_count := 0
	ip := 0 // instruction pointer
	for {
		old_ip := ip
		current_jmp := vals[ip]

		// console log we are jmping
		//fmt.Println("jmp with offset", current_jmp)
		// make the jmp
		ip += current_jmp
		//fmt.Println("ip is now pointing to", ip)
		// increment the val at oldip
		vals[old_ip]++

		//fmt.Println("current state", vals)
		if ip >= len(vals) {
			fmt.Println("ESCAPED", ip, debug_count)
			break
		}
		//if debug_count == 4 {
		//	fmt.Println("debug break")
		//	fmt.Println("len of vals", len(vals))
		//	break
		//}
		//fmt.Println("----------------")
		debug_count++
		if debug_count%100 == 0 {
			fmt.Print("*")
		}
	}

	fmt.Println("You made it, total steps", debug_count+1)
}

func main() {
	vals := readInput("part1.input")
	fmt.Println("puzzle input", vals)
	runProgram(vals)
}
