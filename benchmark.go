package main

import (
	"fmt"
)

const MIN_CARDINALITY int = 1000
const MAX_CARDINALITY int = 1000 * 1000

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
	run(Benchmark) Measurements
}

type Measurements struct {
	avgInsertNanos    int64
	avgDeleteNanos    int64
	avgSearchNanos    int64
	avgIterationNanos int64
}

type Report struct {
	dataStructureName string
	dataSize          int
	measurements      Measurements
}

var workload Workload = Workload{
	insertOps:    100,
	deleteOps:    100,
	searchOps:    100,
	iterationOps: 100,
}

var allBenchmarkers []Benchmarker = []Benchmarker{
	sliceBenchmarker{},
	arrayListBenchmarker{},
	mapBenchmarker{},
	treesetBenchmarker{},
}

func main() {
	allReports := []Report{}
	for _, benchmarker := range allBenchmarkers {
		fmt.Println("=============================")
		for n := MIN_CARDINALITY; n <= MAX_CARDINALITY; n = n * 10 {
			fmt.Printf("Running %q benchmark with data size %d\n", benchmarker.getName(), n)
			b := Benchmark{
				dataSize: n,
				workload: workload,
			}
			measurements := benchmarker.run(b)
			report := Report{
				dataStructureName: benchmarker.getName(),
				dataSize:          n,
				measurements:      measurements,
			}
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
		report.measurements.avgInsertNanos,
		report.measurements.avgDeleteNanos,
		report.measurements.avgSearchNanos,
		report.measurements.avgIterationNanos)
}
