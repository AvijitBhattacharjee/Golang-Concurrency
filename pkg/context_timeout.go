package pkg

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextTimeout() {
	fmt.Println("Implementing context timeout")

	var numberOfJobs = 10
	var numberOfWorker = 3
	var job = make(chan int, numberOfJobs)
	var wg sync.WaitGroup
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg.Add(1)
	go worker_time(ctx, &wg, job)


	for i:=0; i<numberOfWorker;i++ {
		time.Sleep(2*time.Second)
		job <- i
	}

	close(job)
	wg.Wait()
	fmt.Println("Process finished")
}

func worker_time(ctx context.Context, wg *sync.WaitGroup, job chan int) {

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			fmt.Println("working on = ", <-job)
		}
	}
}