package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func spin(a []string, count int) []string {
	for i := 0; i < count; i++ {
		ll := len(a) - 1
		last := a[ll]
		//fmt.Println("spindbg", last)
		a = a[0:ll]
		//fmt.Println("spindbg", a)
		a = append([]string{last}, a...)
		//fmt.Println(a)
	}
	return a
}

func exchange(a []string, v1, v2 int) []string {
	a[v1], a[v2] = a[v2], a[v1]
	return a
}

func partner(a []string, p1, p2 string) []string {
	var pos1 = 0
	var pos2 = 0
	for i := 0; i < len(a); i++ {
		if a[i] == p1 {
			pos1 = i
		}
		if a[i] == p2 {
			pos2 = i
		}
	}
	return exchange(a, pos1, pos2)
}

func part1(filename string, programs []string) {
	// open the file
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	//programs := []string{"a", "b", "c", "d", "e"}

	// make a scanner to read in line by line
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		//fmt.Println(text)
		data := strings.Split(text, ",")
		fmt.Println(data)
		for index := range data {
			instr := data[index]
			if instr[0] == 115 {
				// spin
				count, err := strconv.Atoi(instr[1:])
				if err != nil {
					fmt.Println("could not conver sping", instr)
				}
				programs = spin(programs, count)
			} else if instr[0] == 120 {
				// x
				d2 := strings.Split(instr[1:], "/")
				i1, err := strconv.Atoi(d2[0])
				if err != nil {
					fmt.Println("err conv 1", instr)
				}
				i2, err := strconv.Atoi(d2[1])
				if err != nil {
					fmt.Println("err conv 2", instr)
				}
				programs = exchange(programs, i1, i2)
			} else if instr[0] == 112 {
				// p
				d2 := strings.Split(instr[1:], "/")
				//fmt.Println("P,d2", d2)
				programs = partner(programs, d2[0], d2[1])
			} else {
				fmt.Println("invalid instr", instr)
			}
		}
	}
	fmt.Println("programs", programs)
}

func main() {
	//programs := []string{"a", "b", "c", "d", "e"}
	//part1("test.input", programs)

	programs := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	fmt.Println("total programs", len(programs))
	part1("part1.input", programs)
}
