package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func readInput(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)
	var muls [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		re := regexp.MustCompile(`mul\(\d+,\d+\)`)
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			mul := make([]int, 2)
			n, err := fmt.Sscanf(match, "mul(%d,%d)", &mul[0], &mul[1])
			if n != 2 || err != nil {
				log.Fatalf("Bad input: %s", err)
			}

			muls = append(muls, mul)
		}
	}

	return muls
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

	muls := readInput(file)
	fmt.Println(muls)
}
