package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const INPUT_CSV_FILE = "./data/sales.csv"
const OUTPUT_CSV_FILE = "./go/salesWithTaxes.csv"
const BENCHMARK_CSV_FILE = "./output/benchmarks.csv"
const TAX_PERCENT = 0.05

var logBench bool = false

type Bench struct {
	language  string
	benchType string
	timeUse   time.Duration
	memoryUse string
}

func getMemoryUsage() uint64 {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return mem.Alloc / 1024 / 1024 // Convert from bytes to mb
}

func benchmark(tag string) func() {
	timeStart := time.Now()
	memoryStart := getMemoryUsage()
	return func() {
		timeEnd := time.Now()
		timeUsage := timeEnd.Sub(timeStart)
		memoryUsage := fmt.Sprintf("%.2vmb", getMemoryUsage()-memoryStart)
		if logBench {
			fmt.Println("-------- BENCHMARK ", strings.ToUpper(tag), " RESULTS --------")
			fmt.Println("time: ", timeUsage, " memory: ", memoryUsage)
			fmt.Println("-------- END BENCHMARK  RESULTS --------")
			fmt.Println()
		}
		benchmarkFile, err := os.OpenFile(BENCHMARK_CSV_FILE, os.O_APPEND, 0644)
		if err != nil {
			fmt.Println(err)
		}
		writer := csv.NewWriter(benchmarkFile)
		writer.Comma = '\t'
		defer writer.Flush()
		bench := []string{fmt.Sprintf("golang,%s,%s,%s", tag, timeUsage, memoryUsage)}
		writer.Write(bench)
	}
}

func main() {
	readBench := benchmark("read file")
	inputFile, err := os.Open(INPUT_CSV_FILE)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(inputFile)
	records, _ := reader.ReadAll()
	readBench()

	processBench := benchmark("process data")
	sales := records[1:]

	for _, sale := range sales {
		price, err := strconv.ParseFloat(sale[2][1:], 32)
		if err != nil {
			fmt.Println(err)
		}
		addTax := fmt.Sprintf("$%.2f", price+price*TAX_PERCENT)

		sale[2] = addTax
	}
	processBench()
	writeBench := benchmark("write file")
	outputFile, err := os.Create(OUTPUT_CSV_FILE)
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	writer.WriteAll(records)
	writeBench()
}
