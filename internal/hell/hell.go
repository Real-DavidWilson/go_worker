package hell

import "fmt"

type hellContext struct {
	devils []*devilContext
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

func StartFire(numNonce int, numDevils int) chan bool {
	hellContextInstance := newHellContext()

	closedHell := make(chan bool)
	split := numNonce / numDevils
	rest := numNonce % numDevils

	for i := 0; i < numDevils; i++ {
		deadDevil := hellContextInstance.CreateDevilContext()

		start := split*i + 1
		end := (split * (i + 1))

		if i+1 == numDevils && rest > 0 {
			end += rest
		}

		go enslaveDevil(start, end, deadDevil)
	}

	go func() {
		for _, devilContext := range hellContextInstance.devils {
			<-devilContext.gate
			fmt.Println("A DEVIL WAS DEAD")
		}

		closedHell <- true
	}()

	return closedHell
}
