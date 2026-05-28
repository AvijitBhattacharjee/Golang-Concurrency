package pkg

import (
	"fmt"
	"sync"
	"time"
)

const BufferSize = 3

func producer(buffer chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {

		buffer <- i

		fmt.Printf("Produced: %d\n", i)

		time.Sleep(500 * time.Millisecond)
	}

	close(buffer)
}

func consumer(buffer chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range buffer {

		fmt.Printf("Consumed: %d\n", item)

		time.Sleep(1 * time.Second)
	}
}

func StartProducerConsumer() {

	buffer := make(chan int, BufferSize)

	var wg sync.WaitGroup

	// Start Producer
	wg.Add(1)
	go producer(buffer, &wg)

	// Start Consumer
	wg.Add(1)
	go consumer(buffer, &wg)

	wg.Wait()

	fmt.Println("Producer Consumer Finished")
}
