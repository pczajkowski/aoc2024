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
	cheatedAt *Point
}

func (p *Point) key() string {
	return fmt.Sprintf("%d_%d", p.y, p.x)
}

func findPoint(line string, mark byte) *Point {
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
	var start *Point

	var y int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		matrix = append(matrix, []byte(line))
		if start == nil {
			start = findPoint(line, 'S')
			if start != nil {
				start.y = y
			}
		}

		y++
	}

	return start, matrix
}

var directions [][]int = [][]int{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func getMoves(current Point, matrix [][]byte, xMax, yMax int, cheat bool, cheats map[string]bool) []Point {
	var moves []Point
	for _, direction := range directions {
		move := Point{x: current.x + direction[0], y: current.y + direction[1], cost: current.cost + 1, cheatedAt: current.cheatedAt}
		if move.x <= 0 || move.y <= 0 || move.x >= xMax || move.y >= yMax {
			continue
		}

		if matrix[move.y][move.x] == '#' {
			if cheat && !cheats[move.key()] && move.cheatedAt == nil {
				move.cheatedAt = &move
				moves = append(moves, move)
			}

			continue
		}

		moves = append(moves, move)
	}

	return moves
}

func hike(start *Point, matrix [][]byte, xMax, yMax int, cheat bool, cheats map[string]bool) int {
	cost := 1000000000
	visited := make(map[string]int)
	visited[start.key()] = start.cost

	moves := []Point{*start}
	for len(moves) > 0 {
		current := moves[0]
		moves = moves[1:]
		if matrix[current.y][current.x] == 'E' && current.cost <= cost {
			cost = current.cost
			if cheat && current.cheatedAt != nil {
				cheats[current.cheatedAt.key()] = true
			}
		}

		newMoves := getMoves(current, matrix, xMax, yMax, cheat, cheats)
		for _, newMove := range newMoves {
			if visited[newMove.key()] == 0 || visited[newMove.key()] >= newMove.cost {
				moves = append(moves, newMove)
				visited[newMove.key()] = newMove.cost
			}
		}
	}

	return cost
}

func part1(start *Point, matrix [][]byte, atLeast int) int {
	xMax := len(matrix[0]) - 1
	yMax := len(matrix) - 1

	cheats := make(map[string]bool)
	bestWithoutCheating := hike(start, matrix, xMax, yMax, false, cheats)
	var count int
	savings := make(map[int]int)
	for {
		score := hike(start, matrix, xMax, yMax, true, cheats)
		if score >= bestWithoutCheating {
			break
		}

		saving := bestWithoutCheating - score
		if saving >= atLeast {
			count++
		}
		savings[saving]++
	}
	fmt.Println(savings)
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

	start, matrix := readInput(file)
	fmt.Println("Part1:", part1(start, matrix, 1))
}
