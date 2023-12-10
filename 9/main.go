package day_9

import (
	"aoc-2023/utils"
	"bufio"
	"os"
	"strconv"
)

// generateNextSequence considering last digit of `current` is already appended to `lastDigits`
func generateNextSequence(current, firstDigits []int) ([]int, []int) {
	next := []int{}
	isAllZeroes := true

	for i := 0; i < len(current)-1; i++ {
		next = append(next, current[i+1]-current[i])
	}

	firstDigits = append(firstDigits, next[0])
	for _, value := range next {
		if value != 0 {
			isAllZeroes = false
			break
		}
	}

	if isAllZeroes {
		return next, firstDigits
	}

	return generateNextSequence(next, firstDigits)
}

func extrapolateResult(lastDigits []int) int {
	length := len(lastDigits) - 1
	result := lastDigits[length]
	for i := length - 1; i >= 0; i-- {
		result = lastDigits[i] - result
	}

	return result
}

func Process(file *os.File) string {
	scanner := bufio.NewScanner(file)
	fileContent := [][]int{}
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		digits := utils.NumberRegexp().FindAllString(line, -1)
		numbers := []int{}

		for _, digit := range digits {
			number, _ := strconv.Atoi(digit)
			numbers = append(numbers, number)
		}

		fileContent = append(fileContent, numbers)
	}

	for _, readings := range fileContent {
		_, firstDigits := generateNextSequence(readings, []int{readings[0]})
		result += extrapolateResult(firstDigits)
	}

	return strconv.Itoa(result)

}
