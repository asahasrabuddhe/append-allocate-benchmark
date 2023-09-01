# Append vs Allocate Benchmark

This is a quick benchmark to understand the differences in performance in between various methods
of appending to a slice in Go.

## The Program
```go
package main

func main() {
	Append()
	AllocatedAppend()
	AllocatedAssign()
}

// Append appends to a slice without pre-allocating the length or capacity.
func Append() {
	s := make([]int, 0)
	for i := 0; i < 100000; i++ {
		s = append(s, i)
	}
}

// AllocatedAppend appends to a slice with pre-allocated capacity.
func AllocatedAppend() {
	s := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		s = append(s, i)
	}
}

// AllocatedAssign assigns to a slice with pre-allocated length and capacity.
func AllocatedAssign() {
	// here when you specify the length and not the capacity, the capacity is set to be equal to the length.
	s := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		s[i] = i
	}
}

```

### The Results
```bash
➜ append-allocate-benchmark go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: append-allocate-benchmark
BenchmarkAppend-10                  4116            288248 ns/op         4101401 B/op         28 allocs/op
BenchmarkAllocatedAppend-10        15831             75964 ns/op          802820 B/op          1 allocs/op
BenchmarkAllocatedAssign-10        18546             65520 ns/op          802820 B/op          1 allocs/op
PASS
ok      append-allocate-benchmark       6.099s
```

### The Analysis

#### Append

This is the slowest and least performant method of appending to a slice. This should only be used when the
size of the destination slice is not known. When appending to a slice without knowing the final size, the `append` 
method automatically doubles the capacity of the slice when it reaches the current capacity. This is a very expensive.

#### AllocatedAppend

This is the second fastest method of appending to a slice. This should be used when the size of the destination slice is known.
When appending to a slice with a known size, the `append` method need not increase the capacity of the slice.

#### AllocatedAssign

This is the fastest method of appending to a slice. This should be used when the size of the destination slice is known.
