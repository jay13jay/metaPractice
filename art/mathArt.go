package art

import "fmt"

type Drawing struct {
	Strokes int32
	Vectors []int32
	Directions string
	PlusSigns int64
	X int32
	Y int32
	xCheck map[int32]bool
	yCheck map[int32]bool
	xLines map[int32][2]int32
	yLines map[int32][2]int32
}

func newDrawing(N int32, L []int32, D string) *Drawing {
	return &Drawing{
		PlusSigns: 0,
		Strokes: N, 
		Vectors: L, 
		Directions: D,
		X: 0,
		Y: 0,
		xCheck: make(map[int32]bool),
		yCheck: make(map[int32]bool),
		// map of vertical lines - key is X coord, value maps to a low Y and high Y
		// ex. 2: [1,3] maps to [2,1] and [2,3] forming 
		// a verical line at X coord 2, from 1 - 3
		// xLines[x][0] = x-lowY 
		// xLines[x][1] = x-highY
		xLines: make(map[int32][2]int32),
		// map of horizontal lines - key is Y coord, value maps to a low X and high X
		yLines: make(map[int32][2]int32),
	}

}

func getPlusSignCount(N int32, L []int32, D string) int64 {
  // Write your code here
	d := newDrawing(N, L, D)
	d.yCheck[0] = true
	d.xCheck[0] = true
	d.xLines[0] = [2]int32{0, 0}
	d.yLines[0] = [2]int32{0, 0}

	for i := range d.Vectors {
		d.setNewCoord(i)
	}
	fmt.Printf("X: %v\nY: %v\n", d.xLines, d.yLines)
	for i, v := range d.xLines {
		fmt.Printf("Key: %d\tVal: %v\n", i, v)
	}	
	return d.PlusSigns
}

func (d *Drawing) updateMaps() {
	if ! d.xCheck[d.X] {
		d.xCheck[d.X] = true
		d.xLines[d.X] = [2]int32{d.Y, d.Y}
	} else {
		line := d.xLines[d.X]
		if d.Y < d.xLines[d.X][0] {
			line[0] = d.Y
			d.xLines[d.X] = line
		} else if d.Y > d.xLines[d.X][1] {
			line[1] = d.Y
			d.xLines[d.X] = line
		}
	}
	
	if ! d.yCheck[d.Y] {
		d.yCheck[d.Y] = true
		d.yLines[d.Y] = [2]int32{d.X, d.X}
	} else {
		line := d.yLines[d.Y]
		if d.X < d.yLines[d.Y][0] {
			line[0] = d.X
			d.yLines[d.Y] = line
		} else if d.X > d.yLines[d.Y][1] {
			line[1] = d.X
			d.yLines[d.Y] = line
		}
	}
}

func (d *Drawing) setNewCoord(index int) {
	dir := string(d.Directions[index])
	vector := d.Vectors[index]
	if dir == "U" {
		// fmt.Printf("Moving up | current Y: %d\tNew Y: %d\n", d.Y, d.Y+vector)
		d.Y += vector
		d.updateMaps()
	} else if dir == "L" {
		// fmt.Printf("Moving left | current X: %d\tNew X: %d\n", d.X, d.X+vector)
		d.X -= vector
		d.updateMaps()
	} else if dir == "R" {
		// fmt.Printf("Moving right | current X: %d\tNew X: %d\n", d.X, d.X+vector)
		d.X += vector
		d.updateMaps()
	} else if dir == "D" {
		// fmt.Printf("Moving down | current Y: %d\tNew Y: %d\n", d.Y, d.Y+vector)
		d.Y -= vector
		d.updateMaps()
	} else {
		fmt.Println("Invalid direction")
	}

}

func Solve(N int32, L []int32, D string) int64 {
	return getPlusSignCount(N, L, D)
}