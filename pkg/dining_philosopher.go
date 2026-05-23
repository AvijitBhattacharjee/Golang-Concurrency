package pkg

import (
	"fmt"
	"sync"
	"time"
)

const NumberOfPhilosophers = 5

type Fork struct {
	wg sync.Mutex
}

type Philosopher struct {
	id int
	LeftFork *Fork
	RightFork *Fork
}

func (p *Philosopher)dine(wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("%d is thinking /n", p.id)
	time.Sleep(500*time.Millisecond)

	if p.id %2 == 0 {
		p.LeftFork.wg.Lock()
		p.RightFork.wg.Lock()
	} else {
		p.RightFork.wg.Lock()
		p.LeftFork.wg.TryLock()
	}

	fmt.Printf("%d is eating /n", p.id)
	

	time.Sleep(500*time.Millisecond)
	p.LeftFork.wg.Unlock()
	p.RightFork.wg.Unlock()

	fmt.Printf("%d finished eating /n", p.id)
}

func StartDining() {

	var forks [NumberOfPhilosophers]*Fork

	for i:=0;i<NumberOfPhilosophers;i++ {
		forks[i] = &Fork{}
	}

	var philosophers [NumberOfPhilosophers]Philosopher

	for i:=0;i<NumberOfPhilosophers;i++ {
		philosophers[i] = Philosopher{
			id: i,
			LeftFork: forks[i],
			RightFork: forks[(i+1)%NumberOfPhilosophers],
		}
	}

	var wg sync.WaitGroup

	for i:=0;i<NumberOfPhilosophers;i++ {
		wg.Add(1)
		go philosophers[i].dine(&wg)
	}

	wg.Wait()
}