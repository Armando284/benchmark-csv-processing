package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BENCHMARK_CSV_FILE = "./output/benchmarks.csv"

type Bench struct {
	testType string
	_tests   int8
	_timeSum float64
	_memSum  float64
	timeAvg  float64
	memAvg   float64
}

type LanguageBench struct {
	language string
	benchs   [3]Bench
}

func parseBench(bench string) float64 {
	val, err := strconv.ParseFloat(strings.Fields(bench)[0], 2)
	if err != nil {
		panic(err)
	}
	return val
}

func reduceBench(language LanguageBench, time float64, mem float64, i int8) LanguageBench {
	language.benchs[i]._tests++
	language.benchs[i]._timeSum = language.benchs[i]._timeSum + time
	language.benchs[i]._memSum = language.benchs[i]._memSum + mem
	return language
}

func testToNum(testType string) int8 {
	switch testType {
	case "process data":
		return 1
	case "write file":
		return 2
	}
	return 0
}

func main() {
	benchFile, err := os.Open(BENCHMARK_CSV_FILE)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(benchFile)
	records, _ := reader.ReadAll()
	benchRecords := records[1:]

	JavaScript := LanguageBench{"JavaScript", [3]Bench{{"read file", 0, 0, 0, 0, 0}, {"process data", 0, 0, 0, 0, 0}, {"write file", 0, 0, 0, 0, 0}}}
	Python := LanguageBench{"Python", [3]Bench{{"read file", 0, 0, 0, 0, 0}, {"process data", 0, 0, 0, 0, 0}, {"write file", 0, 0, 0, 0, 0}}}
	PhP := LanguageBench{"PhP", [3]Bench{{"read file", 0, 0, 0, 0, 0}, {"process data", 0, 0, 0, 0, 0}, {"write file", 0, 0, 0, 0, 0}}}
	Golang := LanguageBench{"Golang", [3]Bench{{"read file", 0, 0, 0, 0, 0}, {"process data", 0, 0, 0, 0, 0}, {"write file", 0, 0, 0, 0, 0}}}

	for _, bench := range benchRecords {
		switch bench[0] {
		case "javascript":
			testIdx := testToNum(bench[1])

			time := parseBench(bench[2])
			mem := parseBench(bench[3])
			JavaScript = reduceBench(JavaScript, time, mem, testIdx)
		}
	}

	// TODO: Calculate averages
	benchResults := [4]LanguageBench{JavaScript, Python, PhP, Golang}
	fmt.Println(benchResults)
}
