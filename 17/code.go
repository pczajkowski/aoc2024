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
	fmt.Println(registers, program)
}
