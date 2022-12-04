package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	FirstSection int
	LastSection  int
	TotalTasks   int
}

func NewAssignment(writtenAssignment string) Assignment {
	workRange := strings.Split(writtenAssignment, "-")
	start, _ := strconv.Atoi(workRange[0])
	end, _ := strconv.Atoi(workRange[1])

	totalTasks := end - start + 1
	return Assignment{
		FirstSection: start,
		LastSection:  end,
		TotalTasks:   totalTasks,
	}
}

func LoadAssignmentPairs(path string) [][2]string {
	workPairs := [][2]string{}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lineContent := scanner.Text()
		elfTasks := strings.Split(lineContent, ",")
		pair := [2]string{elfTasks[0], elfTasks[1]}
		workPairs = append(workPairs, pair)
	}
	return workPairs
}

func OverlapingTasks(pair [2]string) bool {
	elf1 := NewAssignment(pair[0])
	elf2 := NewAssignment(pair[1])
	larger := Assignment{}
	smaller := Assignment{}
	if elf1.TotalTasks > elf2.TotalTasks {
		larger = elf1
		smaller = elf2
	} else {
		larger = elf2
		smaller = elf1
	}

	for i := larger.FirstSection; i < larger.LastSection+1; i++ {
		if smaller.FirstSection == i || smaller.LastSection == i {
			return true
		}
	}
	return false
}

func main() {
	example := LoadAssignmentPairs("../input.txt")
	tally := 0
	for _, pair := range example {
		if OverlapingTasks(pair) {
			tally += 1
		}
	}
	fmt.Println(tally)
}
