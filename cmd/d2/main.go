package main

import (
	"log"

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
	return nil
}

func p2(inputpath string) error {
	return nil
}
