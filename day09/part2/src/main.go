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
	Direction string
	Distance  int
}

func LoadInstructions(path string) []Instruction {
	instructionList := []Instruction{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parsedLine := strings.Split(scanner.Text(), " ")
		direction := parsedLine[0]
		distance, err := strconv.Atoi(parsedLine[1])
		if err != nil {
			log.Fatal(err)
		}
		instructionList = append(instructionList, Instruction{direction, distance})

	}

	return instructionList
}

func main() {
	instructions := LoadInstructions("../../example.txt")
	for _, instruction := range instructions {
		fmt.Println(instruction)
	}
}
