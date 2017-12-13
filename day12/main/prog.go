package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type connection struct {
	parent   int
	children []int
}

func readFile(filename string) []connection {
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	programs := make([]connection, 0)
	scanner := bufio.NewScanner(inf)
	for scanner.Scan() {
		text := scanner.Text()
		res := strings.Split(text, " <-> ")
		res2 := strings.Split(res[1], ",")
		parent, _ := strconv.Atoi(res[0])
		children := make([]int, 0)
		for _, v := range res2 {
			tmp, _ := strconv.Atoi(strings.Trim(v, " "))
			children = append(children, tmp)
		}
		c := connection{parent, children}
		programs = append(programs, c)
	}
	return programs
}

type lifo struct {
	data []int
}

// WHAT I HAVE IS A LIFO
func (s *lifo) push(val int) {
	s.data = append(s.data, val)
}
func (s lifo) empty() bool {
	return len(s.data) == 0
}
func (s *lifo) pop() (int, error) {
	if s.empty() {
		// in my case -1 is nothing
		// if this was a real stack
		// then return an error
		return -1, errors.New("empty stack")
	}
	var tmp int
	// this pulls the last thing inserted
	// then resets data to the rest of the list
	tmp, s.data = s.data[0], s.data[1:]
	return tmp, nil
}
func (s lifo) len() int {
	return len(s.data)
}

func isAlreadyFound(af map[int]bool, lookingFor int) bool {
	_, present := af[lookingFor]
	// golang sucks for this type of stuff
	// i think true means it is there
	return present
}

// this function finds the group for the parent
// then removes it from the programMap list
func findOneGroup(programMap map[int]connection, parentIndex int) {
	// by prob definition 0 is already found
	var alreadyFound map[int]bool
	alreadyFound = make(map[int]bool)
	// put parentIndex in the already found map
	alreadyFound[parentIndex] = true

	// save off parentIndex into the stack
	var ll lifo
	ll.push(parentIndex)

	count := 0
	for {
		fmt.Println("loop begin")
		if ll.empty() {
			fmt.Println("ll is empty stopping")
			break
		}

		// look in lifo for first person
		currentGuy, _ := ll.pop()
		currentProgram := programMap[currentGuy]
		fmt.Println("currentGuy", currentGuy, currentProgram)

		// find direct children of current Guy and mark them
		// in the group and push them on the ll
		for i := 0; i < len(currentProgram.children); i++ {
			currentChildIndex := currentProgram.children[i]
			fmt.Println("currentChildIndex", currentChildIndex)

			_, present := alreadyFound[currentChildIndex]
			if !present {
				// !present is golang stupid way of
				// checking hashmap
				// they were not already in the group
				// mark them and push onto lifo
				alreadyFound[currentChildIndex] = true
				ll.push(currentChildIndex)
			}
		}
		count++ // used to stop run away bad program
		fmt.Println("EOL: ll", ll)
		fmt.Println("EOL: af", alreadyFound)
		fmt.Println("----------")
		if count == 10000 {
			fmt.Println("BROKE FROM COUNTER")
			break
		}

	} // end of while looop
	fmt.Println("len of af", len(alreadyFound))
	fmt.Println("total programs", len(programMap))

	// go through the list of programs
	// found in this group and remove them
	// from program map
	for k := range alreadyFound {
		delete(programMap, k)
	}
}

func getNextKey(programMap map[int]connection) int {
	var tmp int
	for k := range programMap {
		tmp = k
		break
	}
	return tmp
}

func main() {
	//programs := readFile("test.input")
	programs := readFile("part1.input")

	// parse the input and save into better
	// format
	var programMap map[int]connection
	programMap = make(map[int]connection)
	for _, p := range programs {
		//fmt.Println(p)
		programMap[p.parent] = p
	}
	fmt.Println(programMap)

	var count = 0
	for {
		// get a key from the program map
		// does not matter which really
		nextParent := getNextKey(programMap)
		fmt.Println("next parent", nextParent)

		// find the group and remove it
		// from the program map
		findOneGroup(programMap, nextParent)
		fmt.Println("after we found a group prgramMap len", len(programMap))
		count++

		if len(programMap) == 0 {
			fmt.Println("out of programs!")
			fmt.Println("count is", count)
			break
		}
		if count == 300 {
			break
		}
	}
	// make 0 the parent
	//findOneGroup(programMap, 0)
	//fmt.Println("after we found a group prgramMap len", len(programMap))

}
