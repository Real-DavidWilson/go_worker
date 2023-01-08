package hell

import (
	"fmt"
	"time"
)

type devilContext struct {
	gate chan bool
}

func newHellContext() hellContext {
	return hellContext{
		devils: make([]*devilContext, 0),
	}
}

func (hellContextInstance *hellContext) CreateDevilContext() chan bool {
	gate := make(chan bool)

	devilContext := &devilContext{
		gate: gate,
	}

	hellContextInstance.devils = append(hellContextInstance.devils, devilContext)

	return gate
}

func enslaveDevil(start, end int, endChannel chan bool) {
	for i := start; i <= end; i++ {
		fmt.Println("PROCESSING", i)
		time.Sleep(time.Millisecond * 2)
	}

	endChannel <- true
}
