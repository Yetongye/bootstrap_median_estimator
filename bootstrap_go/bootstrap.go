package main

/*
import (
	"math/rand"
	"time"
)

// Bootstrap performs B bootstrap resamplings and returns the sample medians
func Bootstrap(data []float64, B int) []float64 {
	n := len(data)
	medians := make([]float64, B)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < B; i++ {
		sample := make([]float64, n)
		for j := 0; j < n; j++ {
			sample[j] = data[rand.Intn(n)]
		}
		medians[i] = Median(sample)
	}
	return medians
}
*/

import (
	"math/rand"
	"sync"
)

// concurrent bootstrap resampling
func BootstrapConcurrent(data []float64, B int, workers int) []float64 {
	n := len(data)
	medians := make([]float64, B)
	wg := sync.WaitGroup{}
	ch := make(chan int, B)

	// Index for work distribution
	for i := 0; i < B; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for i := range ch {
				sample := make([]float64, n)
				for j := 0; j < n; j++ {
					sample[j] = data[rand.Intn(n)]
				}
				medians[i] = Median(sample)
			}
		}()
	}
	wg.Wait()
	return medians
}
