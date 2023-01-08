package hell

import (
	"time"
)

type devilContext struct {
	gate chan bool
}

func enslaveDevil(start, end int, onEndChannel chan bool) {
	for i := start; i <= end; i++ {
		time.Sleep(time.Millisecond)
	}

	onEndChannel <- true
}
