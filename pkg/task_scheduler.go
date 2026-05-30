package pkg

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID    int
	Delay time.Duration
}

func worker_scheduler(
	id int,
	jobs <-chan Task,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	for task := range jobs {

		fmt.Printf(
			"Worker %d executing Task %d at %s\n",
			id,
			task.ID,
			time.Now().Format("15:04:05"),
		)
	}
}

func scheduleTask(
	task Task,
	jobs chan Task,
) {

	time.Sleep(task.Delay)

	jobs <- task
}

func Task_Scheduler() {

	jobs := make(chan Task)

	var wg sync.WaitGroup

	// start workers
	for i := 1; i <= 3; i++ {

		wg.Add(1)

		go worker_scheduler(
			i,
			jobs,
			&wg,
		)
	}

	// schedule tasks
	go scheduleTask(
		Task{
			ID:    1,
			Delay: 2 * time.Second,
		},
		jobs,
	)

	go scheduleTask(
		Task{
			ID:    2,
			Delay: 5 * time.Second,
		},
		jobs,
	)

	go scheduleTask(
		Task{
			ID:    3,
			Delay: 1 * time.Second,
		},
		jobs,
	)

	time.Sleep(7 * time.Second)

	close(jobs)

	wg.Wait()
}