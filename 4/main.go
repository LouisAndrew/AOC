package day_4

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type WinningNumbers struct {
	values []int
}

func (w *WinningNumbers) Includes(card int) bool {
	for _, value := range w.values {
		if value == card {
			return true
		}
	}

	return false
}

type Card struct {
	id     int
	result int
	amount int
}

func numberRegexp() *regexp.Regexp {
	return regexp.MustCompile(`\d+`)
}

func getWinningNumbers(line string) WinningNumbers {
	matches := numberRegexp().FindAllString(line, -1)
	values := make([]int, 0, len(matches))
	for _, match := range matches {
		number, _ := strconv.Atoi(match)
		values = append(values, number)
	}

	return WinningNumbers{values}
}

func parseCardLine(line string) Card {
	splittedByColon := strings.Split(line, ":")
	gameId := regexp.MustCompile(`\d+`).FindString(splittedByColon[0])
	id, _ := strconv.Atoi(gameId)
	cardContents := strings.Split(splittedByColon[1], "|")

	winningNumbers := getWinningNumbers(cardContents[0])
	myNumbers := numberRegexp().FindAllString(cardContents[1], -1)
	result := 0

	for _, numberAsStr := range myNumbers {
		number, _ := strconv.Atoi(numberAsStr)
		if winningNumbers.Includes(number) {
			result++
		}
	}

	amount := 1

	return Card{
		id,
		result,
		amount,
	}
}

func getTotalCards(cards []Card) int {
	result := 0

	for index, card := range cards {
		result += card.amount

		for i := index + 1; i <= index+card.result && i < len(cards); i++ {
			cards[i].amount += card.amount
		}
	}

	return result
}

func Process(file *os.File) string {
	scanner := bufio.NewScanner(file)
	cards := []Card{}

	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, parseCardLine(line))
	}

	result := getTotalCards(cards)

	return strconv.Itoa(result)
}
