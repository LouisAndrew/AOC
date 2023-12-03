package day_3

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func makeRange(start, end int) []int {
	arr := make([]int, end-start+1)
	for i := range arr {
		arr[i] = start + i
	}

	return arr
}

type Axis struct {
	start int
	end   int
}

func (axis *Axis) validateAxis(min, max int) {
	if axis.start < min {
		axis.start = min
	}

	if axis.end >= max {
		axis.end = max - 1
	}
}

func (axis *Axis) makeRange() []int {
	return makeRange(axis.start, axis.end)
}

func (axis *Axis) isCoordinateWithin(coordinate int) bool {
	if coordinate >= axis.start && coordinate <= axis.end {
		return true
	}

	return false
}

type Box struct {
	xAxis Axis
	yAxis Axis
}

func (box *Box) isBoxWithin(outerBox Box) {
	box.xAxis.validateAxis(outerBox.xAxis.start, outerBox.xAxis.end)
	box.yAxis.validateAxis(outerBox.yAxis.start, outerBox.yAxis.end)
}

func (box *Box) isCoordinateWithin(x, y int) bool {
	return box.xAxis.isCoordinateWithin(x) && box.yAxis.isCoordinateWithin(y)
}

type Numberbox struct {
	value      int
	indexRange []int
}

func getIndicesOfSymbols(line string) []int {
	// ^ -> negation
	// ^. -> not a dot
	// ^\p{L} -> unicode letter
	// symbolRegex, _ := regexp.Compile(`[^\d|.|\s]`) <- part one

	symbolRegex, _ := regexp.Compile(`[*]`)
	matches := symbolRegex.FindAllStringIndex(line, -1)
	symbolIndices := []int{}

	// Extract only the single index of symbols
	for _, index := range matches {
		// adjacent symbols
		if index[1]-index[0] > 1 {
			for _, i := range makeRange(index[0], index[1]-1) {
				symbolIndices = append(symbolIndices, index[i])
			}
		} else {
			symbolIndices = append(symbolIndices, index[0])
		}
	}

	return symbolIndices
}

func getNumbers(line string) []Numberbox {
	numberRegex, _ := regexp.Compile(`[\d]+`)
	matches := numberRegex.FindAllStringIndex(line, -1)

	numberBoxes := make([]Numberbox, 0, len(matches))
	for _, indexRange := range matches {
		value, _ := strconv.Atoi(line[indexRange[0]:indexRange[1]])
		numberBoxes = append(numberBoxes, Numberbox{value, indexRange})
	}

	return numberBoxes
}

// Part one
func isNumberValid(numberIndices []int, numberLinePosition int, symbolIndices [][]int, textBlock Box) bool {
	xAxis := Axis{numberIndices[0] - 1, numberIndices[1]}
	yAxis := Axis{numberLinePosition - 1, numberLinePosition + 1}
	box := Box{xAxis, yAxis}
	box.isBoxWithin(textBlock)

	for _, yIndex := range box.yAxis.makeRange() {
		for _, symbolXIndex := range symbolIndices[yIndex] {
			if box.isCoordinateWithin(symbolXIndex, yIndex) {
				return true
			}
		}
	}

	return false
}

func calculateGearRatio(numberIndices [][]Numberbox, symbolIndex, yIndex int, textBlock Box) int {
	xAxis := Axis{symbolIndex - 1, symbolIndex + 1}
	yAxis := Axis{yIndex - 1, yIndex + 1}
	box := Box{xAxis, yAxis}
	box.isBoxWithin(textBlock)
	left, right := 0, 0

	for _, yIndex := range box.yAxis.makeRange() {
		for _, numberBox := range numberIndices[yIndex] {
			indexRange := numberBox.indexRange
			if box.isCoordinateWithin(indexRange[0], yIndex) || box.isCoordinateWithin(indexRange[1]-1, yIndex) {
				if left == 0 {
					left = numberBox.value
				} else {
					right = numberBox.value
				}

			}
		}
	}

	return left * right
}

func Process(file *os.File) string {
	inputFileLength := 150 // must be modified if input file is larger

	scanner := bufio.NewScanner(file)
	result := 0
	numberBoxIndices := make([][]Numberbox, 0, inputFileLength) // 100 -> must be modified if input file is larger
	symbolIndices := make([][]int, 0, inputFileLength)          // 100 -> must be modified if input file is larger
	inputBoxWidth := 0
	inputBoxHeight := 0

	for scanner.Scan() {
		line := scanner.Text()
		symbolIndices = append(symbolIndices, getIndicesOfSymbols(line))
		numberBoxIndices = append(numberBoxIndices, getNumbers(line))

		if inputBoxWidth == 0 {
			// considering input file shape is a box
			inputBoxWidth = len(line)
		}

		inputBoxHeight++
	}

	inputBox := Box{Axis{0, inputBoxWidth}, Axis{0, inputBoxHeight}}
	for yIndex, symbolRow := range symbolIndices {
		for _, symbolIndex := range symbolRow {
			result += calculateGearRatio(numberBoxIndices, symbolIndex, yIndex, inputBox)
		}
	}

	return strconv.Itoa(result)
}
