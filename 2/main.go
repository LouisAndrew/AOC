package day_2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type balls struct {
	red   int
	blue  int
	green int
}

type game struct {
	id    int
	balls balls
}

func getMaxNumberOfBall(substring string, currentAmount int, matcher *regexp.Regexp) int {
	matched := matcher.FindStringSubmatch(substring)
	if len(matched) <= 1 {
		return currentAmount
	}

	amount, err := strconv.Atoi(matched[1])
	if err != nil {
		panic(err)
	}

	if amount > currentAmount {
		return amount
	}

	return currentAmount
}

func parseGameLine(line string) game {
	id := 0
	balls := balls{
		red:   0,
		blue:  0,
		green: 0,
	}

	splittedByGameIdentifier := strings.Split(line, ":")

	if len(splittedByGameIdentifier) <= 1 {
		return game{}
	}

	regex, err := regexp.Compile(`Game (\d+)`)
	matched := regex.FindStringSubmatch(splittedByGameIdentifier[0])
	if err != nil {
		panic(err)
	}

	id, err = strconv.Atoi(matched[1])
	if err != nil {
		panic(err)
	}

	redRegex, _ := regexp.Compile(`(\d+) red`)
	blueRegex, _ := regexp.Compile(`(\d+) blue`)
	greenRegex, _ := regexp.Compile(`(\d+) green`)

	splittedBySemicolon := strings.Split(splittedByGameIdentifier[1], ";")
	for _, summaryDescription := range splittedBySemicolon {
		balls.red = getMaxNumberOfBall(summaryDescription, balls.red, redRegex)
		balls.green = getMaxNumberOfBall(summaryDescription, balls.green, greenRegex)
		balls.blue = getMaxNumberOfBall(summaryDescription, balls.blue, blueRegex)
	}

	return game{
		id,
		balls,
	}
}

func getGamePower(game game) int {
	return game.balls.red * game.balls.blue * game.balls.green
}

func Process(file *os.File) string {
	scanner := bufio.NewScanner(file)
	games := []game{}
	result := 0

	for scanner.Scan() {
		games = append(games, parseGameLine(scanner.Text()))
	}

	for _, game := range games {
		result += getGamePower(game)
	}

	return strconv.Itoa(result)
}
