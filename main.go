package main

import (
	"fmt"
	"main/internal/hell"
	"math"
)

func main() {
	totalNonce, totalWorkers := int64(math.MaxInt64), 10

	hellExecution := hell.StartFire(totalNonce, totalWorkers)

	for {
		select {
		case nonceData := <-hellExecution.OnNonceFound:
			fmt.Println("NONCE FOUNDED", nonceData)
		case <-hellExecution.OnAllFinished:
			fmt.Println("ALL WAS FINISHED")
			return
		}
	}
}
