package day_3

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetIndicesOfSymbols(t *testing.T) {
	want := []int{0, 5}

	got := getIndicesOfSymbols("*$..^*......")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetNumbers(t *testing.T) {
	want := make([]Numberbox, 2)
	want[0] = Numberbox{indexRange: []int{0, 3}, value: 467}
	want[1] = Numberbox{indexRange: []int{5, 8}, value: 114}

	got := getNumbers("467..114..")
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestCalculateGearRatio(t *testing.T) {
	input := "467*114......22*....13"
	numberIndices := getNumbers(input)
	symbolIndices := getIndicesOfSymbols(input)
	fmt.Println(symbolIndices)

	got := calculateGearRatio([][]Numberbox{numberIndices}, symbolIndices[0], 0, Box{Axis{0, len(input)}, Axis{0, 1}})
	want := 467 * 114
	if got != want {
		t.Fatalf("expected: %d, got: %d", want, got)
	}

	got = calculateGearRatio([][]Numberbox{numberIndices}, symbolIndices[1], 0, Box{Axis{0, len()}, Axis{0, 1}})
	want = 0
	if got != want {
		t.Fatalf("expected: %d, got: %d", want, got)
	}
}
