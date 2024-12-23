package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

func getSets(computers map[string][]string) [][]string {
	var sets [][]string

	for key, value := range computers {
		if key[0] != 't' {
			continue
		}

		for i := range value {
			for _, subKey := range computers[value[i]] {
				if subKey == key {
					continue
				}

				subValues := computers[subKey]
				for j := range subValues {
					if subValues[j] == key {
						sets = append(sets, []string{key, value[i], subKey})
					}
				}
			}
		}
	}

	return sets
}

func part1(computers map[string][]string) int {
	sets := getSets(computers)
	for i := range sets {
		slices.Sort(sets[i])
	}

	unique := make(map[string]bool)
	for i := range sets {
		unique[strings.Join(sets[i], ",")] = true
	}

	return len(unique)
}

func contains(nodes []string, connections map[string]bool) bool {
	if len(connections) == 0 {
		return true
	}

	for key, _ := range connections {
		var isThere bool
		for i := range nodes {
			if nodes[i] == key {
				isThere = true
				break
			}
		}

		if !isThere {
			return false
		}
	}

	return true
}

func connected(key string, computers map[string][]string, connections map[string]bool) {
	if !contains(computers[key], connections) {
		return
	}

	connections[key] = true
	for _, value := range computers[key] {
		if connections[value] {
			continue
		}

		connected(value, computers, connections)
	}
}

func part2(computers map[string][]string) string {
	var allConnections []map[string]bool
	for key, _ := range computers {
		connections := make(map[string]bool)
		connected(key, computers, connections)

		allConnections = append(allConnections, connections)
	}

	var biggest int
	var connections map[string]bool
	for i := range allConnections {
		if len(allConnections[i]) > biggest {
			biggest = len(allConnections[i])
			connections = allConnections[i]
		}
	}

	var keys []string
	for key, _ := range connections {
		keys = append(keys, key)
	}

	slices.Sort(keys)
	return strings.Join(keys, ",")
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
	fmt.Println("Part1:", part1(computers))
	fmt.Println("Part2:", part2(computers))
}
