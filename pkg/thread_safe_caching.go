package pkg

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Cache struct {
	data map[int]string
	mu   sync.RWMutex
}

func (c *Cache) Set(key int, val string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = val
}

func (c *Cache) Get(key int) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[key]
	if ok {
		return val
	}
	return "no value found"
}

func Thread_Cache() {

	cache := Cache{
		data: make(map[int]string),
	}

	for i := 0; i < 3; i++ {
		cache.Set(i, strconv.Itoa(i))
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go cache.reader(&wg)

	wg.Wait()
	fmt.Println("Execution finished")
}

func (c *Cache) reader(wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i < 3; i++ {
		value := c.Get(i)
		fmt.Println("this is the key = ", value)
		time.Sleep(1 * time.Second)
	}
}
