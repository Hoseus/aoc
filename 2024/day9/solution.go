package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"slices"
)

func main() {
	memory, memoryAllocations := parseInput()

	part1(memory)
	part2(memory, memoryAllocations)
}

func parseInput() ([]int, [][]int) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	memory := make([]int, 0)
	memoryAllocations := make([][]int, 0)
	for i, memoryBlock := range string(bytes) {
		size := int(memoryBlock - '0')
		memory = append(memory, size)
		if utils.Even(i) {
			memoryAllocations = append(memoryAllocations, slices.Repeat([]int{i / 2}, size))
		} else {
			memoryAllocations = append(memoryAllocations, make([]int, 0))
		}
	}

	return memory, memoryAllocations
}

func part1(memory []int) {
	result := 0
	defragmentedMemory := make([]int, 0)
	l, r := 0, len(memory)-1
	leftFreeMemory, rightRemainingBlocks := -1, -1
	for l < r {
		fileId, fileSize, blocksToWrite := 0, 0, 0
		if utils.Even(l) {
			fileId = l / 2
			fileSize = memory[l]
			blocksToWrite = fileSize
			l++
		} else if utils.Even(r) {
			fileId = r / 2
			fileSize = memory[r]

			if leftFreeMemory < 0 {
				leftFreeMemory = memory[l]
			}

			if rightRemainingBlocks < 0 {
				rightRemainingBlocks = fileSize
			}

			if leftFreeMemory >= rightRemainingBlocks {
				blocksToWrite = rightRemainingBlocks
				leftFreeMemory = leftFreeMemory - blocksToWrite
				rightRemainingBlocks = -1
				r--
			} else {
				blocksToWrite = leftFreeMemory
				rightRemainingBlocks = rightRemainingBlocks - blocksToWrite
				leftFreeMemory = -1
				l++
			}
		} else {
			r--
		}

		if l == r && rightRemainingBlocks > 0 {
			blocksToWrite += rightRemainingBlocks
		}
		defragmentedMemory = append(defragmentedMemory, slices.Repeat([]int{fileId}, blocksToWrite)...)
	}

	for i, fileId := range defragmentedMemory {
		result += i * fileId
	}

	fmt.Println(fmt.Sprintf("Day 9 Part 1 Result: %d", result))
}

func part2(memory []int, memoryAllocations [][]int) {
	l, r := 1, len(memory)-1
	for r >= 0 {
		if utils.Even(r) && r > l {
			fileId := r / 2
			fileSize := memory[r]

			i := l
			fileWasMoved := false
			for i < r && !fileWasMoved {
				memoryAllocation := memoryAllocations[i]
				freeMemorySize := memory[i] - len(memoryAllocation)
				if freeMemorySize >= fileSize {
					memoryAllocations[i] = append(memoryAllocation, slices.Repeat([]int{fileId}, fileSize)...)
					memoryAllocations[r] = slices.Repeat([]int{0}, fileSize)
					fileWasMoved = true
				}
				i += 2
			}
		} else {
			memoryAllocation := memoryAllocations[r]
			memoryAllocationSize := len(memoryAllocation)
			originalFreeMemorySize := memory[r]

			memoryAllocations[r] = append(memoryAllocation, slices.Repeat([]int{0}, originalFreeMemorySize-memoryAllocationSize)...)
		}
		r--
	}

	result := 0
	j := 0
	for _, memoryAllocation := range memoryAllocations {
		for _, fileId := range memoryAllocation {
			result += j * fileId
			j++
		}
	}

	fmt.Println(fmt.Sprintf("Day 9 Part 2 Result: %d", result))
}
