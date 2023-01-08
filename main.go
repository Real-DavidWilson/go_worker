package main

import (
	"main/internal/hell"
	"runtime"
)

func main() {
	totalNonce, totalWorkers := 100, runtime.NumCPU()

	hellOnFire := hell.StartFire(totalNonce, totalWorkers)

	<-hellOnFire
}
