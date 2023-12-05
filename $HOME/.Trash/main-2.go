package day_5

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type Range struct {
	start int
	end   int
}

func (r Range) inRange(point int) bool {
	return r.start <= point && r.end >= point
}

func (r Range) intersects(r2 Range) bool {
	return r.inRange(r2.start) || r.inRange(r2.end) || r2.inRange(r.start) || r2.inRange(r.end)
}

func Intersection(r1, r2 Range) Range {
	if !r1.intersects(r2) {
		return Range{0, 0}
	}

	start := max(r1.start, r2.start)
	end := min(r1.end, r2.end)

	return Range{start, end}
}

type Map struct {
	src        int
	dest       int
	rangeValue Range
	delta      int
}

type AlmanacMapInstruction struct {
	mapIdentifier string
	source        string
	destination   string
}

func (am AlmanacMapInstruction) getContent(fileContent string) []Map {
	regex := regexp.MustCompile(`(?s)` + am.mapIdentifier + ` map:\n(.*?)\n\n`)
	match := regex.FindStringSubmatch(fileContent)
	if len(match) == 0 {
		return []Map{}
	}

	mapContentsString := strings.Split(match[1], "\n")
	mapContents := make([]Map, 0, len(mapContentsString))

	for _, mapContentString := range mapContentsString {
		mapContent := strings.Split(mapContentString, " ")
		source, _ := strconv.Atoi(mapContent[1])
		destination, _ := strconv.Atoi(mapContent[0])
		rangeAmount, _ := strconv.Atoi(mapContent[2])

		mapContents = append(mapContents, Map{src: source, dest: destination, rangeValue: Range{start: source, end: source + rangeAmount}, delta: destination - source})
	}

	return mapContents
}

func (am AlmanacMapInstruction) readInstruction(fileContent string) {
	content := am.getContent(fileContent)
}

// Don't forget to add 2 new lines on the end of the test file
func ProcessPartTwo(file *os.File) string {
	scanner := bufio.NewScanner(file)
	fileContent := []string{}

	for scanner.Scan() {
		// Dump entire file content
		fileContent = append(fileContent, scanner.Text())
	}

	fileContentString := strings.Join(fileContent, "\n")

	seedsAsStrings := numberRegexp().FindAllString(fileContent[0], -1)
	almanacs := groupSeeds(seedsAsStrings)

	for index, almanacMap := range ALMANAC_MAP {
		go almanacMap.ReadAlmanacMap(fileContentString, &almanacs, index)
	}

	result := 0

	LOCATION_INDEX := 6
	for _, almanac := range almanacs {
		location := almanac.GetField(LOCATION_INDEX)
		if result == 0 || almanac.GetField(LOCATION_INDEX) < result {
			result = location
		}
	}

	return strconv.Itoa(result)
}
