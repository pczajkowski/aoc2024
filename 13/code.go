package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Button struct {
	id   byte
	x, y int
}

type Machine struct {
	buttons []Button
	x, y    int
}

func readInput(file *os.File) []Machine {
	scanner := bufio.NewScanner(file)
	var machines []Machine
	var machine Machine

	var buttonsRead int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			buttonsRead = 0
			machines = append(machines, machine)
			machine = Machine{}
			continue
		}

		if buttonsRead < 2 {
			var button Button
			n, err := fmt.Sscanf(line, "Button %c: X+%d, Y+%d", &button.id, &button.x, &button.y)
			if n != 3 || err != nil {
				log.Fatalf("Not able to parse button '%s': %s", line, err)
			}

			machine.buttons = append(machine.buttons, button)
			buttonsRead++
		} else {
			n, err := fmt.Sscanf(line, "Prize: X=%d, Y=%d", &machine.x, &machine.y)
			if n != 2 || err != nil {
				log.Fatalf("Not able to parse machine '%s': %s", line, err)
			}
		}
	}

	machines = append(machines, machine)
	return machines
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

	machines := readInput(file)
	fmt.Println(machines)
}
