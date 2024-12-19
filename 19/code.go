package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const delta byte = 97

func readInput(file *os.File) ([][]string, []string) {
	scanner := bufio.NewScanner(file)
	patterns := make([][]string, 25)
	var towels []string

	var patternsRead bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			patternsRead = true
			continue
		}

		if !patternsRead {
			parts := strings.Split(line, ", ")
			for _, part := range parts {
				patterns[part[0]-delta] = append(patterns[part[0]-delta], part)
			}
		} else {
			towels = append(towels, line)
		}
	}

	return patterns, towels
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

	patterns, towels := readInput(file)
	fmt.Println(patterns, towels)
}
