package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jacobmiller22/aoc2025/pconfig"
	"github.com/jacobmiller22/aoc2025/sgrid"
)

var cache map[*sgrid.Coordinate[byte]]int

func init() {
	cache = make(map[*sgrid.Coordinate[byte]]int, 0)
}

type ProblemConfig struct {
	InputPath string
	Part      int
}

func main() {

	pcfg := pconfig.Parse()

	var err error
	switch pcfg.Part {
	case 1:
		err = p1(pcfg.InputPath)
	case 2:
		err = p2(pcfg.InputPath)
	default:
		log.Fatal("Invalid part")
	}

	if err != nil {
		log.Fatalf("error running problem: %+v\n", err)
	}

}

func p1(inputpath string) error {
	f, err := os.Open(inputpath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	sum1 := 0

	g := sgrid.NewGrid2DFromBytes(
		data,
		map[byte]struct{}{
			'S': {},
			'^': {},
		},
	)
	// visual grid for viewing
	vg := sgrid.NewGrid2DFromBytes(
		data,
		map[byte]struct{}{
			'@': {},
			'^': {},
		},
	)

	nl := byte('|')
	schar := byte('S')

	// find the start
	var startCoord *sgrid.Coordinate[byte]
	for _, coord := range g.CoordsAtY(0) {
		if coord.Occupation != nil && *coord.Occupation == byte('S') {
			startCoord = coord
			break
		}
	}

	if startCoord == nil {
		return fmt.Errorf("didn't find the Starting Coordinate on the first row")
	}

	vg.Coordinate(startCoord.X, startCoord.Y).Occupation = &schar

	fmt.Printf("Starting Coordiante (%d, %d)\n", startCoord.X, startCoord.Y)

	visited := make(map[*sgrid.Coordinate[byte]]bool, 0)

	stack := []*sgrid.Coordinate[byte]{startCoord}

	for len(stack) > 0 {

		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, ok := visited[curr]; ok {
			continue
		}

		// Update the vg
		if c := vg.Coordinate(curr.X, curr.Y); c != nil && c.Occupation == nil || (*c.Occupation != byte('S') && *c.Occupation != byte('^')) {
			c.Occupation = &nl
		}

		visited[curr] = true

		// if we are not on a splitter, we go below
		if curr.Occupation == nil || *curr.Occupation != byte('^') {
			if next := g.Coordinate(curr.X, curr.Y+1); next != nil {
				stack = append(stack, next)

			}
		}

		// if we are on a splitter, add the horizontally adjacent nodes that haven't been visited
		if curr.Occupation != nil && *curr.Occupation == byte('^') {
			sum1++
			if next := g.Coordinate(curr.X-1, curr.Y); next != nil {
				if !visited[next] {
					stack = append(stack, next)
				}

			}
			if next := g.Coordinate(curr.X+1, curr.Y); next != nil {
				if !visited[next] {
					stack = append(stack, next)
				}

			}
		}
	}

	fmt.Printf("%s\n", sgrid.GridToBytes(vg))

	log.Printf("Part 1: %d\n", sum1)

	log.Printf("Part 2: %+v\n", bubbleDown(&g, startCoord))

	return nil
}

func p2(inputpath string) error {
	return nil
}

func bubbleDown(g *sgrid.Grid2D[byte], curr *sgrid.Coordinate[byte]) int {
	if curr == nil {
		return 1
	}

	// if we are on a splitter, we add 2 timelines to the cache
	if curr.Occupation != nil && *curr.Occupation == byte('^') {
		leftChild := g.Coordinate(curr.X-1, curr.Y)
		var left int
		if cached, ok := cache[leftChild]; ok {
			left = cached
		} else {
			left = bubbleDown(g, leftChild)
		}
		rightChild := g.Coordinate(curr.X+1, curr.Y)
		var right int
		if cached, ok := cache[rightChild]; ok {
			right = cached
		} else {
			right = bubbleDown(g, rightChild)
		}
		cache[curr] = left + right
		return left + right
	}

	// if we are on nothing, we go down
	r := bubbleDown(g, g.Coordinate(curr.X, curr.Y+1))
	cache[curr] = r
	return r
}
