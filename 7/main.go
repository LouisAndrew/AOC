package day_7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	content string
	value   int
	index   int
}

var CARDS = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

func getCardValue(card string) int {
	for i, c := range CARDS {
		if c == card {
			return i
		}
	}

	return -1
}

func getHandValue(hand string) int {
	cards := strings.Split(hand, "")
	handMap := make(map[string]int)

	const HIGH_CARD, ONE_PAIR, TWO_PAIR, THREE_OF_A_KIND, FULL_HOUSE, FOUR_OF_A_KIND, FIVE_OF_A_KIND = 0, 1, 2, 3, 4, 5, 6

	for _, card := range cards {
		_, ok := handMap[card]
		if ok {
			handMap[card]++
		} else {
			handMap[card] = 1
		}
	}

	jokerCount, ok := handMap["J"]
	if !ok {
		jokerCount = 0
	}

	switch len(handMap) {
	case 5:
		return HIGH_CARD + jokerCount
	case 4:
		if jokerCount > 0 {
			return THREE_OF_A_KIND
		}

		return ONE_PAIR
	case 3:
		for index, value := range handMap {
			if value == 3 {
				if index == "J" || jokerCount == 1 {
					return FOUR_OF_A_KIND
				}

				return THREE_OF_A_KIND
			}
		}

		if jokerCount == 2 {
			return FOUR_OF_A_KIND
		}

		if jokerCount == 1 {
			return FULL_HOUSE
		}

		return TWO_PAIR
	case 2:
		for index, value := range handMap {
			if value == 4 {
				if index == "J" {
					return FIVE_OF_A_KIND
				}

				return FOUR_OF_A_KIND + jokerCount
			}
		}

		if jokerCount != 0 {
			return FIVE_OF_A_KIND
		}

		return FULL_HOUSE
	case 1:
		return FIVE_OF_A_KIND
	}

	return -1
}

func compareCards(h1, h2 Hand) int {
	h1Contents := strings.Split(h1.content, "")
	h2Contents := strings.Split(h2.content, "")

	for i, h1Card := range h1Contents {
		h1CardValue := getCardValue(h1Card)
		h2CardValue := getCardValue(h2Contents[i])

		if h1CardValue > h2CardValue {
			return -1
		}

		if h1CardValue < h2CardValue {
			return 1
		}
	}

	return 0
}

func compareHand(h1, h2 Hand) int {
	if h1.value > h2.value {
		return -1
	}

	if h1.value < h2.value {
		return 1
	}

	return compareCards(h1, h2)
}

func sortHands(hands []Hand) []Hand {
	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			if compareHand(hands[i], hands[j]) == -1 {
				hands[i], hands[j] = hands[j], hands[i]
			}
		}
	}

	return hands
}

func Process(file *os.File) string {
	scanner := bufio.NewScanner(file)
	index := 0
	hands := []Hand{}
	bets := []int{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		content := line[0]
		bet, _ := strconv.Atoi(line[1])
		value := getHandValue(content)

		hands = append(hands, Hand{content, value, index})
		bets = append(bets, bet)
		index++
	}

	hands = sortHands(hands)
	result := 0

	for i, hand := range hands {
		result += bets[hand.index] * (i + 1)
	}

	return strconv.Itoa(result)
}
