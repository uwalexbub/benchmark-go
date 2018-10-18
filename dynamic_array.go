package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Benchmark_dynamic_array(n int, p Params) {
	fmt.Printf("Benchmarking dynamic array with cardinality %d\n", n)

	rand.Seed(time.Now().UnixNano())
	a := initialize(n)

	n_deletes := int(p.deletes_portion * float32(n))
	run_deletes(a, n_deletes)

	n_inserts := int(p.inserts_portion * float32(n))
	run_inserts(a, n_inserts)

	n_searches := int(p.search_portion * float32(n))
	run_searches(a, n_searches)

	n_iterations := int(p.iterations_portion * float32(n))
	run_iterations(a, n_iterations)
}

func initialize(n int) []int {
	fmt.Println("Initializing")
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(n)
	}

	return a
}

func run_deletes(a []int, n int) {
	fmt.Println("Running delete operations")
	for i := 0; i < int(n); i++ {
		del_index := rand.Intn(len(a))
		a = append(a[:del_index], a[del_index+1:]...)
	}
}

func run_inserts(a []int, n int) {
	fmt.Println("Running insert operations")
	for i := 0; i < n; i++ {
		// Insert in a random position
		insert_index := rand.Intn(len(a))
		value := rand.Intn(n)
		a = append(a[:insert_index],
			append([]int{value}, a[insert_index:]...)...)
	}
}

func run_searches(a []int, n int) {
	fmt.Println("Running search operations")
	for i := 0; i < n; i++ {
		// Pick any index at random and get the value by that index,
		// then search that value in the array. That way we know for sure that value will exist.
		index := rand.Intn(len(a))
		value := a[index]
		search(a, value)
	}
}

func search(array []int, value int) int {
	for i := 0; i < len(array); i++ {
		if value == array[i] {
			return i
		}
	}

	return -1
}

func run_iterations(a []int, n int) {
	fmt.Println("Running iteration operations")
	for i := 0; i < n; i++ {
		for j := 0; j < len(a); j++ {
			// do nothing
		}
	}
}
