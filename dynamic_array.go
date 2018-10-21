// This implements benchmarking of Go's slice data structure which is
// an equivalent of a dynamic array in other programming languages
package main

import (
	"fmt"
	"math/rand"
)

// Dummy struct that implements the benchmarker interface.
// In Go language, a type implements an interface if it implements all the methods of that interface.
type dynamicArrayBenchmarker struct {
}

func (this dynamicArrayBenchmarker) getName() string {
	return "Dynamic Array"
}

func (this dynamicArrayBenchmarker) run(b Benchmark) Report {
	fmt.Printf("Benchmarking %q with data size %d\n", this.getName(), b.dataSize)

	initRandSeed()
	data := this.initialize(b.dataSize)

	report := Report{
		dataStructureName: this.getName(),
		dataSize:          b.dataSize,
	}

	report.avgInsertNanos = this.runInserts(data, b.workload.insertOps)
	report.avgDeleteNanos = this.runDeletes(data, b.workload.deleteOps)
	report.avgSearchNanos = this.runSearches(data, b.workload.searchOps)
	report.avgIterationNanos = this.runIterations(data, b.workload.iterationOps)

	return report
}

func (this dynamicArrayBenchmarker) initialize(n int) []string {
	fmt.Printf("Initializing %q\n", this.getName())
	data := make([]string, n)
	for i := 0; i < n; i++ {
		data[i] = getUniqueValue()
	}

	return data
}

func (this dynamicArrayBenchmarker) runInserts(data []string, n int) int64 {
	fmt.Println("Running insert operations")
	data_copy := make([]string, len(data))
	copy(data_copy, data)
	var opTimeSum int64 = 0
	for i := 0; i < n; i++ {
		// Insert in a random position
		insert_index := rand.Intn(len(data_copy))
		value := getUniqueValue()
		action := func() {
			data_copy = append(data_copy[:insert_index],
				append([]string{value}, data_copy[insert_index:]...)...)
		}
		opTimeSum += measureAction(action)
	}

	return avg(opTimeSum, n)
}

func (this dynamicArrayBenchmarker) runDeletes(data []string, n int) int64 {
	fmt.Println("Running delete operations")
	data_copy := make([]string, len(data))
	copy(data_copy, data)
	var opTimeSum int64 = 0
	for i := 0; i < int(n); i++ {
		del_index := rand.Intn(len(data_copy))
		action := func() {
			data_copy = append(data_copy[:del_index], data_copy[del_index+1:]...)
		}
		opTime := measureAction(action)
		opTimeSum += opTime
	}

	return avg(opTimeSum, n)
}

func (this dynamicArrayBenchmarker) runSearches(a []string, n int) int64 {
	fmt.Println("Running search operations")
	var opTimeSum int64 = 0
	for i := 0; i < n; i++ {
		// Pick any index at random and get the value by that index,
		// then search that value in the array. That way we know for sure that value will exist.
		index := rand.Intn(len(a))
		value := a[index]
		action := func() {
			this.search(a, value)
		}
		opTimeSum += measureAction(action)
	}

	return avg(opTimeSum, n)
}

func (this dynamicArrayBenchmarker) search(array []string, value string) int {
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return i
		}
	}

	return -1
}

func (this dynamicArrayBenchmarker) runIterations(data []string, n int) int64 {
	fmt.Println("Running iteration operations")
	var sumNanos int64 = 0
	for i := 0; i < n; i++ {
		action := func() {
			// This is Go iterator over a collection
			for index, value := range data {
				dummyAction(index, value)
			}
		}
		sumNanos += measureAction(action)
	}

	return avg(sumNanos, n)
}
