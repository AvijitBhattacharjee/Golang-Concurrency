package pkg

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	numberOfJob    int
	numberOfWorker int
	wg             sync.WaitGroup
	jobs           chan int
}

func WorkerPools() {

	var newPool = WorkerPool{
		numberOfJob:    10,
		numberOfWorker: 3,
		jobs:           make(chan int, 10),
	}

	var wg sync.WaitGroup

	for i := 0; i < newPool.numberOfWorker; i++ {
		wg.Add(1)
		go workerJob(i, newPool.jobs, &wg)
	}

	receiveJob(newPool.jobs, newPool.numberOfJob)

	close(newPool.jobs)
	wg.Wait()

	fmt.Println("all jobs finished")

}

func receiveJob(jobs chan int, numberOfJob int) {
	for i := 0; i < numberOfJob; i++ {
		jobs <- i
	}

}

func workerJob(id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d started job %d\n", id, job)
		time.Sleep(2 * time.Second)
		fmt.Printf("worker %d finished job %d\n", id, job)
	}
}
