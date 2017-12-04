package main

import "fmt"

const (
	east  = 1
	north = 2
	west  = 3
	south = 4
)

type Position struct {
	x, y int
}

type Orientation int

type Step struct {
	pos Position
	ori Orientation
}

var occu map[Position]bool
var a [100]Step

func takeStep(n int) {
	if n == 0 {
		a[n] = Step{Position{0, 0}, east}
		occu[Position{0, 0}] = true
		return
	}
	if n == 1 {
		a[n] = Step{Position{1, 0}, east}
		occu[Position{1, 0}] = true
		return
	}
	var last = a[n-1]
	if last.ori == east {
		// look NORTH is it empty?
		lookPos := Position{last.pos.x, last.pos.y - 1}
		_, err := occu[lookPos]
		if err == false {
			// it is empty!
			// move that direction
			fmt.Println("can move to the left!")
			newPos := Position{last.pos.x, last.pos.y - 1}
			occu[newPos] = true
			a[n] = Step{newPos, north}
		} else {
			// to our left is not empty
			// move forward
		}
	}
	// now look
}

func main() {
	occu = make(map[Position]bool)
	takeStep(0)
	takeStep(1)
	takeStep(2)
	fmt.Println("occu", occu)
}
