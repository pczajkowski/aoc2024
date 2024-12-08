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

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func checkAntinode(antinode Point, matrix [][]byte, key byte) bool {
	if antinode.x < 0 || antinode.y < 0 || antinode.y >= len(matrix) || antinode.x >= len(matrix[0]) {
		return false
	}

	if matrix[antinode.y][antinode.x] == key {
		return false
	}

	return true
}

func getAntinodes(groups map[byte][]Point, matrix [][]byte) map[string]Point {
	antinodes := make(map[string]Point)

	for key, items := range groups {
		edge := len(items)
		for i := range items {
			for j := i + 1; j < edge; j++ {
				deltaX := items[j].x - items[i].x
				deltaY := items[j].y - items[i].y

				firstAntinode := Point{x: items[i].x - deltaX, y: items[i].y - deltaY}
				if checkAntinode(firstAntinode, matrix, key) {
					antinodes[firstAntinode.key()] = firstAntinode
				}

				secondAntinode := Point{x: items[j].x + deltaX, y: items[j].y + deltaY}
				if checkAntinode(secondAntinode, matrix, key) {
					antinodes[secondAntinode.key()] = secondAntinode
				}
			}
		}
	}

	return antinodes
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
	antinodes := getAntinodes(groups, matrix)
	fmt.Println("Part1:", len(antinodes))
}
