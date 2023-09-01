package main

import "testing"

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Append()
	}
}

func BenchmarkAllocatedAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocatedAppend()
	}
}

func BenchmarkAllocatedAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllocatedAssign()
	}
}
