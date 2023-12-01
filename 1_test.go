package main

import "testing"

func TestNumberFromInput(t *testing.T) {
	got := getNumberFromInput("aaa123")
	if got != 13 {
		t.Errorf("getNumberFromInput(\"aaa123\") = %d; want 13", got)
	}

	got = getNumberFromInput("aaa1aaaaaa5")
	if got != 15 {
		t.Errorf("getNumberFromInput(\"aaa1aaaaaa5\") = %d; want 15", got)
	}

	got = getNumberFromInput("aaa1aaaaaa")
	if got != 11 {
		t.Errorf("getNumberFromInput(\"aaa1aaaaaa\") = %d; want 11", got)
	}

	got = getNumberFromInput("aaaaaaaaaaa")
	if got != 0 {
		t.Errorf("getNumberFromInput(\"aaaaaaaaaaa\") = %d; want 0", got)
	}

	got = getNumberFromInput("oneaaa2")
	if got != 12 {
		t.Errorf("getNumberFromInput(\"oneaaa2\") = %d; want 12", got)
	}

	got = getNumberFromInput("oneaaathree")
	if got != 13 {
		t.Errorf("getNumberFromInput(\"oneaaathree\") = %d; want 13", got)
	}

	got = getNumberFromInput("asdsasanine")
	if got != 99 {
		t.Errorf("getNumberFromInput(\"asdsasanine\") = %d; want 99", got)
	}
}

func TestParseLiteralDigit(t *testing.T) {
	got, err := parseLiteralDigit("one", 0)

	if got != 1 || err != nil {
		t.Errorf("parseLiteralDigit(\"one\", 0) = %d; want 1", got)
	}

	got, err = parseLiteralDigit("xtwo", 1)
	if got != 2 || err != nil {
		t.Errorf("parseLiteralDigit(\"xtwo\", 1) = %d; want 2", got)
	}

	got, err = parseLiteralDigit("abcd", 0)
	if got != -1 || err == nil {
		t.Errorf("parseLiteralDigit(\"abcd\", 0) = %d; want -1", got)
	}

	got, err = parseLiteralDigit("abcdef12", 0)
	if got != -1 || err == nil {
		t.Errorf("parseLiteralDigit(\"abcd\", 0) = %d; want -1", got)
	}
}
