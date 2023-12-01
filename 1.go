package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLiteralDigit(input string, startindex int) (int, error) {
	const FAIL = -1
	const LITERAL_DIGIT_MAX_LENGTH = 5
	LITERAL_DIGITS := [...]string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	var builder strings.Builder
	for i := startindex; i < len(input) && i < startindex+LITERAL_DIGIT_MAX_LENGTH; i++ {
		builder.WriteByte(input[i])

		for index, literalDigit := range LITERAL_DIGITS {
			if literalDigit == builder.String() {
				return index + 1, nil
			}
		}
	}

	return FAIL, errors.New("NO LITERAL DIGIT FOUND")
}

func getNumberFromInput(input string) int {
	digits := make([]int, 0, 2)

	for index, char := range input {
		parsedLiteral, err := parseLiteralDigit(input, index)
		if err == nil {
			if len(digits) < 2 {
				digits = append(digits, parsedLiteral)
				continue
			}

			digits[1] = parsedLiteral
		}

		num, err := strconv.Atoi(string(char))
		if err == nil {
			if len(digits) < 2 {
				digits = append(digits, num)
				continue
			}

			digits[1] = num
		}
	}

	if len(digits) == 0 {
		return 0
	}

	if len(digits) == 1 {
		digits = append(digits, digits[0])
	}

	return digits[0]*10 + digits[1]
}

func main() {
	file, err := os.Open("inputs/input-1.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		result += getNumberFromInput(scanner.Text())
	}

	fmt.Println(result)
}
