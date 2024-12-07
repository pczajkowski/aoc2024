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

func check(equation Equation, index, result int) bool {
	if result > equation.result {
		return false
	}

	if index >= len(equation.numbers) {
		if result == equation.result {
			return true
		}

		return false
	}

	resultAdd := check(equation, index+1, result+equation.numbers[index])
	resultMul := check(equation, index+1, result*equation.numbers[index])

	if resultAdd {
		return resultAdd
	}

	return resultMul
}

func part1(equations []Equation) int {
	var result int
	for _, equation := range equations {
		if check(equation, 0, 1) {
			result += equation.result
		}
	}

	return result
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
	fmt.Println("Part1:", part1(equations))
}
