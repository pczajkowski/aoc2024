package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x  int
	steps int
}

func (p Point) key() string {
	return fmt.Sprintf("%d_%d", p.y, p.x)
}

func readInput(file *os.File) []Point {
	scanner := bufio.NewScanner(file)
	var points []Point

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var point Point
		n, err := fmt.Sscanf(line, "%d,%d", &point.x, &point.y)
		if n != 2 || err != nil {
			log.Fatalf("Not able to parse byte '%s': %s", line, err)
		}

		points = append(points, point)
	}

	return points
}

var directions [][]int = [][]int{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func getObstaclesMap(obstacles []Point, howMany, xMax, yMax int) map[string]bool {
	obstaclesMap := make(map[string]bool)

	for i := 0; i < howMany; i++ {
		if obstacles[i].x < xMax && obstacles[i].y < yMax {
			obstaclesMap[obstacles[i].key()] = true
		}
	}

	return obstaclesMap
}

func getMoves(current Point, obstaclesMap map[string]bool, xMax, yMax int) []Point {
	var moves []Point
	for _, direction := range directions {
		move := Point{x: current.x + direction[0], y: current.y + direction[1], steps: current.steps + 1}
		if move.x < 0 || move.y < 0 || move.x > xMax || move.y > yMax {
			continue
		}

		if !obstaclesMap[move.key()] {
			moves = append(moves, move)
		}
	}

	return moves
}

func hike(obstaclesMap map[string]bool, xMax, yMax int) int {
	steps := 1000000000
	visited := make(map[string]int)

	goal := Point{x: xMax, y: yMax}
	moves := []Point{Point{x: 0, y: 0}}
	for len(moves) > 0 {
		current := moves[0]
		moves = moves[1:]
		if current.x == goal.x && current.y == goal.y && current.steps < steps {
			steps = current.steps
		}

		newMoves := getMoves(current, obstaclesMap, xMax, yMax)
		for _, newMove := range newMoves {
			if visited[newMove.key()] == 0 || visited[newMove.key()] >= newMove.steps {
				moves = append(moves, newMove)
				visited[newMove.key()] = newMove.steps
			}
		}
	}

	return steps
}

func part1(obstacles []Point, howMany, xMax, yMax int) int {
	obstaclesMap := getObstaclesMap(obstacles, howMany, xMax, yMax)
	return hike(obstaclesMap, xMax, yMax)
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

	obstacles := readInput(file)
	fmt.Println("Part1:", part1(obstacles, 12, 6, 6))
}
