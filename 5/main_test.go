package day_5

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	r1 := Range{start: 0, end: 10}
	r2 := Range{start: 5, end: 15}

	result := Intersection(r1, r2)

	if result.start != 5 || result.end != 10 {
		t.Errorf("Expected 5, 10 but got %d, %d", result.start, result.end)
	}

	r2 = Range{start: 15, end: 20}

	result = Intersection(r1, r2)
	if result.start != 0 || result.end != 0 {
		t.Errorf("Expected empty range but got %d, %d", result.start, result.end)
	}

	r2 = Range{start: 3, end: 10}
	result = Intersection(r1, r2)
	if result.start != 3 || result.end != 10 {
		t.Errorf("Expected 3, 10 but got %d, %d", result.start, result.end)
	}

	r1 = Range{start: 8, end: 12}
	result = Intersection(r1, r2)
	if result.start != 8 || result.end != 10 {
		t.Errorf("Expected 8, 10 but got %d, %d", result.start, result.end)
	}
}
