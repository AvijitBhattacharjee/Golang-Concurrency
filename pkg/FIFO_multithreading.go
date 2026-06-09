//implementing a fixed-size FIFO queue 
// with specific memory constraints and multi-threading considerations.

package pkg

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type FIFOQueue struct {
	buffer   []int
	capacity int

	head int
	tail int
	size int

	mu sync.Mutex
}

func NewFIFOQueue(capacity int) *FIFOQueue {
	return &FIFOQueue{
		buffer:   make([]int, capacity),
		capacity: capacity,
	}
}

func (q *FIFOQueue) Enqueue(val int) error {

	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size == q.capacity {
		return errors.New("queue full")
	}

	q.buffer[q.tail] = val

	q.tail = (q.tail + 1) % q.capacity

	q.size++

	return nil
}

func (q *FIFOQueue) Dequeue() (int, error) {

	q.mu.Lock()
	defer q.mu.Unlock()

	if q.size == 0 {
		return 0, errors.New("queue empty")
	}

	val := q.buffer[q.head]

	q.head = (q.head + 1) % q.capacity

	q.size--

	return val, nil
}

func (q *FIFOQueue) Size() int {

	q.mu.Lock()
	defer q.mu.Unlock()

	return q.size
}

func FIFOQueue_Test() {

	queue := NewFIFOQueue(5)

	var wg sync.WaitGroup

	// Producers
	for i := 1; i <= 3; i++ {

		wg.Add(1)

		go func(id int) {

			defer wg.Done()

			for j := 0; j < 3; j++ {

				err := queue.Enqueue(id*10 + j)

				if err != nil {
					fmt.Printf("Producer %d: queue full\n", id)
				} else {
					fmt.Printf("Producer %d inserted %d\n",
						id,
						id*10+j)
				}

				time.Sleep(100 * time.Millisecond)
			}

		}(i)
	}

	// Consumers
	for i := 1; i <= 2; i++ {

		wg.Add(1)

		go func(id int) {

			defer wg.Done()

			for j := 0; j < 4; j++ {

				val, err := queue.Dequeue()

				if err != nil {
					fmt.Printf("Consumer %d: queue empty\n", id)
				} else {
					fmt.Printf("Consumer %d removed %d\n",
						id,
						val)
				}

				time.Sleep(150 * time.Millisecond)
			}

		}(i)
	}

	wg.Wait()

	fmt.Println("Final Size:", queue.Size())
}