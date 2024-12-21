package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	y, x int
	cost int
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

func readInput(file *os.File) (*Point, *Point, [][]byte) {
	scanner := bufio.NewScanner(file)
	var matrix [][]byte
	var start, finish *Point

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

		if finish == nil {
			finish = findPoint(line, 'E')
			if finish != nil {
				finish.y = y
			}
		}

		y++
	}

	return start, finish, matrix
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

	start, finish, matrix := readInput(file)
	fmt.Println(start, finish, matrix)
}
