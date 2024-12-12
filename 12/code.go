package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x  int
	value byte
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

func getRegion(start Point, matrix [][]byte, xMax, yMax int, checked map[string]bool) []Point {
	plots := []Point{start}
	var region []Point

	for len(plots) > 0 {
		current := plots[0]
		plots = plots[1:]

		if !checked[current.key()] {
			region = append(region, current)
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

func part1(matrix [][]byte) int {
	xMax := len(matrix[0])
	yMax := len(matrix)
	checked := make(map[string]bool)

	for y := range matrix {
		for x := range matrix[y] {
			current := Point{x: x, y: y, value: matrix[y][x]}
			if !checked[current.key()] {
				region := getRegion(current, matrix, xMax, yMax, checked)
				fmt.Println(region)
			}
		}
	}

	return 0
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
	fmt.Println(part1(matrix))
}
