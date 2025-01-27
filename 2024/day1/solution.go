package main

import (
	"aoc2024/utils"
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

func main() {
	h1, h2 := parseInput()
	part1(h1, h2)
	h1, h2 = parseInput()
	part2(h1, h2)
}

func parseInput() (utils.Heap[int], utils.Heap[int]) {
	fileName := "./input"
	file, err := os.Open(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	h1 := utils.NewHeap(func(a, b int) bool { return a < b })
	h2 := utils.NewHeap(func(a, b int) bool { return a < b })

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers, _ := utils.StringsToNumbers(strings.Fields(line))
		h1.PushT(numbers[0])
		h2.PushT(numbers[1])
	}

	return *h1, *h2
}

func part1(h1 utils.Heap[int], h2 utils.Heap[int]) {
	sum := 0
	size := h1.Len()
	for i := 0; i < size; i++ {
		number1 := heap.Pop(&h1).(int)
		number2 := heap.Pop(&h2).(int)

		diff := utils.Abs(number2 - number1)
		sum += diff
	}

	fmt.Println(fmt.Sprintf("Day 1 Part 1 Result: %d", sum))
}

func part2(h1 utils.Heap[int], h2 utils.Heap[int]) {
	sum := 0
	count := 0
	number1 := heap.Pop(&h1).(int)
	number2 := heap.Pop(&h2).(int)
	for h1.Len() > 0 || h2.Len() > 0 {
		if number1 < number2 && h1.Len() > 0 {
			number1 = heap.Pop(&h1).(int)
			count = 0
		} else if number1 > number2 && h2.Len() > 0 {
			number2 = heap.Pop(&h2).(int)
			count = 0
		} else {
			if (h1.Len() > 0 && number1 == h1.PeekT() && h2.Len() > 0 && number2 < h2.PeekT()) || h2.Len() == 0 {
				number1 = heap.Pop(&h1).(int)
			} else {
				count++
				number2 = heap.Pop(&h2).(int)
			}

			if number1 != number2 {
				score := number1 * count
				sum += score
				count = 0
			}
		}
	}

	fmt.Println(fmt.Sprintf("Day 1 Part 2 Result: %d", sum))
}
