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

func checkAntinode(antinode Point, matrix [][]byte) bool {
	if antinode.x < 0 || antinode.y < 0 || antinode.y >= len(matrix) || antinode.x >= len(matrix[0]) {
		return false
	}

	return true
}

func getAntinodes(groups map[byte][]Point, matrix [][]byte, multi bool) map[string]Point {
	antinodes := make(map[string]Point)

	for _, items := range groups {
		edge := len(items)
		for i := range items {
			for j := i + 1; j < edge; j++ {
				deltaX := items[j].x - items[i].x
				deltaY := items[j].y - items[i].y

				firstAntinode := Point{x: items[i].x - deltaX, y: items[i].y - deltaY}
				for checkAntinode(firstAntinode, matrix) {
					antinodes[firstAntinode.key()] = firstAntinode
					if multi {
						firstAntinode = Point{x: firstAntinode.x - deltaX, y: firstAntinode.y - deltaY}
					} else {
						firstAntinode = Point{x: -1, y: -1}
					}
				}

				secondAntinode := Point{x: items[j].x + deltaX, y: items[j].y + deltaY}
				for checkAntinode(secondAntinode, matrix) {
					antinodes[secondAntinode.key()] = secondAntinode
					if multi {
						secondAntinode = Point{x: secondAntinode.x + deltaX, y: secondAntinode.y + deltaY}
					} else {
						secondAntinode = Point{x: -1, y: -1}
					}
				}
			}
		}
	}

	return antinodes
}

func part2(groups map[byte][]Point, matrix [][]byte) int {
	antinodes := getAntinodes(groups, matrix, true)

	for _, items := range groups {
		if len(items) > 1 {
			for _, item := range items {
				antinodes[item.key()] = item
			}
		}
	}

	return len(antinodes)
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
	fmt.Println("Part1:", len(getAntinodes(groups, matrix, false)))
	fmt.Println("Part2:", part2(groups, matrix))
}
