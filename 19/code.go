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

func getPossibleTowers(patterns [][]string, towels []string) []string {
	var possible []string
	for _, towel := range towels {
		checked := make(map[string]bool)
		if checkTowel(towel, 0, patterns, checked) {
			possible = append(possible, towel)
		}
	}

	return possible
}

func checkTowel2(towel string, index int, patterns [][]string, badOnes map[string]bool) int {
	var count int
	if index >= len(towel) {
		return 1
	}

	edge := len(towel) - 1
	for _, pattern := range patterns[towel[index]-delta] {
		if index+len(pattern)-1 > edge {
			badOnes[patternAndIndex(pattern, index)] = true
			continue
		}

		if badOnes[patternAndIndex(pattern, index)] {
			continue
		}

		patternMatch := true
		for i := range pattern {
			if index+i >= len(towel) || pattern[i] != towel[index+i] {
				patternMatch = false
				break
			}
		}

		if patternMatch {
			count += checkTowel2(towel, index+len(pattern), patterns, badOnes)
		} else {
			badOnes[patternAndIndex(pattern, index)] = true
		}
	}

	return count
}

func getMax(patterns [][]string) int {
	var maxPatterns int
	for _, letter := range patterns {
		maxPatterns += len(letter)
	}

	return maxPatterns
}

func part2(patterns [][]string, towels []string) int {
	var count int
	for _, towel := range towels {
		badOnes := make(map[string]bool)
		count += checkTowel2(towel, 0, patterns, badOnes)
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
	possible := getPossibleTowers(patterns, towels)
	fmt.Println("Part1:", len(possible))
	fmt.Println("Part2:", part2(patterns, possible))
}
