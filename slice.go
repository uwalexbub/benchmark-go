// This implements benchmarking of Go's slice data structure which is
// an equivalent of a dynamic array in other programming languages
package main

import (
	"fmt"
	"math/rand"
)

// Dummy struct that implements the benchmarker interface.
// In Go language, a type implements an interface if it implements all the methods of that interface.
type sliceBenchmarker struct {
}

func (this sliceBenchmarker) getName() string {
	return "Slice"
}

func (this sliceBenchmarker) run(b Benchmark) Measurements {
	initRandSeed()
	data := this.initialize(b.dataSize)

	measurements := Measurements{}
	measurements.avgInsertNanos = this.runInserts(data, b.workload.insertOps)
	measurements.avgDeleteNanos = this.runDeletes(data, b.workload.deleteOps)
	measurements.avgSearchNanos = this.runSearches(data, b.workload.searchOps)
	measurements.avgIterationNanos = this.runIterations(data, b.workload.iterationOps)

	return measurements
}

func (this sliceBenchmarker) initialize(n int) []string {
	fmt.Println("Initializing")
	data := make([]string, n)
	for i := 0; i < n; i++ {
		data[i] = getUniqueValue()
	}

	return data
}

func (this sliceBenchmarker) runInserts(data []string, n int) int64 {
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

func (this sliceBenchmarker) runDeletes(data []string, n int) int64 {
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

func (this sliceBenchmarker) runSearches(data []string, n int) int64 {
	fmt.Println("Running search operations")
	var opTimeSum int64 = 0
	for i := 0; i < n; i++ {
		// Pick any index at random and get the value by that index,
		// then search that value in the array. That way we know for sure that value will exist.
		index := rand.Intn(len(data))
		value := data[index]
		action := func() {
			if !this.search(data, value) {
				panic(fmt.Sprintf("Value %q not found in the slice", value))
			}
		}
		opTimeSum += measureAction(action)
	}

	return avg(opTimeSum, n)
}

func (this sliceBenchmarker) search(array []string, value string) bool {
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return true
		}
	}

	return false
}

func (this sliceBenchmarker) runIterations(data []string, n int) int64 {
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
