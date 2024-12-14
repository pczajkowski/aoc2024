package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Robot struct {
	x, y   int
	vX, Vy int
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
		n, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.vX, &robot.Vy)
		if n != 4 || err != nil {
			log.Fatalf("Not able to parse robot '%s': %s", line, err)
		}

		robots = append(robots, robot)
	}

	return robots
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
	fmt.Println(robots)
}
