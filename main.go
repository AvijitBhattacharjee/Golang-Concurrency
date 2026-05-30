package main

import "github.com/AvijitBhattacharjee/Golang-Concurrency/pkg"
import "github.com/AvijitBhattacharjee/Golang-Concurrency/pkg/core_concepts"

func main() {

	// dining philosopher problem
	// pkg.StartDining()

	// // // Sleeping Barber
	// pkg.StartSleepingBarber()

	// // // producer-consumer
	// pkg.StartProducerConsumer()

	// // printing ood-even concurrently
	// pkg.OddEvenPrint()

	// // Dekker's algorithm
	// pkg.DekkersAlgorithm()

	// // channel_blocking_behavior
	// pkg.ChannelBlockingBehavior()

	// // timeOut_channel
	// pkg.TimeOutChannel()

	// // goroutine_lekage
	// //pkg.PlayBook()

	// // implemnting worker pool
	// pkg.WorkerPools()

	// pkg.RateLimiter()

	// pkg.Goroutine_lekage()


	// pkg.Build_H2O()

	// pkg.Thread_Cache()

	// pkg.Reader_Writer()

	pkg.Task_Scheduler()

	pkg.PipelineConcurrency()

	pkg.WebCrawler()
	
	core_concepts.ContextCancellation()

	pkg.core_concepts.ErrorGroup()
}