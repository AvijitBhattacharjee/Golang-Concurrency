package pkg

import (
	"fmt"
	"sync"
	"time"
)

func OddEvenPrint() {

	var wg sync.WaitGroup

	wg.Add(1)
	go printOdd(&wg)
	wg.Add(1)
	go printEven(&wg)

	wg.Wait()

}

func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("Prinitng odd = ", i)
			time.Sleep(1 * time.Second)
		}
	}
}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Prinitng even = ", i)
			time.Sleep(1 * time.Second)
		}
	}
}
