package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jacobmiller22/aoc2025/pconfig"
)

type Rotation string

const RotationLeft Rotation = "L"
const RotationRight Rotation = "R"

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
		return fmt.Errorf("error opening input file: %+v", err)
	}

	scanner := bufio.NewScanner(f)

	ringsize := 100
	clicks := 50 // dial starts at 50

	zerocounts := 0

	for scanner.Scan() {
		direction, distance, err := parseline(scanner.Text())
		if err != nil {
			return err
		}
		prev := clicks

		clicks = rotateStrict(ringsize, prev, direction, distance)

		log.Printf("Rotated %s %d clicks from %d to %d\n", direction, distance, prev, clicks)

		if clicks == 0 {
			zerocounts++
		}
	}

	log.Printf("Landed on zero %d times\n", zerocounts)
	return nil
}

func p2(inputpath string) error {
	f, err := os.Open(inputpath)
	if err != nil {
		return fmt.Errorf("error opening input file: %+v", err)
	}

	scanner := bufio.NewScanner(f)

	ringsize := 100
	prev := 50 // dial starts at 50

	zerocounts := 0

	for scanner.Scan() {
		direction, distance, err := parseline(scanner.Text())
		if err != nil {
			return err
		}

		new, zeroclicks := rotate(ringsize, prev, direction, distance)

		log.Printf("Rotated %s%d to point at %d; clicks %d\n", direction, distance, new, zeroclicks)

		zerocounts += zeroclicks

		prev = new

	}

	log.Printf("Landed on zero %d times\n", zerocounts)
	return nil
}

func parseline(line string) (Rotation, int, error) {
	line = strings.Trim(line, " \n")

	var rotation Rotation
	if strings.HasPrefix(line, "R") {
		rotation = RotationRight
	} else if strings.HasPrefix(line, "L") {
		rotation = RotationLeft
	} else {
		return "", 0, fmt.Errorf("Invalid rotation in line: %s\n", line)
	}

	distancestr := line[1:]
	distance, err := strconv.Atoi(distancestr)
	if err != nil {
		return rotation, 0, fmt.Errorf("Invalid distance: %s", distancestr)
	}

	return rotation, distance, nil

}

func rotate(ringsize, curr int, direction Rotation, distance int) (newpos int, zeroclicks int) {
	switch direction {
	case RotationLeft:
		if distance >= curr {
			zeroclicks = (ringsize - curr + distance) / ringsize
			// edge case for if we start at 0 and go left, shouldn't count the first click
			if curr == 0 && zeroclicks > 0 {
				zeroclicks--
			}
		}
		newpos = curr - distance
	case RotationRight:
		if distance >= ringsize-curr {
			zeroclicks = (curr + distance) / ringsize
		}
		newpos = curr + distance
	default:
		panic("rotate: received invalid Rotation")
	}

	newpos = (newpos + (ringsize * (zeroclicks + 1))) % ringsize

	return newpos, zeroclicks
}

func rotateStrict(ringsize, curr int, direction Rotation, distance int) int {
	new := float64(curr)
	var res int
	switch direction {
	case RotationLeft:
		res = (int(new - float64(distance)))
	case RotationRight:
		res = (int(new + float64(distance)))
	default:
		panic("rotate: received invalid Rotation")
	}

	res = (res + ringsize) % ringsize

	return res
}
