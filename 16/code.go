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
	parent    *Point
}

func (p *Point) key() string {
	return fmt.Sprintf("%d_%d", p.y, p.x)
}

func (p *Point) keyWithDirection() string {
	return fmt.Sprintf("%d_%d_%d_%d", p.y, p.x, p.direction[1], p.direction[0])
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
		move := Point{x: reindeer.x + direction[0], y: reindeer.y + direction[1], cost: reindeer.cost, direction: direction, parent: &reindeer}
		if matrix[move.y][move.x] != '#' {
			if reindeer.direction[0] == direction[0] && reindeer.direction[1] == direction[1] {
				move.cost++
			} else {
				move.cost += 1001
			}

			moves = append(moves, move)
		}
	}

	return moves
}

func hike(reindeer *Point, matrix [][]byte) (int, []*Point) {
	cost := 1000000000
	visited := make(map[string]int)
	paths := make(map[int][]*Point)

	moves := []Point{*reindeer}
	for len(moves) > 0 {
		current := moves[0]
		moves = moves[1:]
		if matrix[current.y][current.x] == 'E' && current.cost <= cost {
			cost = current.cost
			paths[cost] = append(paths[cost], &current)
		}

		newMoves := getMoves(current, matrix)
		for _, newMove := range newMoves {
			if visited[newMove.keyWithDirection()] == 0 || visited[newMove.keyWithDirection()] >= newMove.cost {
				moves = append(moves, newMove)
				visited[newMove.keyWithDirection()] = newMove.cost
			}
		}
	}

	return cost, paths[cost]
}

func countTiles(paths []*Point) int {
	var count int
	checked := make(map[string]bool)

	for _, path := range paths {
		current := path
		for current != nil {
			if !checked[current.key()] {
				checked[current.key()] = true
				count++
			}

			current = current.parent
		}
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

	reindeer, matrix := readInput(file)
	lowestCost, bestPaths := hike(reindeer, matrix)
	fmt.Println("Part1:", lowestCost)
	fmt.Println("Part2:", countTiles(bestPaths))
}
