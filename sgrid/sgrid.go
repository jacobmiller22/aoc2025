// This package contains simple grid functionality
package sgrid

import (
	"slices"
)

type Coordinate[T any] struct {
	X          int
	Y          int
	Occupation *T
}

func (c Coordinate[T]) Occupied() bool {
	return c.Occupation != nil
}

type Grid2D[T any] struct {
	coordsByX map[int][]*Coordinate[T]
	coordsByY map[int][]*Coordinate[T]
	Width     int
	Height    int
}

func NewGrid2D[T any](width, height int, coords []Coordinate[T]) Grid2D[T] {

	coordsByX := make(map[int][]*Coordinate[T], width)
	coordsByY := make(map[int][]*Coordinate[T], height)

	for _, c := range coords {
		coordsByX[c.X] = append(coordsByX[c.X], &c)
		coordsByY[c.Y] = append(coordsByY[c.Y], &c)
	}

	for _, coords := range coordsByX {
		slices.SortFunc(coords, func(a, b *Coordinate[T]) int {
			return a.Y - b.Y
		})
	}

	return Grid2D[T]{
		coordsByX: coordsByX,
		coordsByY: coordsByY,
		Width:     width,
		Height:    height,
	}
}

func NewGrid2DFromBytes(data []byte, occupiedChars map[byte]struct{}) Grid2D[byte] {

	coords := make([]Coordinate[byte], 0)

	width := 0
	y := 0
	for i, b := range data {

		x := (i - y) - width*y

		if b == '\n' {
			width = x
			y++
			continue
		}

		var occupation *byte
		if _, ok := occupiedChars[b]; ok {
			occupation = &b
		}

		coords = append(coords, Coordinate[byte]{
			X:          x,
			Y:          y,
			Occupation: occupation,
		})
	}

	height := y + 1

	return NewGrid2D(width, height, coords)
}

func (g Grid2D[T]) Coordinate(x, y int) *Coordinate[T] {
	if x < 0 || x >= g.Width {
		return nil
	}
	if y < 0 || y >= g.Height {
		return nil
	}
	if coords, ok := g.coordsByX[x]; ok {
		if r, found := slices.BinarySearchFunc(
			coords,
			y,
			func(c *Coordinate[T], t int) int {
				return c.Y - t
			}); found {
			return coords[r]
		}
	}
	return nil
}

func (g Grid2D[T]) StrictAdjacent(x, y int) []*Coordinate[T] {
	adjacentCoords := make([]*Coordinate[T], 0)
	coordsToCheck := [][2]int{
		{x - 1, y},
		{x, y - 1},
		{x + 1, y},
		{x, y + 1},
	}

	for _, pair := range coordsToCheck {
		if c := g.Coordinate(pair[0], pair[1]); c != nil {
			adjacentCoords = append(adjacentCoords, c)
		}
	}
	return adjacentCoords
}

func (g Grid2D[T]) Adjacent(x, y int) []*Coordinate[T] {

	adjacentCoords := make([]*Coordinate[T], 0)

	adjacentCoords = append(adjacentCoords, g.StrictAdjacent(x, y)...)

	coordsToCheck := [][2]int{
		{x - 1, y - 1},
		{x - 1, y + 1},
		{x + 1, y - 1},
		{x + 1, y + 1},
	}

	for _, pair := range coordsToCheck {
		if c := g.Coordinate(pair[0], pair[1]); c != nil {
			adjacentCoords = append(adjacentCoords, c)
		}
	}
	return adjacentCoords
}

func (g Grid2D[T]) Coords() []*Coordinate[T] {

	coords := make([]*Coordinate[T], 0)
	for _, cs := range g.coordsByX {
		coords = append(coords, cs...)
	}
	return coords
}

func (g Grid2D[T]) CoordsAtX(x int) []*Coordinate[T] {
	if _, ok := g.coordsByX[x]; ok {
		return g.coordsByX[x]
	}
	return nil
}
func (g Grid2D[T]) CoordsAtY(y int) []*Coordinate[T] {
	if _, ok := g.coordsByY[y]; ok {
		return g.coordsByY[y]
	}
	return nil
}

func GridToBytes(g Grid2D[byte]) []byte {

	stride := g.Width + 1
	data := make([]byte, stride*g.Height)

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			// Calculate the specific index in the flat slice
			// (Current Row * Row Length) + Current Column
			index := (y * stride) + x

			if c := g.Coordinate(x, y); c != nil && c.Occupied() {
				data[index] = *c.Occupation
			} else {
				data[index] = '.'
			}
		}

		// 2. Insert the newline at the end of the row
		newlineIndex := (y * stride) + g.Width
		data[newlineIndex] = '\n'
	}

	return data
}
