package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var directions [][]int = [][]int{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

type Point struct {
	x, y      int
	direction int
}

func (p *Point) key() string {
	return fmt.Sprintf("%d_%d", p.x, p.y)
}

func findGuard(line string) *Point {
	for i, char := range line {
		if char == '^' {
			return &Point{x: i}
		}
	}

	return nil
}

func readInput(file *os.File) (*Point, [][]byte) {
	scanner := bufio.NewScanner(file)
	var matrix [][]byte
	var guard *Point

	var y int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		matrix = append(matrix, []byte(line))
		if guard == nil {
			guard = findGuard(line)
			if guard != nil {
				guard.y = y
			}
		}

		y++
	}

	return guard, matrix
}

func walk(guard *Point, matrix [][]byte, xMax, yMax int, visited map[string]bool) {
	if guard.x < 0 || guard.x > xMax || guard.y < 0 || guard.y > yMax {
		return
	}

	if matrix[guard.y][guard.x] == '#' {
		guard.x -= directions[guard.direction][0]
		guard.y -= directions[guard.direction][1]
		guard.direction = (guard.direction + 1) % 4
	} else {
		visited[guard.key()] = true
		guard.x += directions[guard.direction][0]
		guard.y += directions[guard.direction][1]
	}

	walk(guard, matrix, xMax, yMax, visited)
}

func part1(guard *Point, matrix [][]byte) int {
	xMax := len(matrix[0]) - 1
	yMax := len(matrix) - 1
	visited := make(map[string]bool)

	walk(guard, matrix, xMax, yMax, visited)
	return len(visited)
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

	guard, matrix := readInput(file)
	fmt.Println("Part1:", part1(guard, matrix))
}
