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

func CalculateSignalStrength(instructions []Instruction, initialCycleCheck, cycleCheckIncrement int) (int, []string) {
	currentCycle := 0
	signalStrengthSum := 0
	xRegistrySum := 1
	crt := []string{}
	for _, instruction := range instructions {
		switch instruction.Name {

		case "noop":
			for cycle := 0; cycle < instruction.CycleCount; cycle++ {
				currentCycle++
				if currentCycle%cycleCheckIncrement == initialCycleCheck {
					signalStrengthSum += xRegistrySum * currentCycle
					fmt.Println("Cycle:", currentCycle, "\tX:", xRegistrySum, "\tSum:", signalStrengthSum)
				}
				crt = append(crt, AddCrtPixel(currentCycle, xRegistrySum))
			}

		case "addx":
			for cycle := 0; cycle < instruction.CycleCount; cycle++ {
				currentCycle++
				if currentCycle%cycleCheckIncrement == initialCycleCheck {
					signalStrengthSum += xRegistrySum * currentCycle
					fmt.Println("Cycle:", currentCycle, "\tX:", xRegistrySum, "\tSum:", signalStrengthSum)
				}
				crt = append(crt, AddCrtPixel(currentCycle, xRegistrySum))
			}
			xRegistrySum += instruction.Value
		}
	}
	return signalStrengthSum, crt
}

func AddCrtPixel(cycle, register int) string {
	position := cycle % 40
	// top line edge case
	if position == 0 {
		position = 40
	}
	// Sprite is 3 wide ###
	lowEnd := register
	highEnd := register + 2
	if position >= lowEnd && position <= highEnd {
		return "#"
	}
	return " "
}

func main() {
	example := LoadProgramInstructions("../input.txt")
	_, crt := CalculateSignalStrength(example, 20, 40)
	for index, val := range crt {
		fmt.Print(val)
		if ((index + 1) % 40) == 0 {
			fmt.Println("")
		}
	}
}
