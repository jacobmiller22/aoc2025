package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	numericRegex := regexp.MustCompile(`^[a-zA-Z0-9]*$`)

	var sum1 uint64 = 0
	sum2 := 0

	scanner := bufio.NewScanner(f)

	operands := make([][]uint64, 0)
	operators := make([]byte, 0)

	i := 0
	for scanner.Scan() {

		elems := strings.Fields(scanner.Text())

		if numericRegex.MatchString(elems[0]) {
			// operand
			row := make([]uint64, 0, len(elems))
			for _, elem := range elems {
				val, err := strconv.Atoi(elem)
				if err != nil {
					return err
				}

				row = append(row, uint64(val))
			}
			operands = append(operands, row)
		} else {
			// operators
			for _, elem := range elems {
				operators = append(operators, []byte(elem)...)
			}
		}
		i++

	}

	if err := validate(operands, operators); err != nil {
		return err
	}

	fmt.Printf("%+v\n%+v\n", operands, operators)

	for i, operator := range operators {
		result := calculateColumn(column(operands, i), operator)

		fmt.Printf("i: %d\n", i)

		sum1 += result
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	log.Printf("Part 1: Sum of invalid ids: %d\n", sum1)
	log.Printf("Part 2: Sum of invalid ids: %d\n", sum2)

	return nil
}

func p2(inputpath string) error {
	return nil
}

func validate(operands [][]uint64, operators []byte) error {
	w := len(operands[0])

	for i, line := range operands {
		if len(line) != w {
			return fmt.Errorf("validate: operand count mismatch on row %d, (want %d, got %d)", i, w, len(line))
		}
	}

	if len(operators) != w {
		return fmt.Errorf("validate: operator count mismatch, (want %d, got %d)", w, len(operators))
	}

	fmt.Printf("Width: %d\n Height: %d\n", w, len(operands))

	return nil
}

func calculateColumn(column []uint64, operator byte) uint64 {

	var fn func(a, b uint64) uint64
	var result uint64
	switch operator {
	case '*':
		fn = func(a, b uint64) uint64 { return a * b }
		result = 1
	case '+':
		fn = func(a, b uint64) uint64 { return a + b }
		result = 0
	default:
		panic(fmt.Sprintf("unknown operator: '%c'", operator))
	}

	for i := range column {
		result = fn(result, column[i])
	}

	return result
}

// column returns the values of column i in the given matrix m
func column(m [][]uint64, i int) []uint64 {

	values := make([]uint64, 0)
	for _, r := range m {
		values = append(values, r[i])
	}
	fmt.Printf("Column: %+v\n", values)
	return values
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
