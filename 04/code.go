package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file *os.File) [][]byte {
	scanner := bufio.NewScanner(file)
	var matrix [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		matrix = append(matrix, []byte(line))
	}

	return matrix
}

func countWordOccurrences(matrix [][]byte, word string) int {
	rows, cols := len(matrix), len(matrix[0])
	count := 0
	found := make(map[string]bool)

	var dfs func(i, j, wordIndex int, wordSoFar string)
	dfs = func(i, j, wordIndex int, wordSoFar string) {
		if wordIndex == len(word) {
			if !found[wordSoFar] {
				found[wordSoFar] = true
				count++
			}

			return
		}

		if i < 0 || i >= rows || j < 0 || j >= cols || matrix[i][j] != word[wordIndex] {
			return
		}

		wordSoFar = fmt.Sprintf("%s%d_%d%c", wordSoFar, i, j, matrix[i][j])

		dfs(i+1, j, wordIndex+1, wordSoFar)   // Down
		dfs(i-1, j, wordIndex+1, wordSoFar)   // Up
		dfs(i, j+1, wordIndex+1, wordSoFar)   // Right
		dfs(i, j-1, wordIndex+1, wordSoFar)   // Left
		dfs(i+1, j+1, wordIndex+1, wordSoFar) // Diagonal down-right
		dfs(i-1, j-1, wordIndex+1, wordSoFar) // Diagonal up-left
		dfs(i+1, j-1, wordIndex+1, wordSoFar) // Diagonal down-left
		dfs(i-1, j+1, wordIndex+1, wordSoFar) // Diagonal up-right
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == word[0] {
				dfs(i, j, 0, "")
			}
		}
	}
	fmt.Println(found)
	return count
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

	matrix := readInput(file)
	fmt.Println("Part1:", countWordOccurrences(matrix, "XMAS"))
}
