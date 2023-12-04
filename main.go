package main

import (
	day_1 "aoc-2023/1"
	day_2 "aoc-2023/2"
	day_3 "aoc-2023/3"
	day_4 "aoc-2023/4"
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
	day := os.Args[1:][0]

	processorMap := map[string]processor{
		"1": {
			filepath: "inputs/input-1.txt",
			process:  day_1.Process,
		},
		"2": {
			filepath: "inputs/input-2.txt",
			process:  day_2.Process,
		},
		"3": {
			filepath: "inputs/input-3.txt",
			process:  day_3.Process,
		},
		"4": {
			filepath: "inputs/input-4.txt",
			process:  day_4.Process,
		},
	}

	processor, ok := processorMap[day]

	if ok {
		file, err := os.Open(processor.filepath)
		checkError(err)

		defer file.Close()
		fmt.Println(processor.process(file))
	} else {
		fmt.Println("No processor found for day", day)
	}
}
