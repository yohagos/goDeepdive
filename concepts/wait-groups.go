package concepts

import (
	"fmt"
	"sync"
	"time"
)

func worker3(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			worker3(i)
		})
	}

	wg.Wait()
}
