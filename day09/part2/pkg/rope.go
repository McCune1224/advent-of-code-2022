package grid

import "fmt"

type knot struct {
	X      int
	Y      int
	Symbol string
}

type Rope struct {
	Knots []knot
}

// New Rope should take in a specified amount of Knots and return a Rope
func NewRope(knotQuantity int) Rope {
	rope := Rope{}
	for i := 0; i < knotQuantity; i++ {
		if i == 0 {
			rope.Knots = append(rope.Knots, knot{X: 0, Y: 0, Symbol: "H"})
		}
		rope.Knots = append(rope.Knots, knot{X: 0, Y: 0, Symbol: fmt.Sprint(i)})
	}
	return rope
}

// TODO:
// A Knot needs to adjust its X,Y coordinates based off if the knot indexed in front of it is not "touching"
// A knot is considered touching another knot based off one of three conditions:
//  1. The Knot's x,y coordinates are the same (overlapping)
//  2. The Knot's x,y coordinates are adjacent
//  3. The Knot's x,y coordinates are diagonal

func (r *Rope) adjustKnots(tail point, head point) {
	for i, knot := range r.Knots {
        if i == 0 {
            continue
        }
        xDif := knot.X - r.Knots[i-1].X
        yDif := knot.Y - r.Knots[i-1].Y
        tempKnot := knot {xDif, yDif, "!"}
        switch tempKnot {
        // case tempKnot.X ==
        }
		// switch (point{head.x - tail.x, head.y - tail.y}) {
		// case point{-2, 1}, point{-1, 2}, point{0, 2}, point{1, 2}, point{2, 1}:
		// 	newTail.y++
		// }
		// switch (point{head.x - tail.x, head.y - tail.y}) {
		// case point{1, 2}, point{2, 1}, point{2, 0}, point{2, -1}, point{1, -2}:
		// 	newTail.x++
		// }
		// switch (point{head.x - tail.x, head.y - tail.y}) {
		// case point{2, -1}, point{1, -2}, point{0, -2}, point{-1, -2}, point{-2, -1}:
		// 	newTail.y--
		// }
		// switch (point{head.x - tail.x, head.y - tail.y}) {
		// case point{-1, -2}, point{-2, -1}, point{-2, -0}, point{-2, 1}, point{-1, 2}:
		// 	newTail.x--
		// }
	}
}
