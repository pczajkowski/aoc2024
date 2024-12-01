package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func readInput(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)
	var lefts, rights []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var left, right int
		n, err := fmt.Sscanf(line, "%d %d", &left, &right)
		if n != 2 || err != nil {
			log.Fatalf("Bad input: %s", line)
		}

		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	return lefts, rights
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func part1(lefts, rights []int) int {
	slices.Sort(lefts)
	slices.Sort(rights)

	var distance int
	for i := range lefts {
		distance += abs(rights[i] - lefts[i])
	}

	return distance
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

	lefts, rights := readInput(file)
	fmt.Println("Part1:", part1(lefts, rights))
}
