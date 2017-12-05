package main

import "fmt"

// allow for 1000 x 1000
type Grid [1000 * 1000]int

// pretend it is 500x500
// offset by 250
func (g *Grid) get(row, col int) int {
	return g[(row+250)*500+(250+col)]
}
func (g *Grid) set(row, col int, val int) {
	fmt.Println("grid set", row, col, val)
	g[(250+row)*500+(250+col)] = val
}

func figure_sum(g *Grid, row, col int) int {
	return g.get(row-1, col) + g.get(row+1, col) +
		g.get(row, col+1) + g.get(row, col-1)
}

const (
	east  = 1
	north = 2
	west  = 3
	south = 4
)

func compute_sum(g *Grid, row, col int) int {
	fmt.Println("compute sum", row, col)
	return g.get(row-1, col) + g.get(row+1, col) +
		g.get(row, col-1) + g.get(row, col+1) +
		g.get(row-1, col-1) + g.get(row-1, col+1) +
		g.get(row+1, col-1) + g.get(row+1, col+1)
}

func main() {
	var grid Grid
	var row, col int = 0, 0
	var current_dir int = east
	grid.set(row, col, 1)
	// still facing east
	col++
	grid.set(row, col, 1)

	count := 1
	for {
		if current_dir == east {
			// facing east look left and see if empty
			new_row := row - 1
			new_col := col
			if grid.get(new_row, new_col) == 0 {
				// we can move so move!
				new_sum := compute_sum(&grid, new_row, new_col)
				grid.set(new_row, new_col, new_sum)
				current_dir = north
			} else {
				// we could not move go forward
			}
		}
		if current_dir == north {

		}
		count++
		if count > 2 {
			break
		}
	}
}
