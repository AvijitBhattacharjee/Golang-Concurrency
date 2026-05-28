package pkg

import (
	"fmt"
	"time"
)

func Goroutine_lekage() {
	ch := make(chan int)

	go work(ch)

	time.Sleep(3 * time.Second)

	fmt.Println("main finished")

}

func work(ch chan int) {

	fmt.Println("worker started")

	// waiting forever as no one sending data to the channel
	// data := <-ch:
	// fmt.Println("received:", data)

	// handling go routine leakage with
	// 1. timeout 2. closing the channel

	select {
	case data := <-ch:
		fmt.Println("received:", data)
	case <-time.After(time.Second * 2):
		fmt.Println("time out exceeded")
	}
}
