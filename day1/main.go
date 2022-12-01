package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func FileToSlice(filePath string) []string {
	fileSlice := []string{}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileSlice = append(fileSlice, fileScanner.Text())
	}
	return fileSlice
}

func PartOne(caloriesList []string) int {
	maxCalories := 0
	currentCalories := 0

	for _, calorie := range caloriesList {
		if calorie == "" {
			if maxCalories < currentCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
			continue
		}

		intCalorie, err := strconv.Atoi(calorie)
		if err != nil {
			log.Fatal(err)
		}

		currentCalories += intCalorie
	}

	return maxCalories
}

func PartTwo(caloriesList []string) int {
	caloriesSum := 0
	elfCaloriesList := []int{}
	currentCaloriesTally := 0

	for _, calorie := range caloriesList {
		if calorie == "" {
			elfCaloriesList = append(elfCaloriesList, currentCaloriesTally)
			currentCaloriesTally = 0
			continue
		}
		intCalorie, err := strconv.Atoi(calorie)
		if err != nil {
			log.Fatal(err)
		}
		currentCaloriesTally += intCalorie
	}

	sort.Slice(elfCaloriesList, func(i, j int) bool {
		return elfCaloriesList[i] < elfCaloriesList[j]
	})

	for _, calorie := range elfCaloriesList[len(elfCaloriesList)-3:] {
		caloriesSum += calorie
	}

	return caloriesSum
}

func main() {
	fileSlice := FileToSlice("input.txt")
	answer1 := PartOne(fileSlice)
	answer2 := PartTwo(fileSlice)
	fmt.Printf("Part 1:\t%v\n", answer1)
	fmt.Printf("Part 2:\t%v\n", answer2)
}
