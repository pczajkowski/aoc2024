package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x      int
	cost      int
	direction []int
}

func (p *Point) key() string {
	return fmt.Sprintf("%d_%d", p.y, p.x)
}

func findReindeer(line string, mark byte) *Point {
	for i := range line {
		if line[i] == mark {
			return &Point{x: i}
		}
	}

	return nil
}

func readInput(file *os.File) (*Point, [][]byte) {
	scanner := bufio.NewScanner(file)
	var matrix [][]byte
	var reindeer *Point

	var y int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		matrix = append(matrix, []byte(line))
		if reindeer == nil {
			reindeer = findReindeer(line, 'S')
			if reindeer != nil {
				reindeer.y = y
				reindeer.direction = []int{1, 0}
			}

			y++
		}

	}

	return reindeer, matrix
}

var directions [][]int = [][]int{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func getMoves(reindeer Point, matrix [][]byte) []Point {
	var moves []Point
	for _, direction := range directions {
		move := Point{x: reindeer.x + direction[0], y: reindeer.y + direction[1], cost: reindeer.cost, direction: direction}
		if matrix[move.y][move.x] != '#' {
			if reindeer.direction[0] == direction[0] && reindeer.direction[1] == direction[1] {
				move.cost++
			} else {
				move.cost += 1000
			}

			moves = append(moves, move)
		}
	}

	return moves
}

func hike(reindeer *Point, matrix [][]byte) int {
	cost := 1000000000
	visited := make(map[string]int)

	moves := []Point{*reindeer}
	for len(moves) > 0 {
		current := moves[0]
		moves = moves[1:]
		if matrix[current.y][current.x] == 'E' && current.cost < cost {
			cost = current.cost
		}

		newMoves := getMoves(current, matrix)
		for _, newMove := range newMoves {
			if visited[newMove.key()] == 0 || visited[newMove.key()] > newMove.cost {
				moves = append(moves, newMove)
				visited[newMove.key()] = newMove.cost
			}
		}
	}

	return cost
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

	reindeer, matrix := readInput(file)
	fmt.Println("Part1:", hike(reindeer, matrix))
}
