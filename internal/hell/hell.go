package hell

import "sync"

type NonceData struct {
	Nonce int64
	Hash  []byte
}

type HellExecution struct {
	OnAllFinished chan bool
	OnNonceFound  chan *NonceData
}

type hellContext struct {
	mu                sync.Mutex
	devils            []*devil
	nonceFoundChannel chan *NonceData
	stopped           bool
}

func newHellContext() *hellContext {
	return &hellContext{
		devils:            make([]*devil, 0),
		nonceFoundChannel: make(chan *NonceData),
		stopped:           false,
	}
}

func (hellContextInstance *hellContext) stopHell() {
	hellContextInstance.mu.Lock()
	hellContextInstance.stopped = true
	hellContextInstance.mu.Unlock()
}

func (hellContextInstance *hellContext) hellStopped() bool {
	hellContextInstance.mu.Lock()
	defer hellContextInstance.mu.Unlock()
	return hellContextInstance.stopped
}

func (hellContextInstance *hellContext) CreateDevil() *devil {
	gate := make(chan bool)

	devilInstance := &devil{
		gate: gate,
	}

	hellContextInstance.devils = append(hellContextInstance.devils, devilInstance)

	return devilInstance
}

func StartFire(numNonce int64, numDevils int) *HellExecution {
	hellContextInstance := newHellContext()
	workFinished := make(chan bool)

	split := numNonce / int64(numDevils)
	rest := numNonce % int64(numDevils)

	for i := 0; i < numDevils; i++ {
		devilInstance := hellContextInstance.CreateDevil()

		start := split * int64(i)
		end := (split * int64(i+1))

		if i+1 == numDevils && rest > 0 {
			end += rest
		}

		go devilInstance.enslaveDevil(hellContextInstance, start, end)
	}

	go func() {
		for _, devilContext := range hellContextInstance.devils {
			<-devilContext.gate
		}

		workFinished <- true
	}()

	return &HellExecution{
		OnNonceFound:  hellContextInstance.nonceFoundChannel,
		OnAllFinished: workFinished,
	}
}
