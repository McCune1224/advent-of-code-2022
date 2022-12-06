package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func LoadStream(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	content := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = scanner.Text()
	}

	return content
}

func isRepeatedSequence(sequence string) bool {
	uniqueTally := map[byte]int{}
	for i := 0; i < len(sequence); i++ {
		if uniqueTally[sequence[i]] != 0 {
			return true
		} else {
			uniqueTally[sequence[i]] += 1
		}
	}
	return false
}

func UniqueSequence(stream string, sequenceLength int) int {
	tally := sequenceLength
	for i := 0; i < len(stream)-sequenceLength+1; i++ {
		foo := string(stream[i : i+sequenceLength])
		fmt.Println(foo, isRepeatedSequence(foo))
		if isRepeatedSequence(foo) == false {
			return tally
		}
		tally += 1
	}

	return -1
}

func main() {
	input := LoadStream("../input.txt")
	// part 1
	foo := UniqueSequence(input, 4)
	// part 2
	bar := UniqueSequence(input, 14)
	fmt.Println(foo, bar)
}
