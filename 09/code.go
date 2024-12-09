package main

import (
	"fmt"
	"log"
	"os"
)

func getDiskMap(data []byte) []int {
	diskMap := make([]int, len(data))
	for i, block := range data {
		if block < 48 || block > 57 {
			continue
		}

		diskMap[i] = int(block) - 48
	}

	return diskMap
}

func getDisk(diskMap []int) (int, []int) {
	var disk []int
	file := true
	var fileID int

	for _, number := range diskMap {
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

	return diskMap[0], disk
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

func calculateChecksum(disk []int) int64 {
	var checksum int64

	for i := range disk {
		if disk[i] != -1 {
			checksum += int64(disk[i] * i)
		}
	}

	return checksum
}

func part1(diskMap []int) int64 {
	free, disk := getDisk(diskMap)
	compacted := compact(disk, free)

	return calculateChecksum(compacted)
}

func findFile(disk []int, id, end int) int {
	for i := end; i > 0; i-- {
		if disk[i] != id {
			return i + 1
		}
	}

	return -1
}

func findFree(disk []int, start, end, goal int) (int, int) {
	var freeStart, freeEnd int
	for i := start; i < end; i++ {
		if disk[i] == -1 {
			if freeStart <= 0 {
				freeStart = i
				for j := freeStart + 1; j < freeStart+goal; j++ {
					if disk[j] != -1 {
						freeStart = -1
						i = j
						break
					}
				}

				if freeStart > 0 {
					return freeStart, freeStart + goal
				}
			}
		}
	}

	return freeStart, freeEnd
}

func defrag(diskMap []int) []int {
	free, disk := getDisk(diskMap)
	end := len(disk) - 1

	for i := end; i > free; i-- {
		if disk[i] > 0 {
			fileStart := findFile(disk, disk[i], i)
			if fileStart > 0 {
				freeStart, freeEnd := findFree(disk, free, fileStart, i-fileStart+1)
				if freeStart > 0 && freeEnd > 0 && freeStart < freeEnd {
					for k := freeStart; k < freeEnd; k++ {
						disk[k] = disk[i]
					}

					for k := fileStart; k <= i; k++ {
						disk[k] = -1
					}
				}

				i = fileStart
			}
		}
	}

	return disk
}

func part2(diskMap []int) int64 {
	defragmented := defrag(diskMap)

	return calculateChecksum(defragmented)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to specify a file!")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	diskMap := getDiskMap(data)
	fmt.Println("Part1:", part1(diskMap))
	fmt.Println("Part2:", part2(diskMap))
}
