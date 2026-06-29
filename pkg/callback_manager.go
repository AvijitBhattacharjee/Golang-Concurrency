package pkg

import (
	"fmt"
	"sync"
	"time"
)

type CallbackManager struct {
	eventRunning bool
	mu           sync.Mutex
	cv           *sync.Cond
}

func NewCallbackManager() *CallbackManager {

	c := &CallbackManager{}

	c.cv = sync.NewCond(&c.mu)

	return c
}

func (c *CallbackManager) reg_cb(id int, cb func()) {

	c.mu.Lock()

	for c.eventRunning {
		fmt.Printf("User %d waiting...\n", id)
		c.cv.Wait()
	}

	c.mu.Unlock()

	cb()
}

func (c *CallbackManager) Event() {

	c.mu.Lock()
	c.eventRunning = true
	c.mu.Unlock()

	fmt.Print("\nEvent Started\n")
	time.Sleep(5 * time.Second)
	fmt.Print("\nEvent Finished\n")

	c.mu.Lock()
	c.eventRunning = false
	c.cv.Broadcast()
	c.mu.Unlock()
}

func CallbackManagerEvent() {

	manager := NewCallbackManager()

	var wg sync.WaitGroup

	// Executes immediately
	wg.Add(1)

	go func() {

		defer wg.Done()

		manager.reg_cb(1, func() {

			fmt.Println("Callback 1 Executed")
		})
	}()

	time.Sleep(time.Second)

	// Start Event
	go manager.Event()

	time.Sleep(time.Second)

	// Waits
	for i := 2; i <= 4; i++ {

		wg.Add(1)

		go func(id int) {

			defer wg.Done()

			manager.reg_cb(id, func() {

				fmt.Printf("Callback %d Executed\n", id)
			})

		}(i)
	}

	wg.Wait()
}