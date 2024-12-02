package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var report []int
		numbers := strings.Split(line, " ")
		for _, number := range numbers {
			level, err := strconv.Atoi(number)
			if err != nil {
				log.Fatalf("Problem parsing input: %s", err)
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports
}

func check(a, b int, direction int) (bool, int) {
	delta := b - a
	if delta == 0 || delta < -3 || delta > 3 {
		return false, direction
	}

	if direction == 0 {
		direction = delta
	} else if direction < 0 && delta > 0 || direction > 0 && delta < 0 {
		return false, direction
	}

	return true, direction
}

func safe(report []int, skip bool) bool {
	var direction int
	var status bool
	edge := len(report)

	for i := 1; i < edge; i++ {
		status, direction = check(report[i], report[i-1], direction)
		if !status {
			if skip && i+1 < edge {
				status, direction = check(report[i+1], report[i-1], direction)
				if status {
					i++
					skip = false
					continue
				}
			}

			return false
		}
	}

	return true
}

func part(reports [][]int, skip bool) int {
	var result int
	for _, report := range reports {
		if safe(report, skip) {
			result++
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

	reports := readInput(file)
	fmt.Println("Part1:", part(reports, false))
	fmt.Println("Part2:", part(reports, true))
}
