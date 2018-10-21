package main

import (
	"fmt"
)

const MIN_CARDINALITY int = 1000
const MAX_CARDINALITY int = 100 * 1000

type Workload struct {
	insertOps    int
	deleteOps    int
	searchOps    int
	iterationOps int
}

type Benchmark struct {
	dataSize int
	workload Workload
}

type Benchmarker interface {
	getName() string
	run(Benchmark) Report
}

type Report struct {
	dataStructureName string
	dataSize          int
	avgInsertNanos    int64
	avgDeleteNanos    int64
	avgSearchNanos    int64
	avgIterationNanos int64
}

func main() {
	workload := Workload{
		insertOps:    50,
		deleteOps:    50,
		searchOps:    50,
		iterationOps: 50,
	}

	allBenchmarkers := []Benchmarker{
		dynamicArrayBenchmarker{},
		mapBenchmarker{},
	}

	allReports := []Report{}
	for _, benchmarker := range allBenchmarkers {
		fmt.Println("=============================")
		fmt.Printf("Running %q benchmark\n", benchmarker.getName())

		for n := MIN_CARDINALITY; n <= MAX_CARDINALITY; n = n * 10 {
			b := Benchmark{
				dataSize: n,
				workload: workload,
			}
			report := benchmarker.run(b)
			printReport(report)
			allReports = append(allReports, report)
		}
	}
}

func printReport(report Report) {
	fmt.Printf("#DataStructureName,DataSize,AvgInsertNanos,AvgDeleteNanos,AvgSearchNano,AvgIterationNanos\n")
	fmt.Printf("%s,%d,%d,%d,%d,%d\n",
		report.dataStructureName,
		report.dataSize,
		report.avgInsertNanos,
		report.avgDeleteNanos,
		report.avgSearchNanos,
		report.avgIterationNanos)
}
