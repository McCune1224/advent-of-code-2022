package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Name       string
	Value      int
	CycleCount int
}

func LoadProgramInstructions(path string) []Instruction {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	instructionLines := []Instruction{}
	for scanner.Scan() {
		parsedLine := strings.Split(scanner.Text(), " ")
		switch parsedLine[0] {
		case "noop":
			instructionLines = append(instructionLines, Instruction{Name: "noop", Value: 0, CycleCount: 1})
		case "addx":
			registryValueChange, err := strconv.Atoi(parsedLine[1])
			if err != nil {
				log.Fatal(err)
			}
			instructionLines = append(instructionLines, Instruction{Name: "addx", Value: registryValueChange, CycleCount: 2})
		}
	}
	return instructionLines
}

func CalculateSignalStrength(instructions []Instruction, initialCycleCheck, cycleCheckIncrement int) int {
	currentCycle := 0
	signalStrengthSum := 0
	xRegistrySum := 1

	for _, instruction := range instructions {
		switch instruction.Name {

		case "noop":
			for cycle := 0; cycle < instruction.CycleCount; cycle++ {
				currentCycle++
				if currentCycle%cycleCheckIncrement == initialCycleCheck {
					signalStrengthSum += xRegistrySum * currentCycle
					fmt.Println("Cycle:", currentCycle, "\tX:", xRegistrySum, "\tSum:", signalStrengthSum)
				}
			}

		case "addx":
			for cycle := 0; cycle < instruction.CycleCount; cycle++ {
				currentCycle++
				if currentCycle%cycleCheckIncrement == initialCycleCheck {
					signalStrengthSum += xRegistrySum * currentCycle
					fmt.Println("Cycle:", currentCycle, "\tX:", xRegistrySum, "\tSum:", signalStrengthSum)
				}
			}
			xRegistrySum += instruction.Value
		}
	}
	return signalStrengthSum
}

func main() {
	example := LoadProgramInstructions("../input.txt")
	foo := CalculateSignalStrength(example, 20, 40)
	fmt.Println(foo)
}
