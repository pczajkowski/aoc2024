package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var mulRegex *regexp.Regexp = regexp.MustCompile(`mul\(\d+,\d+\)`)

func getResults(line string) int {
	var result int
	matches := mulRegex.FindAllString(line, -1)
	for _, match := range matches {
		mul := make([]int, 2)
		n, err := fmt.Sscanf(match, "mul(%d,%d)", &mul[0], &mul[1])
		if n != 2 || err != nil {
			log.Fatalf("Bad input: %s", err)
		}

		result += mul[0] * mul[1]
	}

	return result
}

func readInput(file *os.File) (int, []string) {
	scanner := bufio.NewScanner(file)
	var result int
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		lines = append(lines, line)
		result += getResults(line)
	}

	return result, lines
}

func part1(muls [][]int) int {
	var result int
	for _, mul := range muls {
		result += mul[0] * mul[1]
	}

	return result
}

func part2(lines []string) int {
	var result int
	multiply := true

	for _, line := range lines {
		var startIndex, endIndex int
		reading := true
		for reading {
			if multiply {
				index := strings.Index(line, "don't()")
				if index == -1 {
					endIndex = len(line)
					reading = false
				} else {
					multiply = false
					endIndex = index
				}

				if startIndex > endIndex {
					startIndex++
					continue
				}

				result += getResults(line[startIndex:endIndex])

				line = line[endIndex:]
				startIndex = 0
			} else {
				index := strings.Index(line, "do()")
				if index == -1 {
					reading = false
				} else {
					multiply = true
					startIndex = 0
					line = line[index:]
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

	part1, lines := readInput(file)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2(lines))
}
