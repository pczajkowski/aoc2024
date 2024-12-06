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

func walk(guard *Point, matrix [][]byte, xMax, yMax int, visited map[string]int) {
	if guard.x < 0 || guard.x > xMax || guard.y < 0 || guard.y > yMax {
		return
	}

	if matrix[guard.y][guard.x] == '#' {
		guard.x -= directions[guard.direction][0]
		guard.y -= directions[guard.direction][1]
		guard.direction = (guard.direction + 1) % 4
	} else {
		visited[guard.key()]++
		guard.x += directions[guard.direction][0]
		guard.y += directions[guard.direction][1]
	}

	walk(guard, matrix, xMax, yMax, visited)
}

func part1(guard *Point, matrix [][]byte) map[string]int {
	xMax := len(matrix[0]) - 1
	yMax := len(matrix) - 1
	visited := make(map[string]int)

	newGuard := Point{x: guard.x, y: guard.y}
	walk(&newGuard, matrix, xMax, yMax, visited)
	return visited
}

func walkWithCheck(guard *Point, matrix [][]byte, xMax, yMax int, visited map[string]int) bool {
	if guard.x < 0 || guard.x > xMax || guard.y < 0 || guard.y > yMax {
		return true
	}

	if matrix[guard.y][guard.x] == '#' {
		guard.x -= directions[guard.direction][0]
		guard.y -= directions[guard.direction][1]
		guard.direction = (guard.direction + 1) % 4
	} else {
		visited[guard.key()]++
		if visited[guard.key()] > 4 {
			return false
		}

		guard.x += directions[guard.direction][0]
		guard.y += directions[guard.direction][1]
	}

	return walkWithCheck(guard, matrix, xMax, yMax, visited)
}

func part2(guard *Point, matrix [][]byte, expected map[string]int) int {
	xMax := len(matrix[0]) - 1
	yMax := len(matrix) - 1
	var count int

	for key, _ := range expected {
		var x, y int
		fmt.Sscanf(key, "%d_%d", &x, &y)
		matrix[y][x] = '#'

		visited := make(map[string]int)
		newGuard := Point{x: guard.x, y: guard.y}
		if !walkWithCheck(&newGuard, matrix, xMax, yMax, visited) {
			count++
		}

		matrix[y][x] = '.'
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

	guard, matrix := readInput(file)
	visited := part1(guard, matrix)
	fmt.Println("Part1:", len(visited))
	fmt.Println("Part2:", part2(guard, matrix, visited))
}
