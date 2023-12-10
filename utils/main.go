package utils

import (
	"regexp"
	"strconv"
)

func NumberRegexp() *regexp.Regexp {
	// won't compile integers!
	// return regexp.MustCompile(`\d+`)
	return regexp.MustCompile(`-?\d*\.{0,1}\d+`)
}

func MakeRange(start, end int) []int {
	arr := make([]int, end-start+1)
	for i := range arr {
		arr[i] = start + i
	}

	return arr
}

func ParseNumbers(content string) []int {

	numberRegex := NumberRegexp()
	numbers := numberRegex.FindAllString(content, -1)

	result := make([]int, len(numbers))
	for i, numberAsString := range numbers {
		number, _ := strconv.Atoi(numberAsString)
		result[i] = number
	}

	return result
}

func Max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func Min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}

func RemoveFromArray[T comparable](array []T, index int) []T {
	array[index] = array[len(array)-1]
	return array[:len(array)-1]
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
