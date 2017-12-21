package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}
type particle struct {
	p           point
	v           point
	a           point
	lowestCount int
}

type manyParticles []particle

func (a manyParticles) Len() int {
	return len(a)
}
func (a manyParticles) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a manyParticles) Less(i, j int) bool {
	return a[i].lowestCount < a[j].lowestCount
}

// parses a single line from the input
// into a particle struct
func parseString(s string) particle {
	re := regexp.MustCompile("=<.*?>")
	res := re.FindAllString(s, -1)
	var tmp particle
	for idx := range res {
		ts := res[idx][2:]
		tsS := strings.Split(ts, ",")
		//fmt.Println(tsS)
		var err error
		if idx == 0 {
			tmp.p.x, err = strconv.Atoi(strings.Trim(tsS[0], " "))
			if err != nil {
				fmt.Println("ERR1:", err)
			}
			tmp.p.y, err = strconv.Atoi(tsS[1])
			if err != nil {
				fmt.Println("ERR2:", err)
			}
			tmp.p.z, err = strconv.Atoi(strings.TrimRight(tsS[2], ">"))
			if err != nil {
				fmt.Println("ERR3:", err)
			}
		} else if idx == 1 {
			tmp.v.x, err = strconv.Atoi(strings.Trim(tsS[0], " "))
			if err != nil {
				fmt.Println("ERR4:", err)
			}
			tmp.v.y, err = strconv.Atoi(tsS[1])
			if err != nil {
				fmt.Println("ERR5:", err)
			}
			tmp.v.z, err = strconv.Atoi(strings.TrimRight(tsS[2], ">"))
			if err != nil {
				fmt.Println("ERR6:", err)
			}
		} else if idx == 2 {
			tmp.a.x, err = strconv.Atoi(strings.Trim(tsS[0], " "))
			if err != nil {
				fmt.Println("ERR7:", err)
			}
			tmp.a.y, err = strconv.Atoi(tsS[1])
			if err != nil {
				fmt.Println("ERR8:", err)
			}
			tmp.a.z, err = strconv.Atoi(strings.TrimRight(tsS[2], ">"))
			if err != nil {
				fmt.Println("ERR9:", err)
			}
		}
	}
	//fmt.Println("tmp is set", tmp)

	//fmt.Println(re)
	return tmp
}

func (p *particle) tick() {
	p.v.x += p.a.x
	p.v.y += p.a.y
	p.v.z += p.a.z

	p.p.x += p.a.x
	p.p.y += p.a.y
	p.p.z += p.a.z
}

func (p particle) dist() float64 {
	return math.Abs(float64(p.p.x)) + math.Abs(float64(p.p.y)) + math.Abs(float64(p.p.z))
}

func printDist(p []particle) {
	for idx := range p {
		fmt.Printf("%f ", p[idx].dist())
	}
	fmt.Println()
}

func markClosest(p []particle) {
	var maxValue float64 = 1000000
	maxIndex := -1
	for idx := range p {
		d := p[idx].dist()
		if float64(d) < maxValue {
			maxValue = d
			maxIndex = idx
		}
	}
	// found closest to zero mark it
	p[maxIndex].lowestCount++
}

func printParticles(p []particle) {
	sort.Sort(manyParticles(p))
	for idx := range p {
		if p[idx].lowestCount > 0 {
			fmt.Println(idx, p[idx])
		}
	}
}
func testme() {
	steps := []string{"p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>",
		"p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>"}
	p := make([]particle, 0)
	p = append(p, parseString(steps[0]))
	p = append(p, parseString(steps[1]))
	for i := 0; i < 10; i++ {
		for idx := range p {
			p[idx].tick()
		}
		markClosest(p)
	}
	printParticles(p)
}

func readInputfile(filename string) []particle {
	inf, err := os.Open(filename)
	if err != nil {
		fmt.Println("cannot open file", err)
	}
	// close it later
	defer inf.Close()

	scanner := bufio.NewScanner(inf)
	particles := make([]particle, 0)
	for scanner.Scan() {
		// split the text line on the colon
		text := scanner.Text()
		//fmt.Println(text)
		particles = append(particles, parseString(text))
	}
	return particles
}

func findLowestAcc(p []particle) int {
	var maxV float64 = 1000000
	maxIndex := -1
	for idx := range p {
		pp := p[idx]
		v := math.Abs(float64(pp.a.x)) + math.Abs(float64(pp.a.y)) + math.Abs(float64(pp.a.z))
		if v < maxV {
			maxV = v
			maxIndex = idx
		}
	}
	return maxIndex
}

func part1() {
	p := readInputfile("part1.input")
	fmt.Println("read in", len(p), "particles...")

	for i := 0; i < 10000; i++ {
		for idx := range p {
			p[idx].tick()
		}
		markClosest(p)
	}
	printParticles(p)

	// this is a guess
	// acceleration is the driver in infinite time
	// was looking for the particle with the smallest acceleration
	fmt.Println(findLowestAcc(p))
}

func main() {
	//parseString("p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>")
	//fmt.Println("-----------")
	//parseString("p=<-36,-1100,4900>, v=<62,22,56>, a=<-4,1,-15>")
	//testme()
	part1()
}
