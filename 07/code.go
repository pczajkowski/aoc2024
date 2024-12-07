package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	result  int
	numbers []int
}

func readInput(file *os.File) []Equation {
	scanner := bufio.NewScanner(file)
	var equations []Equation

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var equation Equation
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			log.Fatalf("Bad input: %s", line)
		}

		var err error
		equation.result, err = strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Problem parsing '%s': %s", parts[0], err)
		}

		numbers := strings.Split(parts[1], " ")
		for _, number := range numbers {
			item, err := strconv.Atoi(number)
			if err != nil {
				log.Fatalf("Problem parsing '%s': %s", number, err)
			}

			equation.numbers = append(equation.numbers, item)
		}

		equations = append(equations, equation)
	}

	return equations
}

func concatenate(a, b int) int {
	if a == 0 {
		return b
	}

	result, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return result
}

func check(equation Equation, index, result int, conc bool) bool {
	if result > equation.result {
		return false
	}

	if index >= len(equation.numbers) {
		if result == equation.result {
			return true
		}

		return false
	}

	resultAdd := check(equation, index+1, result+equation.numbers[index], conc)

	resultConc := false
	if conc {
		resultConc = check(equation, index+1, concatenate(result, equation.numbers[index]), conc)
	}

	if result == 0 {
		result++
	}
	resultMul := check(equation, index+1, result*equation.numbers[index], conc)

	if resultAdd {
		return resultAdd
	}

	if resultConc {
		return resultConc
	}

	return resultMul
}

func parts(equations []Equation) (int, int) {
	var part1 int
	var toRecheck []Equation
	for _, equation := range equations {
		if check(equation, 0, 0, false) {
			part1 += equation.result
		} else {
			toRecheck = append(toRecheck, equation)
		}
	}

	part2 := part1
	for _, equation := range toRecheck {
		if check(equation, 0, 0, true) {
			part2 += equation.result
		}
	}

	return part1, part2
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open %s!\n", filePath)
	}

	equations := readInput(file)
	part1, part2 := parts(equations)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
