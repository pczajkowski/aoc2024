package main

import (
	"fmt"
	"log"
	"os"
)

func getDisk(diskMap []byte) (int, []int) {
	var disk []int
	file := true
	var fileID int

	for _, block := range diskMap {
		if block < 48 || block > 57 {
			continue
		}

		number := int(block) - 48
		if file {
			file = false
			for j := 0; j < number; j++ {
				disk = append(disk, fileID)
			}

			fileID++
		} else {
			file = true
			for j := 0; j < number; j++ {
				disk = append(disk, -1)
			}
		}
	}

	return int(diskMap[0]) - 48, disk
}

func compact(disk []int, free int) []int {
	end := len(disk) - 1
	for free < end {
		if disk[free] != -1 {
			free++
			continue
		}

		if disk[end] == -1 {
			end--
			continue
		}

		disk[free], disk[end] = disk[end], disk[free]
		free++
		end--
	}

	return disk
}

func part1(diskMap []byte) int {
	free, disk := getDisk(diskMap)
	compacted := compact(disk, free)
	fmt.Println(compacted)

	return 0
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	diskMap, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1(diskMap))
}
