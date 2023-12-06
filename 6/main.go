package day_6

import (
	utils "aoc-2023/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getNumberOfCombinations(distance, maxTime int) int {
	result := 0

	for holdTime := range utils.MakeRange(0, maxTime) {
		timeLeft := maxTime - holdTime
		speed := holdTime

		if timeLeft*speed > distance {
			result += 1
		}
	}

	return result
}

func joinNumbers(numbers []int) int {
	numberAsStrings := make([]string, len(numbers))

	for i, number := range numbers {
		numberAsStrings[i] = strconv.Itoa(number)
	}

	joinedStrings := strings.Join(numberAsStrings, "")
	result, _ := strconv.Atoi(joinedStrings)

	return result
}

func Process(file *os.File) string {
	scanner := bufio.NewScanner(file)

	fileContent := []string{}

	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	maxTimes := utils.ParseNumbers(fileContent[0])
	distances := utils.ParseNumbers(fileContent[1])

	result := getNumberOfCombinations(joinNumbers(distances), joinNumbers(maxTimes))

	return strconv.Itoa(result)
}
