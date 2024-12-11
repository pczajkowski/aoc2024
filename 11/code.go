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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	stones := readInput(os.Args[1])
	fmt.Println(stones)
}
