package pkg

import (
	"fmt"
	"sync"
	"time"
)

type WorkLimiters struct {
	numberOfWorker int
	numberOfJobs   int
	wg             sync.WaitGroup
	jobs           chan int
	limiter        <-chan time.Time
}

func RateLimiter() {

	var WorkLimiter = WorkLimiters{
		numberOfWorker: 3,
		numberOfJobs:   10,
		jobs:           make(chan int, 10),

		// global rate limiter
		// only 1 request allowed every 2 sec
		limiter: time.Tick(2 * time.Second),
		wg:      sync.WaitGroup{},
	}

	for i := 0; i < WorkLimiter.numberOfWorker; i++ {
		WorkLimiter.wg.Add(1)
		go WorkLimiter.workLimit(i)
	}

	WorkLimiter.receive()

	close(WorkLimiter.jobs)
	WorkLimiter.wg.Wait()
	fmt.Println("all jobs are finished")
}

func (wl *WorkLimiters) workLimit(id int) {

	defer wl.wg.Done()
	for job := range wl.jobs {

		// rate limiter blocks workers
		<-wl.limiter
		fmt.Printf("%d worker started = %d\n", id, job)
		time.Sleep(2 * time.Second)
		fmt.Printf("%d finished\n", id)
	}
}

func (wl *WorkLimiters) receive() {
	for i := 0; i < wl.numberOfJobs; i++ {
		wl.jobs <- i
	}
}
