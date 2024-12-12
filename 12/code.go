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

func (p *Point) key() string {
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
	fmt.Println(matrix)
}
