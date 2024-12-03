package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func readInput(file *os.File) ([][]int, []string) {
	scanner := bufio.NewScanner(file)
	var muls [][]int
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		lines = append(lines, line)

		re := regexp.MustCompile(`mul\(\d+,\d+\)`)
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			mul := make([]int, 2)
			n, err := fmt.Sscanf(match, "mul(%d,%d)", &mul[0], &mul[1])
			if n != 2 || err != nil {
				log.Fatalf("Bad input: %s", err)
			}

			muls = append(muls, mul)
		}
	}

	return muls, lines
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
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	var startIndex int
	for _, line := range lines {
		edge := len(line) - 1
		endIndex := edge
		reading := true
		for reading {
			if multiply {
				index := strings.Index(line[startIndex:edge], "don't()")
				multiply = false
				if index == -1 {
					endIndex = edge
					reading = false
				} else {
					endIndex = index
				}

				matches := re.FindAllString(line[startIndex:endIndex], -1)
				for _, match := range matches {
					mul := make([]int, 2)
					n, err := fmt.Sscanf(match, "mul(%d,%d)", &mul[0], &mul[1])
					if n != 2 || err != nil {
						log.Fatalf("Bad input: %s", err)
					}

					result += mul[0] * mul[1]
				}
			} else {
				index := strings.Index(line[startIndex:edge], "do()")
				if index == -1 {
					startIndex = 0
					reading = false
				} else {
					multiply = true
					startIndex = index
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

	muls, lines := readInput(file)
	fmt.Println("Part1:", part1(muls))
	fmt.Println("Part2:", part2(lines))
}
