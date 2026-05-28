package main

import "github.com/AvijitBhattacharjee/Golang-Concurrency/pkg"

func main() {

	// dining philosopher problem
	pkg.StartDining()

	// // Sleeping Barber
	pkg.StartSleepingBarber()

	// // producer-consumer 
	pkg.StartProducerConsumer()

	// printing ood-even concurrently
	pkg.OddEvenPrint()

	// Dekker's algorithm
	pkg.DekkersAlgorithm()

	// channel_blocking_behavior
	pkg.ChannelBlockingBehavior()

	// timeOut_channel
	pkg.TimeOutChannel()

	// goroutine_lekage
	//pkg.PlayBook()

	// implemnting worker pool
	pkg.WorkerPools()
}