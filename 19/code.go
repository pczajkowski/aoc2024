package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const delta byte = 97

func readInput(file *os.File) ([][]string, []string) {
	scanner := bufio.NewScanner(file)
	patterns := make([][]string, 25)
	var towels []string

	var patternsRead bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			patternsRead = true
			continue
		}

		if !patternsRead {
			parts := strings.Split(line, ", ")
			for _, part := range parts {
				patterns[part[0]-delta] = append(patterns[part[0]-delta], part)
			}
		} else {
			towels = append(towels, line)
		}
	}

	return patterns, towels
}

func checkTowel(towel string, index int, patterns [][]string) bool {
	if index >= len(towel) {
		return true
	}

	for _, pattern := range patterns[towel[index]-delta] {
		patternMatch := true
		for i := range pattern {
			if index+i >= len(towel) || pattern[i] != towel[index+i] {
				patternMatch = false
				break
			}
		}

		if patternMatch && checkTowel(towel, index+len(pattern), patterns) {
			return true
		}
	}

	return false
}

func part1(patterns [][]string, towels []string) int {
	var count int
	for _, towel := range towels {
		if checkTowel(towel, 0, patterns) {
			count++
		}
	}

	return count
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

	patterns, towels := readInput(file)
	fmt.Println("Part1:", part1(patterns, towels))
}
