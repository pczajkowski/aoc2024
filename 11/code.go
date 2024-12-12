package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) map[int]int {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	stones := make(map[int]int)

	parts := strings.Split(strings.Trim(string(data), "\n"), " ")
	for _, part := range parts {
		stone, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("Bad input %s: %s", part, err)
		}

		stones[stone]++
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

func processStones(stones map[int]int) map[int]int {
	newStones := make(map[int]int)
	for key, value := range stones {
		if key == 0 {
			newStones[1] += value
		} else if stoneString := fmt.Sprintf("%d", key); len(stoneString)%2 == 0 {
			first, second := splitStone(stoneString)
			newStones[first] += value
			newStones[second] += value
		} else {
			m := key * 2024
			newStones[m] += value
		}
	}

	return newStones
}

func part1(stones map[int]int) int {
	for i := 0; i < 25; i++ {
		stones = processStones(stones)
	}

	var result int
	for _, value := range stones {
		result += value
	}

	return result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	stones := readInput(os.Args[1])
	fmt.Println("Part1:", part1(stones))
}
