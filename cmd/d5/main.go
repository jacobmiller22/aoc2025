package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/Workiva/go-datastructures/augmentedtree"
	"github.com/jacobmiller22/aoc2025/bounds"
	"github.com/jacobmiller22/aoc2025/pconfig"
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

type IngredientRange struct {
	Id    uint64
	Start int64
	End   int64
}

func (r IngredientRange) LowAtDimension(dim uint64) int64 {
	return r.Start
}

func (r IngredientRange) HighAtDimension(dim uint64) int64 {
	return r.End
}

func (r IngredientRange) OverlapsAtDimension(i augmentedtree.Interval, dim uint64) bool {
	return r.HighAtDimension(dim) >= r.LowAtDimension(dim) && r.LowAtDimension(dim) >= i.HighAtDimension(dim)
}

func (r IngredientRange) ID() uint64 {
	return r.Id
}

func p1(inputpath string) error {
	f, err := os.Open(inputpath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	scanner := bufio.NewScanner(f)

	fridge := false

	sum1 := 0 // count of fresh ingredients
	var sum2 int64 = 0

	tree := augmentedtree.New(1)

	ranges := make([]augmentedtree.Interval, 0)

	intervals := make([]IngredientRange, 0)

	var ln uint64 = 0
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			tree.Add(ranges...)
			fridge = true
			continue
		}

		if fridge == false {
			// Parse as a range

			l, u, err := bounds.Bounds(s)
			if err != nil {
				return err
			}
			ranges = append(ranges, IngredientRange{
				Id:    ln,
				Start: int64(l),
				End:   int64(u),
			})
			intervals = append(intervals, IngredientRange{
				Id:    ln,
				Start: int64(l),
				End:   int64(u),
			})

			ln++
			continue
		}

		// Use as ingredient
		id, err := strconv.Atoi(s)
		if err != nil {
			return err
		}

		matchedIntervals := tree.Query(IngredientRange{
			Start: int64(id),
			End:   int64(id),
		})

		if len(matchedIntervals) > 0 {
			sum1++
		}
	}

	slices.SortFunc(intervals, func(a, b IngredientRange) int {
		return int(a.Start) - int(b.Start)
	})

	aggregatedIntervals := make([]IngredientRange, 0)

	aggregatedIntervals = append(aggregatedIntervals, intervals[0])

	for _, interval := range intervals[1:] {
		if len(aggregatedIntervals) <= 0 || aggregatedIntervals[len(aggregatedIntervals)-1].End < interval.Start {
			aggregatedIntervals = append(aggregatedIntervals, interval)
		} else {
			// Check if our latest aggregate completed covers this current range
			if aggregatedIntervals[len(aggregatedIntervals)-1].End < interval.End {
				// Combine
				aggregatedIntervals[len(aggregatedIntervals)-1].End = interval.End
			}
		}
	}

	for _, interval := range aggregatedIntervals {
		sum2 += interval.End - interval.Start + 1
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	log.Printf("Part 1: Sum of invalid ids: %d\n", sum1)
	log.Printf("Part 2: Sum of invalid ids: %d\n", sum2)

	return nil
}

func p2(inputpath string) error {
	return nil
}
