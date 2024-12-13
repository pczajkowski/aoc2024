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

func calculate(machine Machine, button int, limit int) [2]int {
	var results [2]int
	otherButton := (button + 1) % 2

	if machine.x%machine.buttons[button].x == 0 && machine.y%machine.buttons[button].y == 0 {
		pushes := machine.x / machine.buttons[button].x
		if pushes*machine.buttons[button].y == machine.y {
			results[button] = pushes
			results[otherButton] = 0
			return results
		}
	}

	start := min(machine.x/machine.buttons[button].x, machine.y/machine.buttons[button].y)
	if limit > 0 && start > limit {
		start = limit
	}

	for ; start > 0; start-- {
		deltaX := machine.x - start*machine.buttons[button].x
		if deltaX%machine.buttons[otherButton].x == 0 {
			otherPushes := deltaX / machine.buttons[otherButton].x
			if limit > 0 && otherPushes > limit {
				continue
			}

			if machine.y-start*machine.buttons[button].y != otherPushes*machine.buttons[otherButton].y {
				continue
			}

			results[button] = start
			results[otherButton] = otherPushes
			return results
		}
	}

	return results
}

func checkMachine(machine Machine, limit int) int {
	resultA := calculate(machine, 0, limit)
	resultB := calculate(machine, 1, limit)

	costA := resultA[0]*3 + resultA[1]
	costB := resultB[0]*3 + resultB[1]

	return min(costA, costB)
}

func solve(machines []Machine, limit int) int {
	var result int
	for _, machine := range machines {
		result += checkMachine(machine, limit)
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
	fmt.Println("Part1:", solve(machines, 100))
}
