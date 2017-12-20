package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type point struct {
	x, y, z int
}
type particle struct {
	p point
	v point
	a point
}

func parseString(s string) particle {
	re := regexp.MustCompile("=<.*?>")
	re2 := regexp.MustCompile("([0-9]+),([0-9]+),([0-9]+)")
	res := re.FindAllString(s, -1)
	fmt.Println(len(res))
	var tmp particle
	//var tmp particle
	for idx := range res {
		fmt.Println(res[idx])
		ts := re2.FindAllString(res[idx], -1)
		re2.

		if idx == 0 {
			tmp.p.x, _ = strconv.Atoi(ts[0])
			tmp.p.y, _ = strconv.Atoi(ts[1])
			tmp.p.z, _ = strconv.Atoi(ts[2])
		}
	}
	//fmt.Println(re)
	return tmp
}

func main() {
	parseString("p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>")
}
