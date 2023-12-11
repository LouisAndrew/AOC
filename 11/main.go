package day_11

import (
	"aoc-2023/utils"
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Position [2]int

func galaxyRegexp() *regexp.Regexp {
	return regexp.MustCompile(`(#+)`)
}

func getManhattanDistance(a, b Position) int {
	return utils.Abs(a[1]-b[1]) + utils.Abs(a[0]-b[0])
}

func getGalaxyPairs(galaxies []Position) [][2]Position {
	pairs := [][2]Position{}

	for i, a := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			pairs = append(pairs, [2]Position{a, galaxies[j]})
		}
	}

	return pairs
}

func transpose(a []string) []string {
	b := []string{}
	for i := 0; i < len(a[0]); i++ {
		m := len(a)
		d := make([]string, m)

		for j := 0; j < m; j++ {
			d[j] = string(a[j][i])
		}

		b = append(b, strings.Join(d, ""))
	}

	return b
}

func expandHorizontalSpace(content []string) []int {
	horizontalCoords := []int{}

	for _, line := range content {
		indices := galaxyRegexp().FindAllStringIndex(line, -1)
		if len(indices) == 0 {
			horizontalCoords = append(horizontalCoords, 1)
		} else {
			horizontalCoords = append(horizontalCoords, 0)
		}
	}

	return horizontalCoords
}

func expandVerticalSpace(content []string) []int {
	transposed := transpose(content)
	expanded := expandHorizontalSpace(transposed)

	return expanded
}

func getPosition(current int, expanded []int) int {
	multiplier := 1_000_000
	space := 0
	for i := 0; i < current; i++ {
		if expanded[i] == 1 {
			space += multiplier
		} else {
			space++
		}

	}

	return space
}

func Process(file *os.File) string {
	scanner := bufio.NewScanner(file)
	galaxies := []Position{}
	fileContent := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		fileContent = append(fileContent, line)
	}

	verticalCoords := expandVerticalSpace(fileContent)
	horizontalCoords := expandHorizontalSpace(fileContent)

	for i, row := range fileContent {
		for _, galaxy := range galaxyRegexp().FindAllStringIndex(row, -1) {
			galaxies = append(galaxies, Position{getPosition(i, horizontalCoords), getPosition(galaxy[0], verticalCoords)})
		}
	}

	pairs := getGalaxyPairs(galaxies)

	result := 0
	for _, pair := range pairs {
		result += getManhattanDistance(pair[0], pair[1])
	}

	return strconv.Itoa(result)
}
