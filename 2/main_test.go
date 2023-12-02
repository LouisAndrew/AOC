package day_2

import (
	"reflect"
	"testing"
)

func TestParseGameLine(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	want := game{
		id: 1,
		balls: balls{
			red:   4,
			blue:  6,
			green: 2,
		},
	}

	got := parseGameLine(input)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	input = "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	want = game{
		id: 3,
		balls: balls{
			red:   20,
			blue:  6,
			green: 13,
		},
	}
	got = parseGameLine(input)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
