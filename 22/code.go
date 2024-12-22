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

func generateNumber(number int, iterations int) (int, []int) {
	lastDigits := make([]int, iterations+1)
	lastDigits[0] = number % 10

	for i := 0; i < iterations; i++ {
		number = calculate(number)
		lastDigits[i+1] = number % 10
	}

	return number, lastDigits
}

func part1(numbers []int, iterations int) (int, [][]int) {
	var result int
	var allLastDigits [][]int
	for _, number := range numbers {
		newNumber, lastDigits := generateNumber(number, iterations)
		result += newNumber

		allLastDigits = append(allLastDigits, lastDigits)
	}

	return result, allLastDigits
}

func isDesiredSequence(sequence, desiredSequence []int) bool {
	if len(sequence) != len(desiredSequence) {
		return false
	}

	for i := range sequence {
		if sequence[i] != desiredSequence[i] {
			return false
		}
	}

	return true
}

func part2(allLastDigits [][]int, iterations int, desiredSequence []int) int {
	var result int
	for _, lastDigits := range allLastDigits {
		var sequence []int
		for i := 1; i < iterations; i++ {
			sequence = append(sequence, lastDigits[i]-lastDigits[i-1])
			if len(sequence) > 3 {
				if isDesiredSequence(sequence[len(sequence)-4:], desiredSequence) {
					result += lastDigits[i]
				}
			}
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

	numbers := readInput(file)
	part1Result, allLastDigits := part1(numbers, 2000)
	fmt.Println("Part1:", part1Result)
	fmt.Println("Part2:", part2(allLastDigits, 2001, []int{-2, 1, -1, 3}))
}
