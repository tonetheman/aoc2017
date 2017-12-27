package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	i_assign   = 5
	i_exchange = 6
	i_partner  = 7
)

type instruction struct {
	typ    int
	count  int
	v1, v2 int
	a1, a2 string
}

const asize = 16

var programs map[int]string

func pp() {
	p := programs
	for i := 0; i < asize; i++ {
		fmt.Printf(p[i])
	}
	fmt.Println()
}

func shift1() {
	p := programs
	tmp := p[asize-1] // this is the last one
	//fmt.Println("tmp is saved", tmp)
	for i := asize - 1; i > 0; i-- {
		//fmt.Println("\ti=i-1", i, i-1)
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

func partner(a1, a2 string) {
	var f = func(v string) int {
		for key, val := range programs {
			if v == val {
				return key
			}
		}
		fmt.Println("ERR: v", v)
		return -1
	}
	pos1 := f(a1)
	pos2 := f(a2)
	exchange(pos1, pos2)
}

func readStringFromFile(filename string) string {
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
		return text
	}
	return ""
}

func fill() {
	programs[0] = "a"
	programs[1] = "b"
	programs[2] = "c"
	programs[3] = "d"
	programs[4] = "e"
	programs[5] = "f"
	programs[6] = "g"
	programs[7] = "h"
	programs[8] = "i"
	programs[9] = "j"
	programs[10] = "k"
	programs[11] = "l"
	programs[12] = "m"
	programs[13] = "n"
	programs[14] = "o"
	programs[15] = "p"
}

func parseOnce(inputString string) []instruction {
	tmp := make([]instruction, 0)

	data := strings.Split(inputString, ",")
	//fmt.Println(data)
	for index := range data {
		instr := data[index]
		if instr[0] == 115 {
			// spin
			count, err := strconv.Atoi(instr[1:])
			if err != nil {
				fmt.Println("could not conver sping", instr)
			}
			var tmp_i instruction
			tmp_i.typ = i_assign
			tmp_i.count = count
			tmp = append(tmp, tmp_i)
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
			var tmp_i instruction
			tmp_i.typ = i_exchange
			tmp_i.v1 = i1
			tmp_i.v2 = i2
			tmp = append(tmp, tmp_i)
		} else if instr[0] == 112 {
			// p
			d2 := strings.Split(instr[1:], "/")
			var tmp_i instruction
			tmp_i.typ = i_partner
			tmp_i.a1 = d2[0]
			tmp_i.a2 = d2[1]
			tmp = append(tmp, tmp_i)
		} else {
			fmt.Println("invalid instr", instr)
		}
	}
	return tmp
}

func part1a(inputString string, instructions []instruction) {
	for index := range instructions {
		instr := instructions[index]
		if instr.typ == i_assign {
			shift(instr.count)
		} else if instr.typ == i_exchange {
			exchange(instr.v1, instr.v2)
		} else if instr.typ == i_partner {
			partner(instr.a1, instr.a2)
		}
	}
}

func main() {
	programs = make(map[int]string)
	fill()

	dataInput := readStringFromFile("part1.input")
	instructions := parseOnce(dataInput)
	var i int64
	for i = 0; i < 1000000; i++ {
		part1a(dataInput, instructions)
		if i%1000 == 0 {
			fmt.Printf("*")
		}
	}
	fmt.Println(programs)

	//	pp()
	//	shift(15)
	//exchange(0, 4)
	//shift1()
	//partner("a", "c")
	//	pp()
}
