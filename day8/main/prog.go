package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var regs map[string]int

func update(reg string, instr string, inc_amt int) {
	cval := regs[reg]
	if instr == "inc" {
		cval += inc_amt
		regs[reg] = cval
	} else if instr == "dec" {
		cval -= inc_amt
		regs[reg] = cval
	} else {
		fmt.Println("UNKNOWN instr", instr)
	}
}

func readInput(filename string) {
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
		//fmt.Println(text)
		// split by space first
		tokens := strings.Split(text, " ")
		reg := tokens[0]
		instr := tokens[1]
		inc_amt, _ := strconv.Atoi(tokens[2])
		// if is at pos 3
		checkedReg := tokens[4]
		op := tokens[5]
		operand := tokens[6]
		iOperand, _ := strconv.Atoi(operand)
		//fmt.Println(reg, instr, inc_amt, checkedReg, op, operand)

		checkRegVal := regs[checkedReg]
		switch op {
		case ">":
			if checkRegVal > iOperand {
				update(reg, instr, inc_amt)
			}
		case "<":
			if checkRegVal < iOperand {
				update(reg, instr, inc_amt)
			}
		case ">=":
			if checkRegVal >= iOperand {
				update(reg, instr, inc_amt)
			}
		case "<=":
			if checkRegVal <= iOperand {
				update(reg, instr, inc_amt)
			}

		case "==":
			if checkRegVal == iOperand {
				update(reg, instr, inc_amt)
			}
		case "!=":
			if checkRegVal != iOperand {
				update(reg, instr, inc_amt)
			}
		default:
			fmt.Println("UNKNOWN OP", op)
		}
	}
}

func main() {
	regs = make(map[string]int)
	//readInput("main\\test.input")
	//fmt.Println("final regs", regs)
	readInput("main\\part1.input")
	maxValue := -1
	for _, v := range regs {
		if v > maxValue {
			maxValue = v
		}
	}
	fmt.Println("max value for regs:", maxValue)
}
