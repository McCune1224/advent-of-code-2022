package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func LoadCrateContent(filePath string) ([]string, []string) {
	stacks := make([]string, 0)
	instructions := make([]string, 0)
	parsedStacks := false

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, " 1") {
			continue
		}

		if line == "" {
			parsedStacks = true
			continue
		}

		if parsedStacks {
			instructions = append(instructions, line)
		} else {
			stacks = append(stacks, line)
		}
	}

	return stacks, instructions
}

func InitStacks(lines []string) [][]string {
	stacks := make([][]string, 0)

	n := strings.Count(lines[len(lines)-1], "[")
	for i := 0; i < n; i++ {
		stacks = append(stacks, make([]string, 0))
	}

	for i := len(lines) - 1; i >= 0; i-- {
		for j, c := range lines[i] {
			if unicode.IsLetter(c) {
				fmt.Println(c)
				stacks[j/4] = append(stacks[j/4], string(c))
			}
		}
	}

	return stacks
}

func InitInstructions(lines []string) [][]string {
	instructions := make([][]string, 0)

	for _, line := range lines {
		re := regexp.MustCompile("[0-9]+")
		nums := re.FindAllString(line, -1)
		instructions = append(instructions, nums)
	}

	return instructions
}

func PerformInstructions(stacks [][]string, instructions [][]string) [][]string {
	for _, instruction := range instructions {
		n, _ := strconv.Atoi(instruction[0])
		from, _ := strconv.Atoi(instruction[1])
		to, _ := strconv.Atoi(instruction[2])

		for i := 0; i < n; i++ {
			toMove := stacks[from-1][len(stacks[from-1])-1]
			stacks[to-1] = append(stacks[to-1], toMove)
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
		}
	}

	return stacks
}

func GetTopOfStacks(stacks [][]string) string {
	top := ""
	for _, s := range stacks {
		top += string(s[len(s)-1])
	}
	return top
}

func main() {
	stackLines, instructionLines := LoadCrateContent("../input.txt")
	stacks := InitStacks(stackLines)
	instructions := InitInstructions(instructionLines)
	stacks = PerformInstructions(stacks, instructions)

	fmt.Println(GetTopOfStacks(stacks))
}
