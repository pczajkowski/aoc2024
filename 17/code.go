package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Register struct {
	id    byte
	value int
}

func readInput(file *os.File) ([]Register, []int) {
	scanner := bufio.NewScanner(file)
	var registers []Register
	var program []int

	var registersRead bool
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			registersRead = true
			continue
		}

		if !registersRead {
			var register Register
			n, err := fmt.Sscanf(line, "Register %c: %d", &register.id, &register.value)
			if n != 2 || err != nil {
				log.Fatalf("Not able to parse register '%s': %s", line, err)
			}

			registers = append(registers, register)
		} else {
			numString := strings.TrimPrefix(line, "Program: ")
			parts := strings.Split(numString, ",")
			for _, part := range parts {
				num, err := strconv.Atoi(part)
				if err != nil {
					log.Fatalf("Not able to convert %s: %s", part, err)
				}

				program = append(program, num)
			}
		}
	}

	return registers, program
}

func getCombo(operand int, registers []Register) int {
	if operand >= 0 && operand <= 3 {
		return operand
	}

	switch operand {
	case 4:
		return registers[0].value
	case 5:
		return registers[1].value
	case 6:
		return registers[2].value
	case 7:
		log.Fatal("Bad instruction!")
	}

	return -1000000
}

func process(registers []Register, program []int) []int {
	edge := len(program) - 1
	var instructionPointer int
	var results []int

	for instructionPointer < edge {
		switch program[instructionPointer] {
		case 0:
			registers[0].value = registers[0].value / (getCombo(program[instructionPointer+1], registers) * getCombo(program[instructionPointer+1], registers))
		case 1:
			registers[1].value ^= program[instructionPointer+1]
		case 2:
			registers[1].value = getCombo(program[instructionPointer+1], registers) % 8
		case 3:
			if registers[0].value > 0 {
				instructionPointer = program[instructionPointer+1]
				continue
			}
		case 4:
			registers[1].value ^= registers[2].value
		case 5:
			results = append(results, getCombo(program[instructionPointer+1], registers)%8)
		case 6:
			registers[1].value = registers[0].value / (getCombo(program[instructionPointer+1], registers) * getCombo(program[instructionPointer+1], registers))
		case 7:
			registers[3].value = registers[0].value / (getCombo(program[instructionPointer+1], registers) * getCombo(program[instructionPointer+1], registers))

		}

		instructionPointer += 2
		fmt.Println(instructionPointer, registers, results)
	}

	return results
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

	registers, program := readInput(file)
	fmt.Println(process(registers, program))
}
