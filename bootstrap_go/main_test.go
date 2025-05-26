package main

import (
	"testing"
)

func BenchmarkBootstrapConcurrent(b *testing.B) {
	data := make([]float64, 100)
	for i := range data {
		data[i] = randNorm(0, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = BootstrapConcurrent(data, 1000, 8)
	}
}
