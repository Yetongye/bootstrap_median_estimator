package main

import "sort"

// Median calculates the median of a slice of float64 numbers.
// If the slice is empty, it returns 0.
func Median(data []float64) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	sorted := make([]float64, n)
	copy(sorted, data)
	sort.Float64s(sorted)
	mid := n / 2
	if n%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2.0
	}
	return sorted[mid]
}
