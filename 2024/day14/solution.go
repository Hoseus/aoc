package main

import (
	"aoc2024/utils"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Robot struct {
	Position utils.Tuple[int, int]
	Velocity utils.Tuple[int, int]
}

type Quadrant struct {
	ColRange utils.Tuple[int, int]
	RowRange utils.Tuple[int, int]
	Count    int
}

func main() {
	robots := parseInput()

	part1(robots, 101, 103)
	part2(robots, 101, 103)
}

func parseInput() []Robot {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	lines := strings.Split(string(bytes), "\n")

	result := make([]Robot, 0)
	for _, line := range lines {
		robotMatch := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`).FindStringSubmatch(line)

		pX, _ := strconv.Atoi(robotMatch[1])
		pY, _ := strconv.Atoi(robotMatch[2])

		vX, _ := strconv.Atoi(robotMatch[3])
		vY, _ := strconv.Atoi(robotMatch[4])

		result = append(
			result,
			Robot{
				Position: utils.Tuple[int, int]{Left: pX, Right: pY},
				Velocity: utils.Tuple[int, int]{Left: vX, Right: vY},
			},
		)
	}

	return result
}

func part1(robots []Robot, cols, rows int) {
	movedRobots := make([]Robot, len(robots))
	copy(movedRobots, robots)

	quadrants := getQuadrants(cols, rows)
	for _, robot := range robots {
		xPosition := ((robot.Position.Left+100*robot.Velocity.Left)%cols + cols) % cols
		yPosition := ((robot.Position.Right+100*robot.Velocity.Right)%rows + rows) % rows
		for i := 0; i < len(quadrants); i++ {
			quadrant := &quadrants[i]
			if xPosition >= quadrant.ColRange.Left && xPosition <= quadrant.ColRange.Right &&
				yPosition >= quadrant.RowRange.Left && yPosition <= quadrant.RowRange.Right {
				quadrant.Count++
				break
			}
		}
	}

	result := 1
	for _, quadrant := range quadrants {
		result *= quadrant.Count
	}

	fmt.Println(fmt.Sprintf("Day 14 Part 1 Result: %v", result))
}

func part2(robots []Robot, cols, rows int) {
	movedRobots := make([]Robot, len(robots))
	copy(movedRobots, robots)

	seconds := 0
	minSafetyFactor := math.MaxInt
	for i := 0; i < 10000; i++ {
		quadrants := getQuadrants(cols, rows)
		for j, _ := range movedRobots {
			movedRobot := &movedRobots[j]
			movedRobot.Position.Left = ((movedRobot.Position.Left+movedRobot.Velocity.Left)%cols + cols) % cols
			movedRobot.Position.Right = ((movedRobot.Position.Right+movedRobot.Velocity.Right)%rows + rows) % rows
			for k := 0; k < len(quadrants); k++ {
				quadrant := &quadrants[k]
				if movedRobot.Position.Left >= quadrant.ColRange.Left && movedRobot.Position.Left <= quadrant.ColRange.Right &&
					movedRobot.Position.Right >= quadrant.RowRange.Left && movedRobot.Position.Right <= quadrant.RowRange.Right {
					quadrant.Count++
					break
				}
			}
		}
		safetyFactor := 1
		for _, quadrant := range quadrants {
			safetyFactor *= quadrant.Count
		}

		if safetyFactor < minSafetyFactor {
			seconds = i + 1
			minSafetyFactor = safetyFactor
		}
	}

	fmt.Println(fmt.Sprintf("Day 14 Part 2 Result: %v", seconds))
}

func getQuadrants(cols, rows int) []Quadrant {
	quadrants := make([]Quadrant, 4)

	colsHalf := cols / 2
	rowsHalf := rows / 2

	quadrants[0] = Quadrant{
		ColRange: utils.Tuple[int, int]{Left: 0, Right: colsHalf - 1},
		RowRange: utils.Tuple[int, int]{Left: 0, Right: rowsHalf - 1},
	}
	quadrants[1] = Quadrant{
		ColRange: utils.Tuple[int, int]{Left: cols - colsHalf, Right: cols - 1},
		RowRange: utils.Tuple[int, int]{Left: 0, Right: rowsHalf - 1},
	}
	quadrants[2] = Quadrant{
		ColRange: utils.Tuple[int, int]{Left: 0, Right: colsHalf - 1},
		RowRange: utils.Tuple[int, int]{Left: rows - rowsHalf, Right: rows - 1},
	}
	quadrants[3] = Quadrant{
		ColRange: utils.Tuple[int, int]{Left: cols - colsHalf, Right: cols - 1},
		RowRange: utils.Tuple[int, int]{Left: rows - rowsHalf, Right: rows - 1},
	}

	return quadrants
}
