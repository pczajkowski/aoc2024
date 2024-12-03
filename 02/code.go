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

func bigger(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func check(arr []int, direction int) (bool, bool) {
	edge := len(arr)
	lengths := make([]int, edge)

	for k := 0; k < edge; k++ {
		lengths[k] = 1
		for i := 0; i < k; i++ {
			delta := arr[k] - arr[i]
			if direction < 0 && delta > 0 || direction > 0 && delta < 0 {
				continue
			}

			if delta != 0 && delta <= 3 && delta >= -3 {
				lengths[k] = bigger(lengths[k], lengths[i]+1)
			}
		}
	}

	return lengths[edge-1] == edge, lengths[edge-1] == edge-1
}

func checkReports(reports [][]int) (int, int) {
	var part1, part2 int
	for _, report := range reports {
		direction := report[1] - report[0]
		one, two := check(report, direction)
		if one {
			part1++
			continue
		} else if two {
			part2++
			continue
		}

		one, two = check(report, -direction)
		if one {
			part1++
			continue
		} else if two {
			part2++
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
