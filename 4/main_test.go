package day_4

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseCardLine(t *testing.T) {
	want := Card{1, 4, 1}
	got := parseCardLine("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	got = parseCardLine("Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11")
	want = Card{6, 0, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	got = parseCardLine("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1")
	want = Card{3, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetWinningNumbers(t *testing.T) {
	got := getWinningNumbers("41 48 83 86 17")
	want := WinningNumbers{[]int{41, 48, 83, 86, 17}}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetTotalCards(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	cards := make([]Card, 0, 5)
	for _, line := range strings.Split(input, "\n") {
		cards = append(cards, parseCardLine(line))
	}

	want := 30
	got := getTotalCards(cards)

	if got != want {
		t.Fatalf("expected: %d, got: %d", want, got)
	}
}
