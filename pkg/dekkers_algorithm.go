package pkg

import (
	"fmt"
	"sync"
	"time"
)

var (
	flag [2]bool
	turn = 0

	counter = 0
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	other := 1 - id

	for i := 0; i < 5; i++ {

		// I want to enter critical section
		flag[id] = true

		// Wait while other goroutine wants to enter
		for flag[other] {

			// Give chance to other goroutine
			if turn != id {

				flag[id] = false

				for turn != id {
				}

				flag[id] = true
			}
		}

		// Critical Section
		fmt.Printf("Goroutine %d entering critical section\n", id)

		counter++

		fmt.Printf("Counter = %d\n", counter)

		time.Sleep(500 * time.Millisecond)

		fmt.Printf("Goroutine %d leaving critical section\n", id)

		// Exit critical section
		turn = other
		flag[id] = false
	}
}

func DekkersAlgorithm() {

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(0, &wg)

	wg.Add(1)
	go worker(1, &wg)

	wg.Wait()

	fmt.Printf("Final Counter = %d\n", counter)
}