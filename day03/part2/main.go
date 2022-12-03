package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func FindDuplicateChar(sequences ...string) string {
	seqHashes := []map[string]int{}
	for _, seq := range sequences {
		seqHashes = append(seqHashes, StringCharTally(seq))
	}

	for k := range seqHashes[0] {
		if seqHashes[1][k] != 0 && seqHashes[2][k] != 0 {
			return k
		}
	}

	return ""
}

func StringCharTally(sequence string) map[string]int {
	hash := map[string]int{}

	for _, char := range sequence {
		hash[string(char)] += 1
	}

	return hash
}

func LoadRutsacks(filename string) []string {
	commonItemsList := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	currentGroup := []string{}
	lineCount := 0
	for scanner.Scan() {
		lineCount += 1
		currentGroup = append(currentGroup, scanner.Text())
		if lineCount%3 == 0 {
			dup := FindDuplicateChar(currentGroup...)
			commonItemsList = append(commonItemsList, dup)
			currentGroup = []string{}
		}
	}
	return commonItemsList
}

func GetAsciiCharacterSum(asciiSequence []string) int {
	tally := 0
	for _, char := range asciiSequence {
		asciiTableVal := []rune(char)[0]
		if unicode.IsUpper(asciiTableVal) {
			fmt.Println("UPPER\t", char, asciiTableVal-38)
			tally += int(asciiTableVal - 38)
		} else {
			fmt.Println("LOWER\t", char, asciiTableVal-96)
			tally += int(asciiTableVal) - 96
		}
	}
	return tally
}

func main() {
	example := LoadRutsacks("../example.txt")
	fmt.Println(example)
	tally := GetAsciiCharacterSum(example)

	fmt.Println(tally)
}
