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

func findRobot(line string) *Point {
	for i, char := range line {
		if char == '@' {
			return &Point{x: i}
		}
	}

	return nil
}

func readInput(file *os.File) (*Point, [][]byte, []byte) {
	scanner := bufio.NewScanner(file)
	var matrix [][]byte
	var moves []byte
	var robot *Point

	var y int
	readingMatrix := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if readingMatrix {
				readingMatrix = false
				continue
			}

			break
		}

		if readingMatrix {
			matrix = append(matrix, []byte(line))
			if robot == nil {
				robot = findRobot(line)
				if robot != nil {
					robot.y = y
				}

				y++
			}
		} else {
			moves = append(moves, []byte(line)...)
		}
	}

	return robot, matrix, moves
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

	robot, matrix, moves := readInput(file)
	fmt.Println(robot, matrix, moves)
}
