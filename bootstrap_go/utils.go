package main

//utils.go contains utility functions for the bootstrap resampling program
// It includes functions for generating normally distributed random numbers and initializing the logger.
import (
	"log"
	"math"
	"math/rand"
	"os"
)

// randNorm returns a normally distributed value
func randNorm(mu, sigma float64) float64 {
	u1 := rand.Float64()
	u2 := rand.Float64()
	z := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return mu + sigma*z
}

// InitLogger sets up the logger to write to file
func InitLogger() {
	logFile, err := os.OpenFile("bootstrap.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
