package main

import (
	"fmt"
	"math/rand"
)

type mapBenchmarker struct{}

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
	var opTimeSum int64 = 0
	for i := 0; i < n; i++ {
		key := getUniqueValue()
		value := rand.Intn(len(data_copy))
		action := func() {
			data_copy[key] = value
		}
		opTimeSum += measureAction(action)
	}

	return avg(opTimeSum, n)
}

func (this mapBenchmarker) runDeletes(data mapType, n int) int64 {
	fmt.Println("Running deletes")
	data_copy := make(mapType, len(data))
	copyMap(data_copy, data)
	var opTimeSum int64 = 0
	for i := 0; i < n; i++ {
		key := getUniqueValue()
		value := rand.Intn(len(data_copy))
		action := func() {
			data_copy[key] = value
		}
		opTimeSum += measureAction(action)
	}

	return avg(opTimeSum, n)

}

func copyMap(dest mapType, src mapType) {
	for key, value := range src {
		dest[key] = value
	}
}
