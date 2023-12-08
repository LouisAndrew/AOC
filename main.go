package main

import (
	day_1 "aoc-2023/1"
	day_2 "aoc-2023/2"
	day_3 "aoc-2023/3"
	day_4 "aoc-2023/4"
	day_5 "aoc-2023/5"
	day_6 "aoc-2023/6"
	day_7 "aoc-2023/7"
	day_8 "aoc-2023/8"
	"fmt"
	"os"
)

type processor struct {
	filepath string
	process  func(file *os.File) string
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]
	day := args[0]
	testMode := ""

	if len(args) > 1 {
		testMode = args[1]
	}

	processorMap := map[string]processor{
		"1": {
			filepath: "inputs/input-1",
			process:  day_1.Process,
		},
		"2": {
			filepath: "inputs/input-2",
			process:  day_2.Process,
		},
		"3": {
			filepath: "inputs/input-3",
			process:  day_3.Process,
		},
		"4": {
			filepath: "inputs/input-4",
			process:  day_4.Process,
		},
		"5": {
			filepath: "inputs/input-5",
			process:  day_5.Process,
		},
		"6": {
			filepath: "inputs/input-6",
			process:  day_6.Process,
		},
		"7": {
			filepath: "inputs/input-7",
			process:  day_7.Process,
		},
		"8": {
			filepath: "inputs/input-8",
			process:  day_8.Process,
		},
	}

	processor, ok := processorMap[day]

	if ok {
		filepath := processor.filepath

		if testMode != "" {
			filepath += "-test"
		}

		file, err := os.Open(filepath + ".txt")
		checkError(err)

		defer file.Close()
		fmt.Println(processor.process(file))
	} else {
		fmt.Println("No processor found for day", day)
	}
}
