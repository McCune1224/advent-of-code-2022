package grid

import "fmt"

type Grid struct {
	grid [][]string
}

// Create New Grid filled with "."
func NewGrid(size int) Grid {
	grid := make([][]string, size)
	for i := range grid {
		grid[i] = make([]string, size)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	return Grid{grid}
}

// Method to Print Grid with 0,0 in the bottom left corner
func (g *Grid) PrintGrid() {
	for i := len(g.grid) - 1; i >= 0; i-- {
		fmt.Println(g.grid[i])
	}
}

func (g *Grid) Set(x, y int, value string) {
	g.grid[x][y] = value
}
