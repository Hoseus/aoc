package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var adjPossibilities = []utils.Tuple[int, int]{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func main() {
	topographicMap, trailHeads := parseInput()

	part1(&topographicMap, trailHeads)
	part2(&topographicMap, trailHeads)
}

func parseInput() ([][]int, []utils.Tuple[int, int]) {
	fileName := "./input"
	bytes, _ := os.ReadFile(fileName) // Read the entire file
	splitInput := strings.Split(string(bytes), "\n")

	topographicMap := make([][]int, 0)
	trailHeads := make([]utils.Tuple[int, int], 0)
	for i, row := range splitInput {
		topographicMap = append(topographicMap, make([]int, 0))
		for j, cell := range row {
			height, _ := strconv.Atoi(string(cell))
			topographicMap[i] = append(topographicMap[i], height)
			if height == 0 {
				trailHeads = append(trailHeads, utils.Tuple[int, int]{i, j})
			}
		}
	}

	return topographicMap, trailHeads
}

func part1(topographicMap *[][]int, trailHeads []utils.Tuple[int, int]) {
	count := 0
	for _, trailHead := range trailHeads {
		count += countHikingTrailsPart1(trailHead, topographicMap)
	}
	fmt.Println(fmt.Sprintf("Day 10 Part 1 Result: %d", count))
}

func part2(topographicMap *[][]int, trailHeads []utils.Tuple[int, int]) {
	count := 0
	for _, trailHead := range trailHeads {
		count += countHikingTrailsPart2(trailHead, topographicMap)
	}
	fmt.Println(fmt.Sprintf("Day 10 Part 2 Result: %d", count))
}

func countHikingTrailsPart1(start utils.Tuple[int, int], topographicMap *[][]int) int {
	count := 0

	derefTopographicMap := *topographicMap
	if derefTopographicMap[start.Left][start.Right] != 0 {
		return count
	}

	q := utils.NewDeque[utils.Tuple[int, int]]()
	q.PushBack(start)
	visited := make(map[utils.Tuple[int, int]]bool)
	visited[start] = true

	for q.Len() > 0 {
		curr, _ := q.PopFront()

		row := curr.Left
		col := curr.Right
		height := derefTopographicMap[row][col]

		if height == 9 {
			count++
		}

		for _, adj := range getAdj(curr, topographicMap) {
			if !visited[adj] {
				visited[adj] = true
				q.PushBack(adj)
			}
		}
	}

	return count
}

func countHikingTrailsPart2(start utils.Tuple[int, int], topographicMap *[][]int) int {
	count := 0

	derefTopographicMap := *topographicMap
	if derefTopographicMap[start.Left][start.Right] != 0 {
		return count
	}

	q := utils.NewDeque[utils.Tuple[int, int]]()
	q.PushBack(start)

	for q.Len() > 0 {
		curr, _ := q.PopFront()

		row := curr.Left
		col := curr.Right
		height := derefTopographicMap[row][col]

		if height == 9 {
			count++
		}

		for _, adj := range getAdj(curr, topographicMap) {
			q.PushBack(adj)
		}
	}

	return count
}

func getAdj(point utils.Tuple[int, int], topographicMap *[][]int) []utils.Tuple[int, int] {
	derefTopographicMap := *topographicMap

	rows := len(derefTopographicMap)
	cols := len(derefTopographicMap[0])

	row := point.Left
	col := point.Right
	height := derefTopographicMap[row][col]

	adj := make([]utils.Tuple[int, int], 0)
	for _, adjPossibility := range adjPossibilities {
		adjRow := row + adjPossibility.Left
		adjCol := col + adjPossibility.Right
		if adjRow >= 0 && adjRow < rows && adjCol >= 0 && adjCol < cols &&
			derefTopographicMap[adjRow][adjCol]-height == 1 {
			adj = append(adj, utils.Tuple[int, int]{adjRow, adjCol})
		}
	}
	return adj
}
