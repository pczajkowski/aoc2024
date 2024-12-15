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

func moveBox(x, y int, matrix [][]byte, xPlus, yPlus int) bool {
	field := matrix[y][x]
	if field == '#' {
		return false
	}

	if field == 'O' {
		if !moveBox(x+xPlus, y+yPlus, matrix, xPlus, yPlus) {
			return false
		}
	}

	matrix[y][x] = 'O'
	return true
}

func moveRobot(robot *Point, matrix [][]byte, x, y int) {
	matrix[robot.y][robot.x] = '.'
	field := matrix[robot.y+y][robot.x+x]
	if field == '#' {
		return
	}

	if field == 'O' {
		if !moveBox(robot.x+x, robot.y+y, matrix, x, y) {
			return
		}
	}

	robot.x += x
	robot.y += y
}

func processMoves(robot *Point, matrix [][]byte, moves []byte) {
	for _, move := range moves {
		switch move {
		case '^':
			moveRobot(robot, matrix, 0, -1)
			fmt.Println(robot, matrix)
		case 'v':
			moveRobot(robot, matrix, 0, 1)
			fmt.Println(robot, matrix)
		case '<':
			moveRobot(robot, matrix, -1, 0)
			fmt.Println(robot, matrix)
		case '>':
			moveRobot(robot, matrix, 1, 0)
			fmt.Println(robot, matrix)
		}
	}
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
	processMoves(robot, matrix, moves)
}
