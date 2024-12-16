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
			}

			y++
		}

	}

	return reindeer, matrix
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
	fmt.Println(reindeer, matrix)
}
