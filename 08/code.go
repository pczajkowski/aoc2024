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

func checkInLine(antinode Point, antenna Point) bool {
	return antinode.x == antenna.x || antinode.y == antenna.y || abs(antinode.x-antenna.x) == abs(antinode.y-antenna.y)
}

func getAntinodes(groups map[byte][]Point, matrix [][]byte, multi bool) map[string]Point {
	antinodes := make(map[string]Point)

	for key, items := range groups {
		edge := len(items)
		for i := range items {
			for j := i + 1; j < edge; j++ {
				deltaX := items[j].x - items[i].x
				deltaY := items[j].y - items[i].y

				firstAntinode := Point{x: items[i].x - deltaX, y: items[i].y - deltaY}
				for checkAntinode(firstAntinode, matrix, key) {
					antinodes[firstAntinode.key()] = firstAntinode
					if multi {
						for _, antenna := range items {
							_, ok := antinodes[antenna.key()]
							if ok {
								continue
							}

							if checkInLine(firstAntinode, antenna) {
								antinodes[antenna.key()] = antenna
							}
						}

						firstAntinode = Point{x: firstAntinode.x - deltaX, y: firstAntinode.y - deltaY}
					} else {
						firstAntinode = Point{x: -1, y: -1}
					}
				}

				secondAntinode := Point{x: items[j].x + deltaX, y: items[j].y + deltaY}
				for checkAntinode(secondAntinode, matrix, key) {
					antinodes[secondAntinode.key()] = secondAntinode
					if multi {
						for _, antenna := range items {
							_, ok := antinodes[antenna.key()]
							if ok {
								continue
							}

							if checkInLine(secondAntinode, antenna) {
								antinodes[antenna.key()] = antenna
							}
						}

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
	fmt.Println("Part2:", len(getAntinodes(groups, matrix, true)))
}
