package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(file *os.File) map[string][]string {
	scanner := bufio.NewScanner(file)
	computers := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			log.Fatalf("Bad input: %s", line)
		}

		computers[parts[0]] = append(computers[parts[0]], parts[1])
		computers[parts[1]] = append(computers[parts[1]], parts[0])
	}

	return computers
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

	computers := readInput(file)
	fmt.Println(computers)
}
