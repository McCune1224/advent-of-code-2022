package grid

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	grid := NewGrid(3)
	if len(grid.grid) != 3 {
		t.Errorf("Grid should be 3x3, but is %dx%d", len(grid.grid), len(grid.grid[0]))
	}
}

func TestSet(t *testing.T) {
	grid := NewGrid(30)
	symbol := "#"
	x, y := 5, 3
	grid.Set(x, y, symbol)
	if grid.grid[x][y] != symbol {
		t.Errorf("Grid should be %s, but is %s", symbol, grid.grid[0][0])
	}

	grid.PrintGrid()
}
