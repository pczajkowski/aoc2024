package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	AND = iota
	OR
	XOR
)

type Gate struct {
	id          string
	left, right string
	op          int
	value       int
}

func readInput(file *os.File) (map[string]Gate, map[string]Gate) {
	scanner := bufio.NewScanner(file)
	zs := make(map[string]Gate)
	gates := make(map[string]Gate)

	readingRegisters := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if readingRegisters {
				readingRegisters = false
				continue
			}

			break
		}

		var gate Gate
		if readingRegisters {
			parts := strings.Split(line, ": ")
			if len(parts) != 2 {
				log.Fatalf("Bad register line: %s", line)
			}

			gate.id = parts[0]

			n, err := fmt.Sscanf(parts[1], "%d", &gate.value)
			if n != 1 || err != nil {
				log.Fatalf("Bad input %s: %s", parts[1], err)
			}
		} else {
			var op string
			n, err := fmt.Sscanf(line, "%s %s %s -> %s", &gate.left, &op, &gate.right, &gate.id)
			if n != 4 || err != nil {
				log.Fatalf("Bad input %s: %s", line, err)
			}

			switch op {
			case "AND":
				gate.op = AND
			case "OR":
				gate.op = OR
			case "XOR":
				gate.op = XOR
			}

			gate.value = -1
		}

		if gate.id[0] == 'z' {
			zs[gate.id] = gate
		} else {
			gates[gate.id] = gate
		}

	}

	return gates, zs
}

func calculate(gate Gate, gates map[string]Gate) int {
	if gate.value != -1 {
		return gate.value
	}

	left := calculate(gates[gate.left], gates)
	right := calculate(gates[gate.right], gates)

	switch gate.op {
	case AND:
		gate.value = left & right
	case OR:
		gate.value = left | right
	case XOR:
		gate.value = left ^ right
	}

	return gate.value
}

func zsToNumber(zs map[string]Gate) int64 {
	arr := make([]byte, len(zs))
	keys := make([]string, len(zs))

	var index int
	for key, _ := range zs {
		keys[index] = key
		index++
	}

	slices.Sort(keys)
	slices.Reverse(keys)
	for i := range keys {
		value := zs[keys[i]].value
		switch value {
		case 0:
			arr[i] = '0'
		case 1:
			arr[i] = '1'
		}
	}

	result, err := strconv.ParseInt(string(arr), 2, 64)
	if err != nil {
		log.Fatalf("Can't convert binary %s: %s", arr, err)
	}

	return result
}

func calculateZs(zs, gates map[string]Gate) {
	for key, value := range zs {
		value.value = calculate(value, gates)
		zs[key] = value
	}
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

	gates, zs := readInput(file)
	calculateZs(zs, gates)
	fmt.Println("Part1:", zsToNumber(zs))
}
