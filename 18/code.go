package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x int
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
	fmt.Println(obstacles)
}
