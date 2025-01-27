package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	start, layout := parseInput()
	part1(start, &layout)
	part2(start, &layout)
}

func parseInput() (*utils.Tuple[utils.Tuple[int, int], uint8], []string) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	layout := strings.Split(string(bytes), "\n")

	var guardDirection uint8 = '^'
	var start *utils.Tuple[utils.Tuple[int, int], uint8]
	i := 0
	for i < len(layout) && start == nil {
		row := layout[i]
		j := 0
		for j < len(row) && start == nil {
			entity := row[j]
			if entity == guardDirection {
				start = &utils.Tuple[utils.Tuple[int, int], uint8]{
					utils.Tuple[int, int]{
						Left:  i,
						Right: j,
					},
					guardDirection,
				}
			}
			j++
		}
		i++
	}

	return start, layout
}

func part1(start *utils.Tuple[utils.Tuple[int, int], uint8], layout *[]string) {
	derefLayout := *layout
	rows := len(derefLayout)
	cols := len(derefLayout[0])

	i := start.Left.Left
	j := start.Left.Right
	currentGuardDirection := start.Right
	visited := make(map[utils.Tuple[int, int]]bool)

	for i >= 0 && i < rows && j >= 0 && j < cols {
		currentPosition := utils.Tuple[int, int]{i, j}
		visited[currentPosition] = true
		nextI, nextJ, nextGuardDirection := getNext(i, j, currentGuardDirection, layout)
		i = nextI
		j = nextJ
		currentGuardDirection = nextGuardDirection
	}
	fmt.Println(fmt.Sprintf("Day 6 Part 1 Result: %d", len(visited)))
}

func part2(start *utils.Tuple[utils.Tuple[int, int], uint8], layout *[]string) {
	derefLayout := *layout
	rows := len(derefLayout)
	cols := len(derefLayout[0])

	startI := start.Left.Left
	startJ := start.Left.Right
	startGuardDirection := start.Right

	i := startI
	j := startJ
	currentGuardDirection := startGuardDirection

	visited := make(map[utils.Tuple[int, int]]bool)
	addedObstructions := make(map[utils.Tuple[int, int]]bool)

	for i >= 0 && i < rows && j >= 0 && j < cols {
		currentPosition := utils.Tuple[int, int]{i, j}
		visited[currentPosition] = true

		nextI, nextJ, nextGuardDirection := getNext(i, j, currentGuardDirection, layout)
		nextPosition := utils.Tuple[int, int]{nextI, nextJ}

		if nextI >= 0 && nextI < rows && nextJ >= 0 && nextJ < cols &&
			nextGuardDirection == currentGuardDirection &&
			!visited[nextPosition] {
			k := i
			l := j
			currentSimulationGuardDirection := getNextGuardDirection(currentGuardDirection)
			simulationVisited := make(map[utils.Tuple[utils.Tuple[int, int], uint8]]bool)
			simulationVisited[utils.Tuple[utils.Tuple[int, int], uint8]{currentPosition, currentGuardDirection}] = true
			canLoopIfObstructed := false
			for k >= 0 && k < rows && l >= 0 && l < cols {
				currentSimulationPosition := utils.Tuple[utils.Tuple[int, int], uint8]{utils.Tuple[int, int]{k, l}, currentSimulationGuardDirection}
				if simulationVisited[currentSimulationPosition] {
					canLoopIfObstructed = true
					break
				}
				simulationVisited[currentSimulationPosition] = true

				nextK, nextL, nextSimulationGuardDirection := getNext(k, l, currentSimulationGuardDirection, layout)

				//The next is the simulated obstacle, so we must treat it as #
				if nextK == nextI && nextL == nextJ {
					//Stay in the same place
					nextK = k
					nextL = l
					//Change direction
					nextSimulationGuardDirection = getNextGuardDirection(currentSimulationGuardDirection)
				}

				k = nextK
				l = nextL
				currentSimulationGuardDirection = nextSimulationGuardDirection
			}

			if canLoopIfObstructed {
				addedObstructions[nextPosition] = true
			}
		}

		i = nextI
		j = nextJ
		currentGuardDirection = nextGuardDirection
	}

	fmt.Println(fmt.Sprintf("Day 6 Part 2 Result: %d", len(addedObstructions)))
}

func getNext(i, j int, currentGuardDirection uint8, layout *[]string) (int, int, uint8) {
	derefLayout := *layout
	rows := len(derefLayout)
	cols := len(derefLayout[0])

	nextI, nextJ := getNextPosition(i, j, currentGuardDirection)

	var nextValue uint8 = '.'
	if nextI >= 0 && nextI < rows && nextJ >= 0 && nextJ < cols {
		nextValue = derefLayout[nextI][nextJ]
	}

	nextGuardDirection := currentGuardDirection
	if nextValue == '#' {
		nextI = i
		nextJ = j
		nextGuardDirection = getNextGuardDirection(currentGuardDirection)
	}

	return nextI, nextJ, nextGuardDirection
}

func getNextPosition(i, j int, currentGuardDirection uint8) (int, int) {
	nextI := i
	nextJ := j
	switch currentGuardDirection {
	case '^':
		nextI = i - 1
	case '>':
		nextJ = j + 1
	case 'v':
		nextI = i + 1
	case '<':
		nextJ = j - 1
	}

	return nextI, nextJ
}

func getNextGuardDirection(currentGuardDirection uint8) uint8 {
	var nextGuardDirection uint8
	switch currentGuardDirection {
	case '^':
		nextGuardDirection = '>'
	case '>':
		nextGuardDirection = 'v'
	case 'v':
		nextGuardDirection = '<'
	case '<':
		nextGuardDirection = '^'
	}

	return nextGuardDirection
}
