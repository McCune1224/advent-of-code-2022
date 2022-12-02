package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	OPPONENT_ROCK     = "A"
	OPPONENT_PAPER    = "B"
	OPPONENT_SCISSORS = "C"
	PLAYER_ROCK       = "X"
	PLAYER_PAPER      = "Y"
	PLAYER_SCISSORS   = "Z"
	DRAW              = 0
	WIN               = 1
	LOSE              = -1
)

type TurnStrategy struct {
	OpponentOption string
	PlayerOption   string
}

func SymbolToValue(symbol string) int {
	switch symbol {
	case OPPONENT_ROCK, PLAYER_ROCK:
		return 1
	case OPPONENT_PAPER, PLAYER_PAPER:
		return 2
	case OPPONENT_SCISSORS, PLAYER_SCISSORS:
		return 3
	default:
		return -1
	}
}

func RockPaperScissors(playerOption, opponentOption string) int {
	if playerOption == PLAYER_ROCK {
		switch opponentOption {
		case OPPONENT_ROCK:
			return DRAW
		case OPPONENT_PAPER:
			return LOSE
		case OPPONENT_SCISSORS:
			return WIN
		}
	}
	if playerOption == PLAYER_PAPER {
		switch opponentOption {
		case OPPONENT_ROCK:
			return WIN
		case OPPONENT_PAPER:
			return DRAW
		case OPPONENT_SCISSORS:
			return LOSE
		}
	}
	if playerOption == PLAYER_SCISSORS {
		switch opponentOption {
		case OPPONENT_ROCK:
			return LOSE
		case OPPONENT_PAPER:
			return WIN
		case OPPONENT_SCISSORS:
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
				PlayerOption:   turn[1],
			})
	}
	return fileSlice
}

func SumRPSResults(games []TurnStrategy) int {
	tally := 0
	for _, game := range games {
		switch RockPaperScissors(game.PlayerOption, game.OpponentOption) {
		case DRAW:
			tally += 3 + SymbolToValue(game.PlayerOption)
		case WIN:
			tally += 6 + SymbolToValue(game.PlayerOption)
		case LOSE:
			tally += 0 + SymbolToValue(game.PlayerOption)
		}
	}
	return tally
}

func main() {
	rpsStrategy := FileToStrategySlice("input.txt")
	totalGamesPoints := SumRPSResults(rpsStrategy)
	fmt.Println(totalGamesPoints)
}
