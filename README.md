Benchmark data structures in Go language

Guidance:
* Use workload of fixed size for every cardinality
* Size of data should not change by magnitude during benchmarking
* Test iterator to see how it performs

Consider:
* Experiment with different kinds of workloads: balanced vs write-intensive vs read-intensive