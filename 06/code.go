package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var directions [][]int = [][]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1},
}

type Point struct {
	x, y      int
	direction int
}

func (p *Point) key() string {
	return fmt.Sprintf("%d_%s", p.x, p.y)
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
	fmt.Println(guard, matrix)
}
