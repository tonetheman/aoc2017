package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Triple struct {
	x, y, z int
}

type InfHex struct {
	currentPos Triple
}

func directionToTriple(s string) Triple {
	switch s {
	case "nw":
		return Triple{-1, +1, 0}
	case "n":
		return Triple{0, 1, -1}
	case "ne":
		return Triple{1, 0, -1}
	case "s":
		return Triple{0, -1, 1}
	case "se":
		return Triple{1, -1, 0}
	case "sw":
		return Triple{-1, 0, 1}
	default:
		fmt.Println("ERRRRR")
		return Triple{-100, -100, -100}
	}
}

func dist(a, b Triple) int {
	tmp1 := math.Abs(float64(a.x - b.x))
	tmp2 := math.Abs(float64(a.y - b.y))
	tmp3 := math.Abs(float64(a.z - b.z))
	return int(math.Max(math.Max(tmp1, tmp2), tmp3))
}

func (i *InfHex) walk(s string) {
	steps := strings.Split(s, ",")
	for _, step := range steps {
		cd := directionToTriple(step)
		fmt.Println(cd)
		i.currentPos.x += cd.x
		i.currentPos.y += cd.y
		i.currentPos.z += cd.z
	}
}

func check(s string) {
	var Inf InfHex
	fmt.Println("before walk", Inf)
	Inf.walk(s)
	fmt.Println("after walk", Inf)
	fmt.Println("dist", dist(Triple{0, 0, 0}, Inf.currentPos))
}

func checks_part1() {
	check("ne,ne,ne")
	check("ne,ne,sw,sw")
	check("ne,ne,s,s")
	check("se,sw,se,sw,sw")
}

func main() {
	//checks_part1()
	data, err := ioutil.ReadFile("part1.input")
	if err != nil {
		fmt.Println("err", err)
	}
	check(string(data))
}
