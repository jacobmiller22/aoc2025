package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

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

func p1(inputpath string) error {
	f, err := os.Open(inputpath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	ranges := strings.Split(strings.Trim(string(data), " \n"), ",")
	invalidSum := 0
	invalidSum2 := 0

	for _, r := range ranges {
		lb, rb, err := bounds(r)
		if err != nil {
			return fmt.Errorf("invalid range: %w", err)
		}

		for p := lb; p <= rb; p++ {
			sp := strconv.Itoa(p)
			if !isvalid1(sp) {
				invalidSum += p
			}
			if !isvalid2(sp) {
				invalidSum2 += p
			}
		}
	}

	log.Printf("Part 1: Sum of invalid ids: %d\n", invalidSum)
	log.Printf("Part 2: Sum of invalid ids: %d\n", invalidSum2)

	return nil
}

func isvalid1(s string) bool {

	lhalf, rhalf := halfs(s)

	if len(lhalf) != len(rhalf) {
		return true
		// continue // odd number ids can't repeat unless they are all the same?
	}

	if repeating(s) {
		return false
		// continue // invalid
	}

	if strings.Compare(lhalf, rhalf) == 0 {
		return false
		// invalidSum += p
		// continue
	}
	return true
}

func isvalid2(s string) bool {
	return !repeatingSequence(s)
}

func p2(inputpath string) error {
	return nil
}

func bounds(r string) (int, int, error) {
	bounds := strings.Split(r, "-")
	if len(bounds) != 2 {
		return 0, 0, fmt.Errorf("bounds: multiple hyphens %s", r)
	}

	lbound, err := strconv.Atoi(bounds[0])
	if err != nil {
		return 0, 0, fmt.Errorf("bounds: bad lower bound: %w", err)
	}

	rbound, err := strconv.Atoi(bounds[1])
	if err != nil {
		return 0, 0, fmt.Errorf("bounds: bad upper bound: %w", err)
	}

	return lbound, rbound, nil
}

func halfs(s string) (string, string) {
	return s[0 : len(s)/2], s[len(s)/2:]
}

func repeating(s string) bool {

	curr := s[0]
	for i := 1; i < len(s); i++ {
		if curr != s[i] {
			return false
		}
		curr = s[i]
	}
	return true
}

// repeatingSequence returns the shortest repeating sequence that composes the entire string
func repeatingSequence(s string) bool {

	n := len(s)

	// For each divisor of n
	for j := 1; j <= n/2; j++ {
		if n%j != 0 {
			continue
		}
		curr := s[0:j]
		isMatch := true

		for i := j; i+j <= n; i += j {
			if strings.Compare(curr, s[i:i+j]) != 0 {
				isMatch = false
				break
			}
			curr = s[i : i+j]
		}
		if isMatch {
			return true
		}
	}

	return false
}
