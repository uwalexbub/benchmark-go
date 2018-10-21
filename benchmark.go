package main

import (
	"fmt"
	"os"
	"time"
)

const MIN_CARDINALITY int = 100
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
	collectionType string
	collectionSize int
	measurements   Measurements
}

const reportHeader string = "#CollectionType,CollectionSize,AvgInsertNanos,AvgDeleteNanos,AvgSearchNano,AvgIterationNanos"

var workload Workload = Workload{
	insertOps:    100,
	deleteOps:    100,
	searchOps:    100,
	iterationOps: 100,
}

var allBenchmarkers []Benchmarker = []Benchmarker{
	sliceBenchmarker{},
	arrayListBenchmarker{},
	treesetBenchmarker{},
	mapBenchmarker{},
}

func main() {
	allReports := []Report{}
	for _, benchmarker := range allBenchmarkers {
		fmt.Println("====================================================")
		for n := MIN_CARDINALITY; n <= MAX_CARDINALITY; n = n * 10 {
			fmt.Printf("Running a benchmark test for %q collection with size %d\n",
				benchmarker.getName(),
				n)
			b := Benchmark{
				dataSize: n,
				workload: workload,
			}
			measurements := benchmarker.run(b)
			report := Report{
				collectionType: benchmarker.getName(),
				collectionSize: n,
				measurements:   measurements,
			}
			printReport(report)
			allReports = append(allReports, report)
		}
	}

	writeReports(allReports)
}

func printReport(report Report) {
	fmt.Printf("Result: { %s }\n", report.ToTabString())
}

func writeReports(reports []Report) {
	t := time.Now()
	filename := fmt.Sprintf("report-%d-%02d-%02d-%02d%02d%02d.csv",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	fmt.Printf("Writing reports to file %q\n", filename)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(fmt.Sprintf("%s\n", reportHeader))
	for _, v := range reports {
		file.WriteString(fmt.Sprintf("%s\n", v.ToCsvString()))
	}

}

func (report Report) ToCsvString() string {
	return fmt.Sprintf("%s,%d,%d,%d,%d,%d",
		report.collectionType,
		report.collectionSize,
		report.measurements.avgInsertNanos,
		report.measurements.avgDeleteNanos,
		report.measurements.avgSearchNanos,
		report.measurements.avgIterationNanos)
}

func (report Report) ToTabString() string {
	return fmt.Sprintf(
		"CollectionType: %s, "+
			"CollectionSize: %d, "+
			"AvgInsertNanos: %d, "+
			"AvgDeleteNanos: %d, "+
			"AvgSearchNanos: %d, "+
			"AvgIterationNanos: %d",
		report.collectionType,
		report.collectionSize,
		report.measurements.avgInsertNanos,
		report.measurements.avgDeleteNanos,
		report.measurements.avgSearchNanos,
		report.measurements.avgIterationNanos)
}
