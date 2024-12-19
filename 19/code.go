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

func patternAndIndex(pattern string, index int) string {
	return fmt.Sprintf("%s_%d", pattern, index)
}

func checkTowel(towel string, index int, patterns [][]string, checked map[string]bool) bool {
	if index >= len(towel) {
		return true
	}

	for _, pattern := range patterns[towel[index]-delta] {
		if checked[patternAndIndex(pattern, index)] {
			continue
		}

		checked[patternAndIndex(pattern, index)] = true
		patternMatch := true
		for i := range pattern {
			if index+i >= len(towel) || pattern[i] != towel[index+i] {
				patternMatch = false
				break
			}
		}

		if patternMatch && checkTowel(towel, index+len(pattern), patterns, checked) {
			return true
		}
	}

	return false
}

func part1(patterns [][]string, towels []string) int {
	var count int
	for _, towel := range towels {
		checked := make(map[string]bool)
		if checkTowel(towel, 0, patterns, checked) {
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
