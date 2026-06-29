package pkg

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	numberOfJobs    int
	numberOfWorkers int
	jobs            chan int
	limiter         <-chan time.Time
	callback        func(int)
	wg              sync.WaitGroup
}

func combination() {

	// Stop the whole program after 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	callback := func(job int) {
		fmt.Printf(">>> Callback: Job %d completed\n", job)
	}

	w := &Worker{
		numberOfJobs:    10,
		numberOfWorkers: 3,
		jobs:            make(chan int, 10),

		// Global rate limiter:
		// Only ONE job starts every 2 seconds
		limiter: time.Tick(2 * time.Second),

		callback: callback,
	}

	// Start workers
	for i := 0; i < w.numberOfWorkers; i++ {
		w.wg.Add(1)
		go w.special_worker(i, ctx)
	}

	// Produce jobs
	for i := 0; i < w.numberOfJobs; i++ {
		w.jobs <- i
	}

	close(w.jobs)

	w.wg.Wait()

	fmt.Println("Process Finished")
}

func (w *Worker) special_worker(id int, ctx context.Context) {

	defer w.wg.Done()

	for {
		select {

		case <-ctx.Done():
			fmt.Printf("Worker %d timed out\n", id)
			return

		case job, ok := <-w.jobs:

			if !ok {
				return
			}

			// Wait for rate limiter OR timeout
			select {

			case <-ctx.Done():
				fmt.Printf("Worker %d timed out\n", id)
				return

			case <-w.limiter:
			}

			fmt.Printf("Worker %d started Job %d\n", id, job)

			time.Sleep(100 * time.Millisecond)

			fmt.Printf("Worker %d completed Job %d\n", id, job)

			w.callback(job)
		}
	}
}