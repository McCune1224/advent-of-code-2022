package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ROCK        = "A"
	PAPER       = "B"
	SCISSORS    = "C"
	PLAYER_LOSE = "X"
	PLAYER_DRAW = "Y"
	PLAYER_WIN  = "Z"
	DRAW        = 0
	WIN         = 1
	LOSE        = -1
)

type TurnStrategy struct {
	OpponentOption string
	PlayerResult   string
}

func SymbolToValue(symbol string) int {
	switch symbol {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	default:
		return -1
	}
}

func RockPaperScissors(playerOption, opponentOption string) int {
	if playerOption == ROCK {
		switch opponentOption {
		case ROCK:
			return DRAW
		case PAPER:
			return LOSE
		case SCISSORS:
			return WIN
		}
	}
	if playerOption == PAPER {
		switch opponentOption {
		case ROCK:
			return WIN
		case PAPER:
			return DRAW
		case SCISSORS:
			return LOSE
		}
	}
	if playerOption == SCISSORS {
		switch opponentOption {
		case ROCK:
			return LOSE
		case PAPER:
			return WIN
		case SCISSORS:
			return DRAW
		}
	}
	return -2
}

func FileToStrategySlice(filePath string) []TurnStrategy {
	fileSlice := []TurnStrategy{}
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		turn := strings.Split(fileScanner.Text(), " ")
		fileSlice = append(fileSlice,
			TurnStrategy{
				OpponentOption: turn[0],
				PlayerResult:   turn[1],
			})
	}
	return fileSlice
}

func SelectPlayerOption(playerDesiredOutcome string, opponentOption string) string {
	switch playerDesiredOutcome {
	case PLAYER_WIN:
		if opponentOption == ROCK {
			return PAPER
		}
		if opponentOption == PAPER {
			return SCISSORS
		}
		if opponentOption == SCISSORS {
			return ROCK
		}
	case PLAYER_LOSE:
		if opponentOption == ROCK {
			return SCISSORS
		}
		if opponentOption == PAPER {
			return ROCK
		}
		if opponentOption == SCISSORS {
			return PAPER
		}
	case PLAYER_DRAW:
		return opponentOption
	}
	return ""
}

func SumRPSResults(games []TurnStrategy) int {
	tally := 0
	for _, game := range games {

		playerOption := SelectPlayerOption(game.PlayerResult, game.OpponentOption)

		switch RockPaperScissors(playerOption, game.OpponentOption) {
		case DRAW:
			tally += 3 + SymbolToValue(playerOption)
		case WIN:
			tally += 6 + SymbolToValue(playerOption)
		case LOSE:
			tally += 0 + SymbolToValue(playerOption)
		}
	}
	return tally
}

func main() {
	rpsStrategy := FileToStrategySlice("input.txt")
	totalGamesPoints := SumRPSResults(rpsStrategy)
	fmt.Println(totalGamesPoints)
}
