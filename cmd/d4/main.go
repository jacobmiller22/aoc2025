package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jacobmiller22/aoc2025/pconfig"
	"github.com/jacobmiller22/aoc2025/sgrid"
)

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
	sum2 := 0

	g := sgrid.NewGrid2DFromBytes(
		data,
		map[byte]struct{}{
			'@': {},
		},
	)
	// visual grid for viewing
	vg := sgrid.NewGrid2DFromBytes(
		data,
		map[byte]struct{}{
			'@': {},
		},
	)

	nl := byte('x')

	fmt.Printf("%s\n", sgrid.GridToBytes(g))

	makePass := func() []*sgrid.Coordinate[byte] {

		toRemove := make([]*sgrid.Coordinate[byte], 0)
		for _, c := range g.Coords() {
			if !c.Occupied() {
				continue
			}
			adjacent := g.Adjacent(c.X, c.Y)
			nOccupied := 0
			for _, adjc := range adjacent {
				if adjc.Occupied() {
					nOccupied++
				}
			}
			if nOccupied < 4 {
				toRemove = append(toRemove, c)
				// c.Occupation = &nl
				vg.Coordinate(c.X, c.Y).Occupation = &nl
			}
		}
		return toRemove
	}

	for passes := 0; ; passes++ {
		toRemove := makePass()
		if passes == 0 {
			sum1 += len(toRemove)
		}
		for i := range toRemove {
			toRemove[i].Occupation = nil
		}
		sum2 += len(toRemove)

		if len(toRemove) == 0 {
			break
		}
		fmt.Printf("Pass %d:\n%s\n\n", passes, sgrid.GridToBytes(vg))
	}

	log.Printf("Part 1: Sum of invalid ids: %d\n", sum1)
	log.Printf("Part 2: Sum of invalid ids: %d\n", sum2)

	return nil
}

func p2(inputpath string) error {
	return nil
}
