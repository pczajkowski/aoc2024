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

var directions [][]int = [][]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1},
}

func checkPath(i, j int, matrix [][]byte, word string, index int, direction []int, rows, cols int) int {
	if i < 0 || i >= rows || j < 0 || j >= cols || matrix[i][j] != word[index] {
		return 0
	}

	if index == len(word)-1 {
		return 1
	}

	i += direction[0]
	j += direction[1]
	index++

	return checkPath(i, j, matrix, word, index, direction, rows, cols)
}

func check(i, j int, matrix [][]byte, word string, rows, cols int) int {
	var count int
	for _, direction := range directions {
		count += checkPath(i, j, matrix, word, 0, direction, rows, cols)
	}

	return count
}

func part1(matrix [][]byte, word string) int {
	var count int
	rows, cols := len(matrix), len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == word[0] {
				count += check(i, j, matrix, word, rows, cols)
			}
		}
	}

	return count
}

func isXmas(i, j int, matrix [][]byte, rows, cols int) bool {
	if i-1 < 0 || i+1 >= rows || j-1 < 0 || j+1 >= cols {
		return false
	}

	opposites := [][]int{{i - 1, j - 1}, {i - 1, j + 1}, {i + 1, j - 1}, {i + 1, j + 1}}
	for _, opposite := range opposites {
		if matrix[opposite[0]][opposite[1]] != 'M' && matrix[opposite[0]][opposite[1]] != 'S' {
			return false
		}
	}

	if matrix[opposites[0][0]][opposites[0][1]] == matrix[opposites[3][0]][opposites[3][1]] || matrix[opposites[1][0]][opposites[1][1]] == matrix[opposites[2][0]][opposites[2][1]] {
		return false
	}

	return true
}

func part2(matrix [][]byte) int {
	var count int
	rows, cols := len(matrix), len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 'A' {
				if isXmas(i, j, matrix, rows, cols) {
					count++
				}
			}
		}
	}

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
	fmt.Println("Part1:", part1(matrix, "XMAS"))
	fmt.Println("Part2:", part2(matrix))
}
