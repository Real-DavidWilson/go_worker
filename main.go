package main

import (
	"main/internal/hell"
	"math"
	"runtime"
)

func main() {
	totalNonce, totalWorkers := math.MaxInt64, runtime.NumCPU()

	<-hell.StartFire(totalNonce, totalWorkers)
}
