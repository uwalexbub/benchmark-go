# Benchmark of data collections in Go language

To read the code, start with `benchmark.go`

Notes:
* Data collection is populated with initial randomized data before actual benchmarking takes place 
* Workload of fixed size is used for every cardinality of data collections
* Insert and delete operations are executed against a copy of initiliazed data collection to minimize undesirable impact to search and iteration operations
* [Monotonic clock](https://golang.org/pkg/time/) is used to time operations on data collections
* Results are measured in nanoseconds
