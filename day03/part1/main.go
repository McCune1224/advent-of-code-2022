package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type Rutsack struct {
	FirstCompartment  string
	SecondCompartment string
	CommonItem        string
}

func StringCharTally(sequence string) map[string]int {
	hash := map[string]int{}

	for _, char := range sequence {
		hash[string(char)] += 1
	}

	return hash
}

func FindDuplicateChar(first, second string) string {
	firstHash := StringCharTally(first)
	secondHash := StringCharTally(second)

	for k := range firstHash {
		if secondHash[k] != 0 {
			return k
		}
	}

	return ""
}

func LoadRutsacks(filename string) []Rutsack {
	rutSacks := []Rutsack{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		itemTally := map[string]int{}
		bagContent := scanner.Text()
		firstCompartment := bagContent[:len(bagContent)/2]
		secondCompartment := bagContent[len(bagContent)/2:]
		for _, item := range bagContent {
			itemTally[string(item)] += 1
		}

		rutSacks = append(rutSacks,
			Rutsack{
				FirstCompartment:  firstCompartment,
				SecondCompartment: secondCompartment,
				CommonItem:        FindDuplicateChar(firstCompartment, secondCompartment),
			})
	}
	return rutSacks
}

func main() {
	example := LoadRutsacks("input.txt")

	tally := 0
	for _, bag := range example {
		charRune := []rune(bag.CommonItem)[0]
		if unicode.IsUpper(charRune) {
			fmt.Println("UPPER\t", bag.CommonItem, charRune-38)
			tally += int(charRune) - 38
		} else {
			fmt.Println("LOWER\t", bag.CommonItem, charRune-96)
			tally += int(charRune) - 96
		}
	}
	fmt.Println(tally)
}
