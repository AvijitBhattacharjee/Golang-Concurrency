
package pkg

import (
	"fmt"
	"sync"
	"time"
)

const NumPhilosophers = 5

type Fork struct {
	sync.Mutex
}

type Philosopher struct {
	id        int
	leftFork  *Fork
	rightFork *Fork
}

func (p Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {

		// Thinking
		fmt.Printf("Philosopher %d is thinking\n", p.id)
		time.Sleep(time.Millisecond * 500)

		// Deadlock prevention
		// Even philosophers pick right first
		// Odd philosophers pick left first

		if p.id%2 == 0 {

			p.rightFork.Lock()
			fmt.Printf("Philosopher %d picked RIGHT fork\n", p.id)

			p.leftFork.Lock()
			fmt.Printf("Philosopher %d picked LEFT fork\n", p.id)

		} else {

			p.leftFork.Lock()
			fmt.Printf("Philosopher %d picked LEFT fork\n", p.id)

			p.rightFork.Lock()
			fmt.Printf("Philosopher %d picked RIGHT fork\n", p.id)
		}

		// Eating
		fmt.Printf("Philosopher %d is eating\n", p.id)
		time.Sleep(time.Millisecond * 500)

		// Release forks
		p.leftFork.Unlock()
		p.rightFork.Unlock()

		fmt.Printf("Philosopher %d finished eating\n", p.id)
	}
}

func StartDining() {

	var forks [NumPhilosophers]*Fork

	// Initialize forks
	for i := 0; i < NumPhilosophers; i++ {
		forks[i] = &Fork{}
	}

	var philosophers [NumPhilosophers]Philosopher

	// Assign forks to philosophers
	for i := 0; i < NumPhilosophers; i++ {

		philosophers[i] = Philosopher{
			id:        i,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%NumPhilosophers],
		}
	}

	var wg sync.WaitGroup

	// Start dining
	for i := 0; i < NumPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].dine(&wg)
	}

	wg.Wait()

	fmt.Println("All philosophers are done eating")
}