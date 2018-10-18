package main

const MIN_CARDINALITY int = 1000 * 10

// TODO: Adjust in the end
const MAX_CARDINALITY int = 1000 * 100

type Params struct {
	inserts_portion    float32
	search_portion     float32
	deletes_portion    float32
	iterations_portion float32
}

type benchmark_action func(int /* cardinality */, Params)

type result struct {
	total_nanos     int64
	insert_nanos    int64
	search_nanos    int64
	delete_nanos    int64
	iteration_nanos int64
}

func main() {
	p := Params{
		inserts_portion:    0.25,
		search_portion:     0.25,
		deletes_portion:    0.25,
		iterations_portion: 0.25,
	}

	run_benchmark(Benchmark_dynamic_array, p)
}

func run_benchmark(target benchmark_action, p Params) {
	for n := MIN_CARDINALITY; n <= MAX_CARDINALITY; n = n * 10 {
		Benchmark_dynamic_array(n, p)
	}
}
