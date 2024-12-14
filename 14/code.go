package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Robot struct {
	x, y   int
	vX, vY int
}

func readInput(file *os.File) []Robot {
	scanner := bufio.NewScanner(file)
	var robots []Robot

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var robot Robot
		n, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.vX, &robot.vY)
		if n != 4 || err != nil {
			log.Fatalf("Not able to parse robot '%s': %s", line, err)
		}

		robots = append(robots, robot)
	}

	return robots
}

func getMaxXY(robots []Robot) (int, int) {
	var maxX, maxY int
	for _, robot := range robots {
		if robot.x > maxX {
			maxX = robot.x
		}

		if robot.y > maxY {
			maxY = robot.y
		}
	}

	return maxX + 1, maxY + 1
}

func robotsAfter(robots []Robot, maxX, maxY, after int) []Robot {
	robotsMoved := make([]Robot, len(robots))
	for i, robot := range robots {
		robot.x = (robot.vX*after + robot.x) % maxX
		if robot.x < 0 {
			robot.x = maxX + robot.x
		}

		robot.y = (robot.vY*after + robot.y) % maxY
		if robot.y < 0 {
			robot.y = maxY + robot.y
		}

		robotsMoved[i] = robot
	}

	return robotsMoved
}

func part1(robots []Robot, maxX, maxY, after int) int {
	midX := maxX / 2
	midY := maxY / 2

	var q1, q2, q3, q4 int
	robotsMoved := robotsAfter(robots, maxX, maxY, after)
	for _, robot := range robotsMoved {
		if robot.x < midX && robot.y < midY {
			q1++
		}

		if robot.x > midX && robot.y < midY {
			q2++
		}

		if robot.x < midX && robot.y > midY {
			q3++
		}

		if robot.x > midX && robot.y > midY {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
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

	robots := readInput(file)
	maxX, maxY := getMaxXY(robots)
	fmt.Println("Part1:", part1(robots, maxX, maxY, 100))
}
