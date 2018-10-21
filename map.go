// This implements benchmarking of Go's map data structure
// which is an equivalent of hash map in other programming languages
package main

import (
	"fmt"
	"math/rand"
)

// Dummy struct that implements the benchmarker interface.
// In Go language, a type implements an interface if it implements all the methods of that interface.
type mapBenchmarker struct {
}

type mapType map[string]int

func (this mapBenchmarker) getName() string {
	return "Map"
}

func (this mapBenchmarker) run(b Benchmark) Report {
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

func (this mapBenchmarker) initialize(n int) mapType {
	fmt.Printf("Initializing %q\n", this.getName())
	data := make(mapType)
	for i := 0; i < n; i++ {
		key := getUniqueValue()
		data[key] = rand.Intn(n)
	}
	return data
}

func (this mapBenchmarker) runInserts(data mapType, n int) int64 {
	fmt.Println("Running insert operations")
	data_copy := make(mapType, len(data))
	copyMap(data_copy, data)
	var sumNanos int64 = 0
	for i := 0; i < n; i++ {
		key := getUniqueValue()
		value := rand.Intn(len(data_copy))
		action := func() {
			data_copy[key] = value
		}
		sumNanos += measureAction(action)
	}

	return avg(sumNanos, n)
}

func (this mapBenchmarker) runDeletes(data mapType, n int) int64 {
	fmt.Println("Running delete operations")
	data_copy := make(mapType, len(data))
	copyMap(data_copy, data)
	var sumNanos int64 = 0
	i := 0
	for key, _ := range data {
		if rand.Intn(len(data))%n == 0 {
			action := func() {
				delete(data_copy, key)
			}
			sumNanos += measureAction(action)
			i++
		}
		if i >= n {
			break
		}
	}

	return avg(sumNanos, n)
}

func (this mapBenchmarker) runSearches(data mapType, n int) int64 {
	fmt.Println("Running search operations")
	data_copy := make(mapType, len(data))
	copyMap(data_copy, data)
	var sumNanos int64 = 0
	i := 0
	for key, _ := range data {
		if rand.Intn(len(data))%n == 0 {
			action := func() {
				value, present := data_copy[key]
				if present {
					dummyAction(value, key)
				} else {
					panic(fmt.Sprintf("Key %q not found in the map", key))
				}
			}
			sumNanos += measureAction(action)
			i++
		}
		if i >= n {
			break
		}
	}

	return avg(sumNanos, n)
}

func (this mapBenchmarker) runIterations(data mapType, n int) int64 {
	fmt.Println("Running iteration operations")
	var sumNanos int64 = 0
	for i := 0; i < n; i++ {
		action := func() {
			for key, value := range data {
				dummyAction(value, key)
			}
		}
		sumNanos += measureAction(action)
	}

	return avg(sumNanos, n)
}

func copyMap(dest mapType, src mapType) {
	for key, value := range src {
		dest[key] = value
	}
}
