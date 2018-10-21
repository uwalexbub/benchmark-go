// This implements benchmarking of TreeSet data structure
// from 3-rd party library https://github.com/emirpasic/gods#treeset
package main

import (
	"fmt"
	"math/rand"

	"github.com/emirpasic/gods/sets/treeset"
)

// Dummy struct that implements the benchmarker interface.
// In Go language, a type implements an interface if it implements all the methods of that interface.
type treesetBenchmarker struct {
}

func (this treesetBenchmarker) getName() string {
	return "TreeSet"
}

func (this treesetBenchmarker) run(b Benchmark) Measurements {
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

func (this treesetBenchmarker) initialize(n int) *treeset.Set {
	fmt.Println("Initializing")
	data := treeset.NewWithStringComparator()
	for i := 0; i < n; i++ {
		value := getUniqueValue()
		data.Add(value)
	}
	return data
}

func (this treesetBenchmarker) runInserts(data *treeset.Set, n int) int64 {
	fmt.Println("Running insert operations")
	data_copy := treeset.NewWithStringComparator()
	copyTreeSet(data_copy, data)
	var sumNanos int64 = 0
	for i := 0; i < n; i++ {
		value := getUniqueValue()
		action := func() {
			data_copy.Add(value)
		}
		sumNanos += measureAction(action)
	}

	return avg(sumNanos, n)
}

func (this treesetBenchmarker) runDeletes(data *treeset.Set, n int) int64 {
	fmt.Println("Running delete operations")
	data_copy := treeset.NewWithStringComparator()
	copyTreeSet(data_copy, data)
	var sumNanos int64 = 0
	i := 0
	for _, value := range data.Values() {
		if rand.Intn(data.Size())%n == 0 {
			action := func() {
				data_copy.Remove(value)
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

func (this treesetBenchmarker) runSearches(data *treeset.Set, n int) int64 {
	fmt.Println("Running search operations")
	data_copy := treeset.NewWithStringComparator()
	copyTreeSet(data_copy, data)
	var sumNanos int64 = 0
	i := 0
	for _, value := range data.Values() {
		if rand.Intn(data.Size())%n == 0 {
			action := func() {
				if data_copy.Contains(value) {
					dummyAction(0, value.(string))
				} else {
					panic(fmt.Sprintf("Value %q not found in the tree set", value))
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

func (this treesetBenchmarker) runIterations(data *treeset.Set, n int) int64 {
	fmt.Println("Running iteration operations")
	var sumNanos int64 = 0
	for i := 0; i < n; i++ {
		action := func() {
			for _, value := range data.Values() {
				dummyAction(0, value.(string))
			}
		}
		sumNanos += measureAction(action)
	}

	return avg(sumNanos, n)
}

func copyTreeSet(dest *treeset.Set, src *treeset.Set) {
	for _, value := range src.Values() {
		dest.Add(value)
	}
}
