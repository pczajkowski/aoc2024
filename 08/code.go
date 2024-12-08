package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

func (p *Point) key() string {
	return fmt.Sprintf("%d_%d", p.y, p.x)
}

func readInput(file *os.File) (map[byte][]Point, [][]byte) {
	scanner := bufio.NewScanner(file)
	var matrix [][]byte
	groups := make(map[byte][]Point)

	var y int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		matrix = append(matrix, []byte(line))
		for x := range line {
			if line[x] != '.' {
				groups[line[x]] = append(groups[line[x]], Point{x: x, y: y})
			}
		}

		y++
	}

	return groups, matrix
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

	groups, matrix := readInput(file)
	fmt.Println(groups, matrix)
}
