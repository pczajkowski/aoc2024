package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x   int
	value  byte
	fences int
}

func (p *Point) key() string {
	return fmt.Sprintf("%d_%d", p.y, p.x)
}

func readInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	var matrix [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		matrix = append(matrix, []byte(line))
	}

	return matrix
}

var directions [][]int = [][]int{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func getPlots(current Point, matrix [][]byte, xMax, yMax int) []Point {
	var plots []Point
	for _, direction := range directions {
		plot := Point{x: current.x + direction[0], y: current.y + direction[1]}
		if plot.x >= 0 && plot.x < xMax &&
			plot.y >= 0 && plot.y < yMax && matrix[plot.y][plot.x] == current.value {
			plot.value = matrix[plot.y][plot.x]
			plots = append(plots, plot)
		}
	}

	return plots
}

func countFences(current Point, matrix [][]byte, xMax, yMax int) int {
	var fences int
	for _, direction := range directions {
		plot := Point{x: current.x + direction[0], y: current.y + direction[1]}
		if plot.x >= 0 && plot.x < xMax &&
			plot.y >= 0 && plot.y < yMax && matrix[plot.y][plot.x] == current.value {
			continue
		}

		fences++
	}

	return fences
}

func getRegion(start Point, matrix [][]byte, xMax, yMax int, checked map[string]bool) map[string]Point {
	plots := []Point{start}
	region := make(map[string]Point)

	for len(plots) > 0 {
		current := plots[0]
		plots = plots[1:]

		if !checked[current.key()] {
			if current.fences == 0 {
				current.fences = countFences(current, matrix, xMax, yMax)
			}
			region[current.key()] = current
		}

		checked[current.key()] = true

		newPlots := getPlots(current, matrix, xMax, yMax)
		for _, plot := range newPlots {
			if !checked[plot.key()] {
				plots = append(plots, plot)
			}
		}
	}

	return region
}

func calculate(region map[string]Point) int {
	var perimeter int
	for _, plot := range region {
		perimeter += plot.fences
	}

	return perimeter * len(region)
}

func part1(matrix [][]byte) int {
	xMax := len(matrix[0])
	yMax := len(matrix)
	checked := make(map[string]bool)
	var result int

	for y := range matrix {
		for x := range matrix[y] {
			current := Point{x: x, y: y, value: matrix[y][x]}
			if !checked[current.key()] {
				current.fences = countFences(current, matrix, xMax, yMax)
				region := getRegion(current, matrix, xMax, yMax, checked)
				result += calculate(region)
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

	matrix := readInput(file)
	fmt.Println("Part1:", part1(matrix))
}
