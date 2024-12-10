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

func key(p Point) string {
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

func getMoves(reindeer Point, matrix [][]byte, xMax, yMax int) []Point {
	var moves []Point
	for _, direction := range directions {
		move := Point{x: reindeer.x + direction[0], y: reindeer.y + direction[1]}
		if move.x >= 0 && move.x < xMax &&
			move.y >= 0 && move.y < yMax && matrix[move.y][move.x]-reindeer.value == 1 {
			move.value = matrix[move.y][move.x]
			moves = append(moves, move)
		}
	}

	return moves
}

func hike(reindeer Point, matrix [][]byte, xMax, yMax int, hash func(Point) string) int {
	var nines int
	visited := make(map[string]bool)

	moves := []Point{reindeer}
	for len(moves) > 0 {
		current := moves[0]
		moves = moves[1:]
		if matrix[current.y][current.x] == '9' {
			nines++
		}

		newMoves := getMoves(current, matrix, xMax, yMax)
		for _, newMove := range newMoves {
			if !visited[hash(newMove)] {
				moves = append(moves, newMove)
				visited[hash(newMove)] = true
			}
		}
	}

	return nines
}

func part1(matrix [][]byte) int {
	var result int
	xMax := len(matrix[0])
	yMax := len(matrix)

	for y := range matrix {
		for x := range matrix[y] {
			if matrix[y][x] == '0' {
				reindeer := Point{x: x, y: y, value: '0'}
				result += hike(reindeer, matrix, xMax, yMax, key)
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
