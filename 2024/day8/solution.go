package main

import (
	"aoc2024/utils"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	grid, antennaLocations := parseInput()

	part1(&grid, &antennaLocations)
	part2(&grid, &antennaLocations)
}

func parseInput() ([]string, map[rune][]utils.Tuple[int, int]) {
	fileName := "./input"
	bytes, _ := os.ReadFile(fileName) // Read the entire file
	grid := strings.Split(string(bytes), "\n")

	antennaLocations := make(map[rune][]utils.Tuple[int, int])
	for i, row := range grid {
		for j, cell := range row {
			if unicode.IsDigit(cell) || unicode.IsLetter(cell) {
				antennaLocations[cell] = append(antennaLocations[cell], utils.Tuple[int, int]{i, j})
			}
		}
	}

	return grid, antennaLocations
}

func part1(grid *[]string, antennaLocations *map[rune][]utils.Tuple[int, int]) {
	derefGrid := *grid
	rows := len(derefGrid)
	cols := len(derefGrid[0])

	antinodeLocations := make(map[utils.Tuple[int, int]]bool)
	for _, locations := range *antennaLocations {
		l, r := 0, len(locations)-1
		for l < len(locations)-1 {
			if l == r {
				l++
				r = len(locations) - 1
			} else {
				location1 := locations[l]
				location2 := locations[r]
				antinode1, antinode2 := getAntinodesPart1(location1, location2, rows, cols)
				if antinode1 != nil {
					antinodeLocations[*antinode1] = true
				}

				if antinode2 != nil {
					antinodeLocations[*antinode2] = true
				}
				r--
			}
		}
	}
	fmt.Println(fmt.Sprintf("Day 8 Part 1 Result: %d", len(antinodeLocations)))
}

func part2(grid *[]string, antennaLocations *map[rune][]utils.Tuple[int, int]) {
	derefGrid := *grid
	rows := len(derefGrid)
	cols := len(derefGrid[0])

	antinodeLocations := make(map[utils.Tuple[int, int]]bool)
	for _, locations := range *antennaLocations {
		l, r := 0, len(locations)-1
		for l < len(locations)-1 {
			if l == r {
				l++
				r = len(locations) - 1
			} else {
				location1 := locations[l]
				location2 := locations[r]
				antinodes := getAntinodesPart2(location1, location2, rows, cols)
				for _, antinode := range antinodes {
					antinodeLocations[antinode] = true
				}
				r--
			}
		}
	}
	fmt.Println(fmt.Sprintf("Day 8 Part 2 Result: %d", len(antinodeLocations)))
}

func isInBounds(x, y int, m, n int) bool {
	return x >= 0 && x < m && y >= 0 && y < n
}

func getAntinodesPart1(location1, location2 utils.Tuple[int, int], m, n int) (*utils.Tuple[int, int], *utils.Tuple[int, int]) {
	x1 := location1.Left
	y1 := location1.Right

	x2 := location2.Left
	y2 := location2.Right

	diffx := x2 - x1
	diffY := y2 - y1

	antinode1X := x1 - diffx
	antinode1Y := y1 - diffY

	antinode2X := x2 + diffx
	antinode2Y := y2 + diffY

	var antinode1 *utils.Tuple[int, int] = nil
	if isInBounds(antinode1X, antinode1Y, m, n) {
		antinode1 = &utils.Tuple[int, int]{antinode1X, antinode1Y}
	}

	var antinode2 *utils.Tuple[int, int] = nil
	if isInBounds(antinode2X, antinode2Y, m, n) {
		antinode2 = &utils.Tuple[int, int]{antinode2X, antinode2Y}
	}

	return antinode1, antinode2
}

func getAntinodesPart2(location1, location2 utils.Tuple[int, int], m, n int) []utils.Tuple[int, int] {
	x1 := location1.Left
	y1 := location1.Right

	x2 := location2.Left
	y2 := location2.Right

	diffx := x2 - x1
	diffY := y2 - y1

	antinodes := make([]utils.Tuple[int, int], 0)

	antinode1X := x1
	antinode1Y := y1
	for isInBounds(antinode1X, antinode1Y, m, n) {
		antinodes = append(antinodes, utils.Tuple[int, int]{antinode1X, antinode1Y})
		antinode1X = antinode1X - diffx
		antinode1Y = antinode1Y - diffY
	}

	antinode2X := x2
	antinode2Y := y2
	for isInBounds(antinode2X, antinode2Y, m, n) {
		antinodes = append(antinodes, utils.Tuple[int, int]{antinode2X, antinode2Y})
		antinode2X = antinode2X + diffx
		antinode2Y = antinode2Y + diffY
	}

	return antinodes
}
