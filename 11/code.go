package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) []int {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var stones []int

	parts := strings.Split(strings.Trim(string(data), "\n"), " ")
	for _, part := range parts {
		stone, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("Bad input %s: %s", part, err)
		}

		stones = append(stones, stone)
	}

	return stones
}

func splitStone(stoneString string) (int, int) {
	half := len(stoneString) / 2

	first, err1 := strconv.Atoi(stoneString[:half])
	if err1 != nil {
		log.Fatalf("Can't convert %s: %s", stoneString[:half], err1)
	}

	second, err2 := strconv.Atoi(stoneString[half:])
	if err1 != nil {
		log.Fatalf("Can't convert %s: %s", stoneString[half:], err2)
	}

	return first, second
}

func processStones(stones []int) []int {
	var newStones []int
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}

		stoneString := fmt.Sprintf("%d", stone)
		if len(stoneString)%2 == 0 {
			first, second := splitStone(stoneString)
			newStones = append(newStones, first)
			newStones = append(newStones, second)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}

	return newStones
}

func part1(stones []int) int {
	for i := 0; i < 25; i++ {
		stones = processStones(stones)
	}

	return len(stones)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	stones := readInput(os.Args[1])
	fmt.Println("Part1:", part1(stones))
}
