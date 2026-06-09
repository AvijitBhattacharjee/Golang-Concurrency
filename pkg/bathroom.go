package pkg

import (
	"fmt"
	"sync"
	"time"
)

const Capacity = 3

type Bathroom struct {
	mu sync.Mutex
	cv *sync.Cond

	inside int

	currentGender string
}

func NewBathroom() *Bathroom {

	b := &Bathroom{}

	b.cv = sync.NewCond(&b.mu)

	return b
}

func (b *Bathroom) EnterMan(name string) {

	b.mu.Lock()

	for b.currentGender == "W" ||
		b.inside == Capacity {

		b.cv.Wait()
	}

	b.currentGender = "M"
	b.inside++

	fmt.Printf("%s entered (M), inside=%d\n",
		name,
		b.inside)

	b.mu.Unlock()
}

func (b *Bathroom) LeaveMan(name string) {

	b.mu.Lock()

	b.inside--

	fmt.Printf("%s left (M), inside=%d\n",
		name,
		b.inside)

	if b.inside == 0 {
		b.currentGender = ""
	}

	b.cv.Broadcast()

	b.mu.Unlock()
}

func (b *Bathroom) EnterWoman(name string) {

	b.mu.Lock()

	for b.currentGender == "M" ||
		b.inside == Capacity {

		b.cv.Wait()
	}

	b.currentGender = "W"
	b.inside++

	fmt.Printf("%s entered (W), inside=%d\n",
		name,
		b.inside)

	b.mu.Unlock()
}

func (b *Bathroom) LeaveWoman(name string) {

	b.mu.Lock()

	b.inside--

	fmt.Printf("%s left (W), inside=%d\n",
		name,
		b.inside)

	if b.inside == 0 {
		b.currentGender = ""
	}

	b.cv.Broadcast()

	b.mu.Unlock()
}
func Bathroom_Test() {

	b := NewBathroom()

	var wg sync.WaitGroup

	people := []struct {
		name   string
		gender string
	}{
		{"John", "M"},
		{"Bob", "M"},
		{"Mike", "M"},
		{"Alice", "W"},
		{"Mary", "W"},
		{"Emma", "W"},
	}

	for _, p := range people {

		wg.Add(1)

		go func(name, gender string) {

			defer wg.Done()

			if gender == "M" {

				b.EnterMan(name)

				time.Sleep(2 * time.Second)

				b.LeaveMan(name)

			} else {

				b.EnterWoman(name)

				time.Sleep(2 * time.Second)

				b.LeaveWoman(name)
			}

		}(p.name, p.gender)
	}

	wg.Wait()
}