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

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func calculate(machine Machine, button int) [2]int {
	var results [2]int
	otherButton := (button + 1) % 2

	if machine.x%machine.buttons[button].x == 0 && machine.y%machine.buttons[button].y == 0 {
		results[button] = machine.x / machine.buttons[button].x
		results[otherButton] = 0
		return results
	}

	start := min(machine.x/machine.buttons[button].x, machine.y/machine.buttons[button].y)
	for start > 0 {
		if (machine.x-start*machine.buttons[button].x)%machine.buttons[otherButton].x == 0 && (machine.y-start*machine.buttons[button].y)%machine.buttons[otherButton].y == 0 {
			results[button] = start
			results[otherButton] = (machine.x - start*machine.buttons[button].x) / machine.buttons[otherButton].x
			return results
		}

		start--
	}

	return results
}

func checkMachine(machine Machine) int {
	results := calculate(machine, 0)

	return results[0]*3 + results[1]
}

func part1(machines []Machine) int {
	var result int
	for _, machine := range machines {
		result += checkMachine(machine)
	}

	return result
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
	fmt.Println("Part1:", part1(machines))
}
