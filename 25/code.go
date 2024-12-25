package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readLine(line string, arr []int) {
	for i := range line {
		if line[i] == '#' {
			arr[i]++
		}
	}
}

func readInput(file *os.File) ([][]int, [][]int) {
	scanner := bufio.NewScanner(file)
	var locks, keys [][]int

	var isKey bool
	var index int
	arr := make([]int, 5)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if isKey {
				keys = append(keys, arr)
			} else {
				locks = append(locks, arr)
			}

			arr = make([]int, 5)
			index = 0
			continue
		}

		if index == 0 {
			isKey = line[0] == '.'
		}

		if index != 0 && index != 6 {
			readLine(line, arr)
		}

		index++
	}

	if isKey {
		keys = append(keys, arr)
	} else {
		locks = append(locks, arr)
	}

	return locks, keys
}

func countMatches(locks, keys [][]int) int {
	var count int
	for _, lock := range locks {
		for _, key := range keys {
			fits := true
			for i := range lock {
				if lock[i]+key[i] > 5 {
					fits = false
					break
				}
			}

			if fits {
				count++
			}
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

	locks, keys := readInput(file)
	fmt.Println("Part1:", countMatches(locks, keys))
}
