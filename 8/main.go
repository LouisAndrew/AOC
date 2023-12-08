package day_8

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func buildRawNodes(content []string) [][3]string {
	rawNodes := make([][3]string, len(content))

	for i, v := range content {
		splitted := strings.Split(v, " = ")
		value := splitted[0]

		splitted = strings.Split(splitted[1], ", ")
		leftValue := splitted[0][1:]
		rightValue := splitted[1][:len(splitted[1])-1]

		rawNodes[i] = [3]string{value, leftValue, rightValue}
	}

	return rawNodes
}

func buildMap(content []string) map[string][2]string {
	directionMap := make(map[string][2]string)
	rawNodes := buildRawNodes(content)

	for _, rawNode := range rawNodes {
		directionMap[rawNode[0]] = [2]string{rawNode[1], rawNode[2]}
	}

	return directionMap
}

func getStepCount(value string, originalInstruction []string, directionMap map[string][2]string) int {
	instructionLength := len(originalInstruction)
	instruction := originalInstruction
	current := value
	stepCount := 0

	for i := 0; i < instructionLength; i++ {
		if strings.HasSuffix(current, "Z") {
			return stepCount
		}

		if instruction[i] == "R" {
			current = directionMap[current][1]
		} else {
			current = directionMap[current][0]
		}

		stepCount++

		if i == instructionLength-1 {
			instruction = append(instruction, originalInstruction...)
			instructionLength = len(instruction)
		}
	}

	return stepCount
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func Process(file *os.File) string {
	scanner := bufio.NewScanner(file)
	result := 0
	content := []string{}

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	instruction := strings.Split(content[0], "")
	content = content[2:]

	directionMap := buildMap(content)
	current := []string{}

	for key := range directionMap {
		if strings.HasSuffix(key, "A") {
			current = append(current, key)
		}
	}

	steps := make([]int, len(current))
	for i, value := range current {
		steps[i] = getStepCount(value, instruction, directionMap)
	}

	result = steps[0]
	for i := 1; i < len(steps); i++ {
		result = lcm(result, steps[i])
	}

	return strconv.Itoa(result)
}
