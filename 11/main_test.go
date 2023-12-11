package day_11

import (
	"reflect"
	"strings"
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	got := getManhattanDistance([2]int{6, 1}, [2]int{11, 5})
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestTranspose(t *testing.T) {
	original := []string{"abc", "def", "ghi"}
	got := transpose(original)
	want := []string{"adg", "beh", "cfi"}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("got %v want %v", got, want)
	}

	got = transpose(want)
	want = original
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("got %v want %v", got, want)
	}
}

func parse(c string) []string {
	return strings.Split(c, "\n")
}

func TestExpand(t *testing.T) {
	want := parse(`....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......`)
	got := expandSpace(parse(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`))

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("got \n%v \n\nwant \n%v", got, want)
	}
}
