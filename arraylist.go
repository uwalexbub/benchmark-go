// This implements benchmarking of ArrayList data structure
// from 3-rd party library https://github.com/emirpasic/gods#arraylist
package main

import (
	"fmt"
	"math/rand"

	"github.com/emirpasic/gods/lists/arraylist"
)

// Dummy struct that implements the benchmarker interface.
// In Go language, a type implements an interface if it implements all the methods of that interface.
type arrayListBenchmarker struct {
}

func (this arrayListBenchmarker) getName() string {
	return "ArrayList"
}

func (this arrayListBenchmarker) run(b Benchmark) Measurements {
	fmt.Printf("Benchmarking %q with data size %d\n", this.getName(), b.dataSize)

	initRandSeed()
	data := this.initialize(b.dataSize)

	measurements := Measurements{}
	measurements.avgInsertNanos = this.runInserts(data, b.workload.insertOps)
	measurements.avgDeleteNanos = this.runDeletes(data, b.workload.deleteOps)
	measurements.avgSearchNanos = this.runSearches(data, b.workload.searchOps)
	measurements.avgIterationNanos = this.runIterations(data, b.workload.iterationOps)

	return measurements
}

func (this arrayListBenchmarker) initialize(n int) *arraylist.List {
	fmt.Println("Initializing")
	data := arraylist.New()
	for i := 0; i < n; i++ {
		value := getUniqueValue()
		data.Add(value)
	}
	return data
}

func (this arrayListBenchmarker) runInserts(data *arraylist.List, n int) int64 {
	fmt.Println("Running insert operations")
	data_copy := arraylist.New()
	copyArrayList(data_copy, data)
	var opTimeSum int64 = 0
	for i := 0; i < n; i++ {
		// Insert in a random position
		insert_index := rand.Intn(data_copy.Size())
		value := getUniqueValue()
		action := func() {
			data_copy.Insert(insert_index, value)
		}
		opTimeSum += measureAction(action)
	}

	return avg(opTimeSum, n)
}

func (this arrayListBenchmarker) runDeletes(data *arraylist.List, n int) int64 {
	fmt.Println("Running delete operations")
	data_copy := arraylist.New()
	copyArrayList(data_copy, data)
	var opTimeSum int64 = 0
	for i := 0; i < int(n); i++ {
		del_index := rand.Intn(data_copy.Size())
		action := func() {
			data_copy.Remove(del_index)
		}
		opTime := measureAction(action)
		opTimeSum += opTime
	}

	return avg(opTimeSum, n)
}

func (this arrayListBenchmarker) runSearches(data *arraylist.List, n int) int64 {
	fmt.Println("Running search operations")
	var opTimeSum int64 = 0
	for i := 0; i < n; i++ {
		// Pick any index at random and get the value by that index,
		// then search that value in the array. That way we know for sure that value will exist.
		index := rand.Intn(data.Size())
		value, _ := data.Get(index)
		action := func() {
			if !data.Contains(value) {
				panic(fmt.Sprintf("Value %q not found in the array list", value))
			}
		}
		opTimeSum += measureAction(action)
	}

	return avg(opTimeSum, n)
}

func copyArrayList(dest *arraylist.List, src *arraylist.List) {
	for _, value := range src.Values() {
		dest.Add(value)
	}
}

func (this arrayListBenchmarker) runIterations(data *arraylist.List, n int) int64 {
	fmt.Println("Running iteration operations")
	var sumNanos int64 = 0
	for i := 0; i < n; i++ {
		action := func() {
			for index, value := range data.Values() {
				dummyAction(index, value.(string))
			}
		}
		sumNanos += measureAction(action)
	}

	return avg(sumNanos, n)
}
