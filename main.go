package main

import (
	core "github.com/AvijitBhattacharjee/Golang-Concurrency/pkg/core_concepts"
	"github.com/AvijitBhattacharjee/Golang-Concurrency/pkg"
)

func main() {

	// ============================================================
	// CLASSICAL CONCURRENCY & SYNCHRONIZATION PROBLEMS
	// These are well-known operating system and interview problems
	// that focus on coordination, ordering, and deadlock prevention.
	// ============================================================

	// Dining Philosophers Problem
	pkg.StartDining()

	// Sleeping Barber Problem
	pkg.StartSleepingBarber()

	// Building H2O Molecules (2 Hydrogen + 1 Oxygen)
	pkg.Build_H2O()

	// Dekker's Algorithm (Mutual Exclusion)
	pkg.DekkersAlgorithm()

	// Print Odd and Even Numbers Concurrently
	pkg.OddEvenPrint()

	// ============================================================
	// CHANNELS & GOROUTINE FUNDAMENTALS
	// Demonstrates how goroutines communicate, block, timeout,
	// and how leaks can occur when channels are misused.
	// ============================================================

	// Understanding Channel Blocking Behavior
	core.ChannelBlockingBehavior()

	// Implementing Timeouts Using Channels
	core.TimeOutChannel()

	// Detecting and Preventing Goroutine Leaks
	core.Goroutine_lekage()

	// Gracefully Stop Goroutines Using Context Cancellation
	core.ContextCancellation()

	// ============================================================
	// PRODUCTION-GRADE CONCURRENCY PATTERNS
	// Common patterns used in backend services, distributed systems,
	// storage systems, and cloud-native applications.
	// ============================================================

	// Producer Consumer Pattern
	pkg.StartProducerConsumer()

	// Worker Pool Pattern
	pkg.WorkerPools()

	// Rate Limiting Using Tickers and Channels
	pkg.RateLimiter()

	// Pipeline Processing (Generator -> Transform -> Consumer)
	pkg.PipelineConcurrency()

	// Error Propagation Across Concurrent Tasks
	core.ErrorGroup()

	// ============================================================
	// THREAD SAFETY & SHARED STATE
	// Demonstrates safe concurrent access to shared memory using
	// mutexes and read-write locks.
	// ============================================================

	// Thread-Safe Cache using RWMutex
	pkg.Thread_Cache()

	// Multiple Readers and Writers accessing shared data
	pkg.Reader_Writer()

	// ============================================================
	// SYSTEM DESIGN & REAL-WORLD INTERVIEW PROBLEMS
	// Frequently asked in Rubrik, Datadog, CockroachDB,
	// Confluent, Nutanix and other backend interviews.
	// ============================================================

	// Concurrent Task Scheduler
	pkg.Task_Scheduler()

	// Concurrent Web Crawler with Deduplication
	pkg.WebCrawler()

	// Thread-Safe Request Counter grouped by IP and Browser
	pkg.ThreadSafeRequestCounter()
}