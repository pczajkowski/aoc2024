package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Can't parse a number %s: %s", line, err)
		}

		numbers = append(numbers, number)
	}

	return numbers
}

func calculate(number int) int {
	a := number * 64
	number = (number ^ a) % 16777216

	b := number / 32
	number = (number ^ b) % 16777216

	c := number * 2048
	number = (number ^ c) % 16777216

	return number
}

func generateNumber(number int, iterations int) int {
	for i := 0; i < iterations; i++ {
		number = calculate(number)
	}

	return number
}

func part1(numbers []int, iterations int) int {
	var result int
	for _, number := range numbers {
		result += generateNumber(number, iterations)
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

	numbers := readInput(file)
	fmt.Println("Part1:", part1(numbers, 2000))
}
