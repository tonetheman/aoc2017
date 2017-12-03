package main

import "fmt"

const (
	east  = 0
	north = 1
	west  = 2
	south = 3
)

// Point holds x and y values
type Point struct {
	x, y int
}

// Step holds a point and orientation
type Step struct {
	x, y        int
	orientation int
}

var occupied map[Point]bool

func calcLeftFacing(facing int) int {
	if facing == east {
		return north
	}
	if facing == north {
		return west
	}
	if facing == west {
		return south
	}
	return east
}

func calcPointToCheck(x, y int, facing int) Point {
	if facing == north {
		return Point{x, y - 1}
	}
	if facing == east {
		return Point{x + 1, y}
	}
	if facing == west {
		return Point{x - 1, y}
	}
	// east
	return Point{x, y + 1}
}

func newStepGoingForward(old Step) Step {
	if old.orientation == north {
		return Step{old.x, old.y - 1, north}
	}
	if old.orientation == south {
		return Step{old.x, old.y + 1, south}
	}
	if old.orientation == east {
		return Step{old.x + 1, old.y, east}
	}
	//west
	return Step{old.x - 1, old.y, west}
}

func calcSpiral(n int, a *[25]Step) bool {
	if n == 0 {
		a[n] = Step{0, 0, east}
		occupied[Point{0, 0}] = true
		return true
	}
	if n == 1 {
		a[n] = Step{1, 0, east}
		occupied[Point{1, 0}] = true
		return false
	}
	// to calc n look at n-1
	// look left if space is free
	// turn left and move forward
	dirToCheck := calcLeftFacing(a[n-1].orientation)
	pointToCheck := calcPointToCheck(a[n-1].x,
		a[n-1].y, dirToCheck)
	fmt.Println("dirToCheck", dirToCheck)
	fmt.Println("pointToCheck", pointToCheck)
	_, err := occupied[pointToCheck]
	if err == false {
		// the key is not in the map
		a[n] = Step{pointToCheck.x, pointToCheck.y,
			dirToCheck}
		occupied[pointToCheck] = true
	} else {
		// the key is in the map
		// walk forward
		a[n] = newStepGoingForward(a[n-1])
		occupied[Point{a[n].x, a[n].y}] = true
	}
	fmt.Println("--------------")
	return false
}

func main() {
	occupied = make(map[Point]bool)
	var steps [25]Step
	calcSpiral(0, &steps)
	calcSpiral(1, &steps)
	calcSpiral(2, &steps)
	calcSpiral(3, &steps)
	fmt.Println("occupied", occupied)
	fmt.Println("steps", steps)
}
