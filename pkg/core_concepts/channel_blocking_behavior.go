package coreconcepts

import (
	"fmt"
	"sync"
	"time"
)

func ChannelBlockingBehavior() {

	// This creates a slice of nil channels. We need to initialize each channel before using it.
	// Sending/receiving on a nil channel blocks forever.
	// So, if we try to send or receive on any of these channels without initializing them, the program will block indefinitely.
	var ch = make([]chan int, 3)

	// Initialize each channel in the slice
	for i := 0; i < len(ch); i++ {
		ch[i] = make(chan int)
	}
	var wg sync.WaitGroup

	wg.Add(2)
	go assign(&wg, ch)
	go print(&wg, ch)

	wg.Wait()
	fmt.Println("process is finished!")

}

func assign(wg *sync.WaitGroup, ch []chan int) {
	defer wg.Done()

	for i := 0; i < len(ch); i++ {
		ch[i] <- i
		time.Sleep(1 * time.Second)
		fmt.Println("assigning = ", i)
	}
}

func print(wg *sync.WaitGroup, ch []chan int) {

	defer wg.Done()

	for i := 0; i < len(ch); i++ {
		fmt.Println("here it is = ", <-ch[i])
		time.Sleep(1 * time.Second)
	}
}
