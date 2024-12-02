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

func safe(report []int) int {
	var direction int
	edge := len(report)

	for i := 1; i < edge; i++ {
		delta := report[i] - report[i-1]
		if delta == 0 || delta < -3 || delta > 3 {
			return i
		}

		if direction == 0 {
			direction = delta
		} else if direction < 0 && delta > 0 || direction > 0 && delta < 0 {
			return i
		}
	}

	return 0
}

func checkReports(reports [][]int) (int, int) {
	var part1, part2 int
	for _, report := range reports {
		bad := safe(report)
		if bad == 0 {
			part1++
		} else {
			bad = safe(append(report[:bad], report[bad+1:]...))
			if bad == 0 {
				part2++
			}
		}
	}

	return part1, part2
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
	part1, part2 := checkReports(reports)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part1+part2)
}
