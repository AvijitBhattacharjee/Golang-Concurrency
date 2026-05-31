// Implement a thread-safe data structure that can keep track of the
// number of incoming requests grouped by IP Address over a time window. Add support for
// grouping by other attributes such as BrowserAgent.

package pkg

import (
	"fmt"
	"sync"
	"time"
)

type RequestCounter struct {
	mu sync.Mutex
	IP map[string][]time.Time
	browser map[string][]time.Time
}

func NewRequestCounter() *RequestCounter {
	return &RequestCounter{
		IP: make(map[string][]time.Time),
		browser: make(map[string][]time.Time),
	}
}

func (r *RequestCounter)addRequest(IP string, browser string) {
	defer r.mu.Unlock()
	r.mu.Lock()

	r.IP[IP] = append(r.IP[IP], time.Now())
	r.browser[browser] = append(r.browser[browser], time.Now())
}

func (r *RequestCounter)countByIP(IP string, window time.Duration) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	var count = 0 

	for _, request := range r.IP[IP] {
		if time.Since(request) <= window {
			count++
		}
	}
	return count
}

func ThreadSafeRequestCounter() {
	counter := NewRequestCounter()

	counter.addRequest(
		"10.1.1.1",
		"Chrome",
	)

	counter.addRequest(
		"10.1.1.1",
		"Chrome",
	)

	counter.addRequest(
		"10.1.1.1",
		"Firefox",
	)

	counter.addRequest(
		"10.1.1.2",
		"Chrome",
	)

	fmt.Println(
		"IP Count:",
		counter.countByIP(
			"10.1.1.1",
			time.Minute,
		),
	)
}