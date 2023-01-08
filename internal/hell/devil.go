package hell

import (
	"time"
)

type devil struct {
	gate chan bool
}

func (devilInstance *devil) enslaveDevil(hellContextInstance *hellContext, start, end int64) {
	for i := start; i <= end; i++ {
		if hellContextInstance.hellStopped() {
			break
		}

		time.Sleep(time.Millisecond)

		if i == 1000 {
			hellContextInstance.nonceFoundChannel <- &NonceData{
				Nonce: int64(i),
				Hash:  make([]byte, 0),
			}

			hellContextInstance.stopHell()

			break
		}
	}

	devilInstance.gate <- true
}
