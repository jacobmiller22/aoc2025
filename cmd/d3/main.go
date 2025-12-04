package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	p1Sum := 0
	p2Sum := int64(0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		bank := strings.Trim(scanner.Text(), " \n")

		p1Sum += largestTwoInBank(bank)
		p2Sum += largestTwelveInBank(bank)
	}

	log.Printf("Part 1: Sum of largest 2 batteries in all banks: %d\n", p1Sum)
	log.Printf("Part 2: Sum of largest 12 batteries in all banks : %d\n", p2Sum)

	return nil
}

func p2(inputpath string) error {
	return nil
}

func largestTwoInBank(bank string) int {

	runes := []rune(bank)

	m := 0
	for l := 0; l < len(runes)-1; l++ {
		for r := l + 1; r < len(runes); r++ {
			val, err := strconv.Atoi(string(runes[l]) + string(runes[r]))
			if err != nil {
				panic("invalid string")
			}
			m = int(math.Max(float64(m), float64(val)))
		}
	}
	return m
}

func largestTwelveInBank(bank string) int64 {
	l := 0
	r := ""

	for remaining := 12; remaining > 0; remaining-- {

		mc := byte('0')
		mi := -1

		for i := l; i <= len(bank)-remaining; i++ {

			if bank[i] > mc {
				mc = bank[i]
				mi = i
			}
		}
		r += string(mc)
		l = mi + 1
	}

	val, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		panic("invalid string")
	}
	return val
}
