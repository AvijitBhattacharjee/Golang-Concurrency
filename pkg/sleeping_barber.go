package pkg

import (
	"fmt"
	"sync"
	"time"
)

const Chairs = 3

func barber(customers chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for customer := range customers {
		fmt.Printf("Barber is cutting hair of customer %d\n", customer)

		time.Sleep(time.Second)

		fmt.Printf("Barber finished customer %d\n", customer)
	}

	fmt.Println("Barber is going home")
}

func customer(id int, customers chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	select {

	case customers <- id:
		fmt.Printf("Customer %d entered waiting room\n", id)

	default:
		fmt.Printf("Customer %d left, no empty chair\n", id)
	}
}

func StartSleepingBarber() {

	customers := make(chan int, Chairs)

	var wg sync.WaitGroup

	// Start barber
	wg.Add(1)
	go barber(customers, &wg)

	// Customers arriving
	for i := 1; i <= 5; i++ {

		wg.Add(1)
		go customer(i, customers, &wg)

		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(5 * time.Second)

	close(customers)

	wg.Wait()
}
