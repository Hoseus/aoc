package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Registers struct {
	A int
	B int
	C int
}

type Context struct {
	InstructionPointer int
	Outputs            []int
}

func main() {
	registers, program := parseInput()
	part1(registers, program)
	part2(registers, program)
}

func parseInput() (Registers, []int) {
	fileName := "./input"
	bytes, err := os.ReadFile(fileName) // Read the entire file
	if err != nil {
		panic(fmt.Sprintf("Failure reading file: %s", fileName))
	}
	programParts := strings.Split(string(bytes), "\n\n")

	registerLines := strings.Split(programParts[0], "\n")
	registerRegex := regexp.MustCompile(`\d+`)
	registerA, _ := strconv.Atoi(registerRegex.FindString(registerLines[0]))
	registerB, _ := strconv.Atoi(registerRegex.FindString(registerLines[1]))
	registerC, _ := strconv.Atoi(registerRegex.FindString(registerLines[2]))
	registers := Registers{
		A: registerA,
		B: registerB,
		C: registerC,
	}

	programRegex := regexp.MustCompile(`\d+(?:,\d+)*`)
	program := make([]int, 0)
	for _, instructionMemberString := range strings.Split(programRegex.FindString(programParts[1]), ",") {
		instructionMember, _ := strconv.Atoi(instructionMemberString)
		program = append(program, instructionMember)
	}

	return registers, program
}

func part1(registers Registers, program []int) {
	outputs := runProgram(registers, program)
	result := join(outputs, ",")
	fmt.Println(fmt.Sprintf("Day 17 Part 1 Result: %s", result))
}

func part2(_ Registers, program []int) {
	fmt.Println(fmt.Sprintf("Day 17 Part 2 Result: %d", reverseEngineer(program)))
}

func runProgram(registers Registers, program []int) []int {
	context := Context{}
	for context.InstructionPointer < len(program) {
		oldInstructionPointer := context.InstructionPointer
		call(program[context.InstructionPointer], &registers, program[context.InstructionPointer+1], &context)
		if context.InstructionPointer == oldInstructionPointer {
			context.InstructionPointer += 2
		}
	}

	return context.Outputs
}

func call(opCode int, registers *Registers, operand int, context *Context) {
	instruction := instructions[opCode]
	instruction(registers, operand, context)
}

var instructions = []func(*Registers, int, *Context){
	adv,
	bxl,
	bst,
	jnz,
	bxc,
	out,
	bdv,
	cdv,
}

func adv(registers *Registers, operand int, _ *Context) {
	comboOperand := solveComboOperand(registers, operand)
	registers.A = dv(registers.A, comboOperand)
}

func bdv(registers *Registers, operand int, _ *Context) {
	comboOperand := solveComboOperand(registers, operand)
	registers.B = dv(registers.A, comboOperand)
}

func cdv(registers *Registers, operand int, _ *Context) {
	comboOperand := solveComboOperand(registers, operand)
	registers.C = dv(registers.A, comboOperand)
}

func dv(a, b int) int {
	return int(float64(a) / (math.Pow(2.0, float64(b))))
}

func bxl(registers *Registers, operand int, _ *Context) {
	registers.B = registers.B ^ operand
}

func bxc(registers *Registers, _ int, _ *Context) {
	registers.B = registers.B ^ registers.C
}

func bst(registers *Registers, operand int, _ *Context) {
	comboOperand := solveComboOperand(registers, operand)
	registers.B = mod8(comboOperand)
}

func out(registers *Registers, operand int, context *Context) {
	comboOperand := solveComboOperand(registers, operand)
	context.Outputs = append(context.Outputs, mod8(comboOperand))
}

func mod8(a int) int {
	modulo := 8
	return (a%modulo + modulo) % modulo
}

func jnz(registers *Registers, operand int, context *Context) {
	if registers.A != 0 {
		context.InstructionPointer = operand
	}
}

func solveComboOperand(registers *Registers, operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	} else if operand == 4 {
		return registers.A
	} else if operand == 5 {
		return registers.B
	} else if operand == 6 {
		return registers.C
	}
	return -1
}

func join(arr []int, separator string) string {
	stringOutputs := make([]string, 0)
	for _, output := range arr {
		stringOutputs = append(stringOutputs, strconv.Itoa(output))
	}
	return strings.Join(stringOutputs, separator)
}

func reverseEngineer(program []int) int {
	validValsForA := map[int]bool{0: true}
	for i := len(program) - 1; i >= 0; i-- {
		instructionPointer := program[i]
		nextValsForA := make(map[int]bool)
		for aVal := range validValsForA {
			aShifted := aVal * 8
			for candidateA := aShifted; candidateA < aShifted+8; candidateA++ {
				output := runProgram(Registers{A: candidateA}, program)
				if len(output) > 0 && output[0] == instructionPointer {
					nextValsForA[candidateA] = true
				}
			}
		}

		validValsForA = nextValsForA
	}

	validValsForASlice := make([]int, 0)
	for validValForA := range validValsForA {
		validValsForASlice = append(validValsForASlice, validValForA)
	}

	return slices.Min(validValsForASlice)
}
