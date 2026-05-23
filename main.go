package main

import "github.com/AvijitBhattacharjee/Golang-Concurrency/pkg"

func main() {

	// dining philosopher problem
	pkg.StartDining()

	// Sleeping Barber
	pkg.StartSleepingBarber()

	pkg.StartProducerConsumer()
}