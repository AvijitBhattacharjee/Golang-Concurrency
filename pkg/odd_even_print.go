package pkg
import (
	"fmt"
	"sync"
)

const limit = 10

func odd(wg *sync.WaitGroup, evenChan, oddChan chan bool) {
	defer wg.Done()
		for i := 1; i <= limit; i += 2 {
			<-oddChan // Wait for turn signal
			fmt.Printf("Odd: %d\n", i)
			evenChan <- true // Hand over to even thread
		}
}


func even(wg *sync.WaitGroup, evenChan, oddChan chan bool) {
	defer wg.Done()
		for i := 2; i <= limit; i += 2 {
			<-evenChan // Wait for turn signal
			fmt.Printf("Even: %d\n", i)
			if i < limit {
				oddChan <- true // Hand over to odd thread
			}
		}
	}

func OddEvenPrint() {
	// Channels to signal turns
	oddChan := make(chan bool)
	evenChan := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)

	// Thread 1: Odd Numbers
	go odd(&wg,evenChan, oddChan)

	// Thread 2: Even Numbers
	go even(&wg, evenChan, oddChan)

	// Start the sequence
	evenChan <- true
	wg.Wait()
}