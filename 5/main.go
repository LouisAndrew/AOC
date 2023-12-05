package day_5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func numberRegexp() *regexp.Regexp {
	return regexp.MustCompile(`\d+`)
}

type Almanac struct {
	seed int

	// Optional fields
	values [7]int // soil index 0, fertilizer index 1, etc
}

func (a *Almanac) GetField(index int) int {
	return a.values[index]
}

func (a *Almanac) SetField(index, value int) {
	for i := index; i < len(a.values); i++ {
		a.values[i] = value
	}
}

func buildAlmanac(seed int) Almanac {
	values := [7]int{0, 0, 0, 0, 0, 0}
	alm := Almanac{seed, values}
	alm.SetField(0, seed)

	return alm
}

type AlmanacMapContent struct {
	source      int
	destination int
	rangeAmount int
}

type AlmanacMap struct {
	mapIdentifier string
	source        string
	destination   string
}

func (am AlmanacMap) getContent(fileContent string) []AlmanacMapContent {
	regex := regexp.MustCompile(`(?s)` + am.mapIdentifier + ` map:\n(.*?)\n\n`)
	match := regex.FindStringSubmatch(fileContent)
	if len(match) == 0 {
		return []AlmanacMapContent{}
	}

	mapContentsString := strings.Split(match[1], "\n")
	mapContents := make([]AlmanacMapContent, 0, len(mapContentsString))

	for _, mapContentString := range mapContentsString {
		mapContent := strings.Split(mapContentString, " ")
		source, _ := strconv.Atoi(mapContent[1])
		destination, _ := strconv.Atoi(mapContent[0])
		rangeAmount, _ := strconv.Atoi(mapContent[2])

		mapContents = append(mapContents, AlmanacMapContent{source, destination, rangeAmount})
	}

	return mapContents
}

func (am AlmanacMap) ReadAlmanacMap(fileContent string, ac *Almanac, valueIndex int) {
	content := am.getContent(fileContent)

	for _, almanacContent := range content {
		almanacFieldValue := (*ac).GetField(valueIndex)
		isInRange := almanacFieldValue >= almanacContent.source && almanacFieldValue < almanacContent.source+almanacContent.rangeAmount
		if isInRange {
			delta := almanacContent.destination - almanacContent.source
			destinationValue := almanacFieldValue + delta

			(*ac).SetField(valueIndex, destinationValue)
		}

	}
}

var ALMANAC_MAP = []AlmanacMap{
	{
		mapIdentifier: "seed-to-soil",
		source:        "seed",
	},
	{
		mapIdentifier: "soil-to-fertilizer",
		source:        "Soil",
		destination:   "Fertilizer",
	},
	{
		mapIdentifier: "fertilizer-to-water",
		source:        "Fertilizer",
		destination:   "Water",
	},
	{
		mapIdentifier: "water-to-light",
		source:        "Water",
		destination:   "Light",
	},
	{
		mapIdentifier: "light-to-temperature",
		source:        "Light",
		destination:   "Temperature",
	},
	{
		mapIdentifier: "temperature-to-humidity",
		source:        "Temperature",
		destination:   "Humidity",
	},
	{
		mapIdentifier: "humidity-to-location",
		source:        "Humidity",
		destination:   "Location",
	},
}

func groupSeeds(seedsAsStrings []string) [][2]int {
	seeds := make([]int, 0, len(seedsAsStrings))
	for _, seedAsString := range seedsAsStrings {
		seed, _ := strconv.Atoi(seedAsString)
		seeds = append(seeds, seed)
	}

	seedGroups := make([][2]int, len(seeds)/2)
	for i, seed := range seeds {
		seedGroupIndex := i / 2
		seedGroups[seedGroupIndex][i%2] = seed
	}

	return seedGroups
}

func runForSeedGroup(fileContent string, seedGroup [2]int, minValues *[]int, index int, wg *sync.WaitGroup) {
	length := seedGroup[1]
	result := make([]int, 0, length)
	x := make([]int, 0, length)
	fmt.Printf("length: %v\n", length)

	for i := 0; i < length; i++ {
		alm := buildAlmanac(seedGroup[0] + i)

		for index, almanacMap := range ALMANAC_MAP {
			almanacMap.ReadAlmanacMap(fileContent, &alm, index)
		}

		result = append(result, alm.GetField(6))
		x = append(x, alm.seed)
	}

	sort.Ints(result)
	minValue := result[0]
	(*minValues)[index] = minValue

	defer wg.Done()
}

// Don't forget to add 2 new lines on the end of the test file
func Process(file *os.File) string {
	scanner := bufio.NewScanner(file)
	fileContent := []string{}

	for scanner.Scan() {
		// Dump entire file content
		fileContent = append(fileContent, scanner.Text())
	}

	fileContentString := strings.Join(fileContent, "\n")

	seedsAsStrings := numberRegexp().FindAllString(fileContent[0], -1)
	seedGroups := groupSeeds(seedsAsStrings)
	minValues := make([]int, len(seedGroups))

	var wg sync.WaitGroup
	wg.Add(len(seedGroups))

	for i, seedGroup := range seedGroups {
		fmt.Printf("seedGroup: %v\n", seedGroup)
		go runForSeedGroup(fileContentString, seedGroup, &minValues, i, &wg)
	}
	wg.Wait()
	/* for index, almanacMap := range ALMANAC_MAP {
		go almanacMap.ReadAlmanacMap(fileContentString, &almanacs, index)
	} */

	sort.Ints(minValues)
	result := minValues[0]
	fmt.Printf("minValues: %v\n", minValues)

	return strconv.Itoa(result)
}
