package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}
type particle struct {
	p point
	v point
	a point
}

// parses a single line from the input
// into a particle struct
func parseString(s string) particle {
	re := regexp.MustCompile("=<.*?>")
	res := re.FindAllString(s, -1)
	//fmt.Println(len(res))
	var tmp particle
	//var tmp particle
	for idx := range res {
		//fmt.Println(res[idx])
		ts := res[idx][2:]
		//fmt.Println(ts)
		tsS := strings.Split(ts, ",")
		//fmt.Println(tsS)
		if idx == 0 {
			tmp.p.x, _ = strconv.Atoi(strings.Trim(tsS[0], " "))
			tmp.p.y, _ = strconv.Atoi(tsS[1])
			tmp.p.z, _ = strconv.Atoi(strings.TrimRight(tsS[2], ">"))
		} else if idx == 1 {
			tmp.v.x, _ = strconv.Atoi(strings.Trim(tsS[0], " "))
			tmp.v.y, _ = strconv.Atoi(tsS[1])
			tmp.v.z, _ = strconv.Atoi(strings.TrimRight(tsS[2], ">"))
		} else if idx == 2 {
			tmp.a.x, _ = strconv.Atoi(strings.Trim(tsS[0], " "))
			tmp.a.y, _ = strconv.Atoi(tsS[1])
			tmp.a.z, _ = strconv.Atoi(strings.TrimRight(tsS[2], ">"))
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

func testme() {
	steps := []string{"p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>",
		"p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>"}

}
func main() {
	parseString("p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>")
	fmt.Println("-----------")
	parseString("p=<-36,-1100,4900>, v=<62,22,56>, a=<-4,1,-15>")
}
