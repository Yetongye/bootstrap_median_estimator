package main

// This program performs bootstrap resampling to estimate the standard error of the median
/*
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Generate synthetic data: 100 normal samples
	data := make([]float64, 100)
	for i := range data {
		data[i] = randNorm(0, 1)
	}

	start := time.Now()

	// Perform 1000 bootstrap resamples
	B := 1000
	medians := Bootstrap(data, B)

	// Calculate standard error of the median
	mean := 0.0
	for _, m := range medians {
		mean += m
	}
	mean /= float64(B)

	variance := 0.0
	for _, m := range medians {
		variance += (m - mean) * (m - mean)
	}
	stdErr := math.Sqrt(variance / float64(B))

	fmt.Printf("Standard Error of the Median: %.6f\n", stdErr)
	fmt.Printf("Execution Time: %v\n", time.Since(start))
}


*/

//concurrent bootstrap resampling
import (
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

func main() {
	// Initialize pprof for profiling
	go func() {
		log.Println("Starting pprof on :6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	// Command line flags for bootstrap parameters
	b := flag.Int("b", 1000, "Number of bootstrap resamples")
	n := flag.Int("n", 100, "Sample size")
	workers := flag.Int("w", 4, "Number of goroutines")
	flag.Parse()

	// Validate input parameters
	if *b <= 0 || *n <= 0 || *workers <= 0 {
		fmt.Println("Error: All input values (-b, -n, -w) must be greater than 0.")
		os.Exit(1)
	}

	// Initialize logger to write to a file
	logFile, err := os.OpenFile("bootstrap.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Print program information
	fmt.Println("This is a bootstrap median estimator using Go:")
	fmt.Printf("Using B=%d, sample size=%d, workers=%d\n\n", *b, *n, *workers)

	// Generate synthetic data: n normal samples
	// Seed the random number generator
	data := make([]float64, *n)
	for i := range data {
		data[i] = randNorm(0, 1)
	}

	// Perform bootstrap resampling
	start := time.Now()
	medians := BootstrapConcurrent(data, *b, *workers)
	duration := time.Since(start)

	mean := 0.0
	for _, m := range medians {
		mean += m
	}
	mean /= float64(*b)

	variance := 0.0
	for _, m := range medians {
		variance += (m - mean) * (m - mean)
	}
	stdErr := math.Sqrt(variance / float64(*b))

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Print results
	fmt.Printf("Standard Error of the Median: %.6f\n", stdErr)
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Used: %.2f KB\n", float64(m.Alloc)/1024.0)

	log.Printf("Standard Error: %.6f", stdErr)
	log.Printf("Execution Time: %v", duration)
	log.Printf("Memory Used: %.2f KB", float64(m.Alloc)/1024.0)
}
