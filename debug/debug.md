# Golang Concurrency Debugging Notes

## 1. WaitGroup Add() must happen before goroutine starts

Bad:

```go
for i := 0; i < 5; i++ {
	go func() {
		wg.Add(1)
		defer wg.Done()
	}()
}
wg.Wait()
```

Problem:

* `wg.Wait()` may execute before `Add()`
* Can cause early exit or WaitGroup misuse panic

Correct:

```go
for i := 0; i < 5; i++ {
	wg.Add(1)

	go func() {
		defer wg.Done()
	}()
}
```

---

## 2. Loop Variable Capture

Bad:

```go
for i := 0; i < 5; i++ {
	go func() {
		fmt.Println(i)
	}()
}
```

Possible output:

```text
5
5
5
5
5
```

Fix:

```go
for i := 0; i < 5; i++ {
	go func(i int) {
		fmt.Println(i)
	}(i)
}
```

---

## 3. Race Condition on Shared Variable

Bad:

```go
total += i
```

Multiple goroutines execute simultaneously.

`count++` is NOT atomic.

Internally:

```go
tmp := count
tmp++
count = tmp
```

Fix:

```go
mu.Lock()
total += i
mu.Unlock()
```

or

```go
atomic.AddInt64(...)
```

---

## 4. WaitGroup != Thread Safety

Bad:

```go
var count int

go func() {
	count++
}()
```

Even if:

```go
wg.Wait()
```

exists, it only waits.

It does NOT protect:

```go
count++
```

Use:

```go
Mutex
Atomic
Channel
```

for synchronization.

---

## 5. Unbuffered Channel Requires Sender + Receiver

Bad:

```go
ch := make(chan int)

ch <- 10
```

Runtime:

```text
fatal error: all goroutines are asleep - deadlock!
```

Reason:

No receiver exists.

---

## 6. Send On Closed Channel

Bad:

```go
close(ch)

ch <- 100
```

Runtime:

```text
panic: send on closed channel
```

Rule:

Only sender should close channel.

---

## 7. Closing Channel Multiple Times

Bad:

```go
go func() {
	close(done)
}()

go func() {
	close(done)
}()
```

Runtime:

```text
panic: close of closed channel
```

Fix:

```go
var once sync.Once

once.Do(func() {
	close(done)
})
```

---

## 8. Receiving From Closed Channel

```go
x := <-ch
```

After channel close:

```go
x == zero value
```

Example:

```go
0 for int
"" for string
nil for pointer
```

Safer:

```go
x, ok := <-ch

if !ok {
	return
}
```

Best:

```go
for x := range ch {
}
```

---

## 9. Worker Pool Deadlock

Bad:

```go
for job := range jobs {
}
```

without:

```go
close(jobs)
```

Workers never exit.

```go
wg.Wait()
```

blocks forever.

---

## 10. Goroutine Leak After Channel Close

Bad:

```go
for {
	x := <-ch
	fmt.Println(x)
}
```

After close:

```text
0
0
0
0
...
```

forever.

Fix:

```go
for x := range ch {
}
```

---

## 11. Mutex Is Not Reentrant

Bad:

```go
mu.Lock()

mu.Lock()
```

Same goroutine.

Deadlock.

Go mutexes are NOT reentrant.

---

## 12. Circular Wait Deadlock

Bad:

```go
G1:
mu1.Lock()
mu2.Lock()

G2:
mu2.Lock()
mu1.Lock()
```

Possible state:

```text
G1 owns mu1
G1 waits mu2

G2 owns mu2
G2 waits mu1
```

Deadlock.

Fix:

Always acquire locks in same order.

---

## 13. RWMutex Lock Upgrade Deadlock

Bad:

```go
mu.RLock()

mu.Lock()
```

Current goroutine is a reader.

Writer waits for all readers.

Current goroutine is one of the readers.

Deadlock.

```text
Reader waiting for writer
Writer waiting for reader
```

---

## 14. Check-Then-Act Race

Bad:

```go
c.mu.RLock()

if _, ok := cache[key]; !ok {

	c.mu.RUnlock()

	c.mu.Lock()

	cache[key] = value

	c.mu.Unlock()
}
```

Two goroutines may both observe:

```go
!ok
```

Both decide to initialize.

Not a data race.

Causes duplicate work.

Fix:

Double-check after acquiring write lock.

---

## 15. Context Cancellation

Best practice:

```go
select {

case <-ctx.Done():
	return

case job := <-jobs:
	process(job)
}
```

Meaning:

```text
Either process job
OR stop immediately
```

---

## 16. break inside select != exit goroutine

Bad:

```go
for {

	select {

	case <-ctx.Done():
		break
	}
}
```

`break` exits only:

```go
select
```

NOT:

```go
for
```

Loop continues forever.

Correct:

```go
case <-ctx.Done():
	return
```

---

## 17. Busy Loop

Bad:

```go
for {

	select {

	default:
	}
}
```

Equivalent:

```go
for {
}
```

Consumes 100% CPU.

---

## 18. Context Cancellation Latency

Bad:

```go
for {

	select {

	case <-ctx.Done():
		return

	default:
		time.Sleep(time.Second)
	}
}
```

Cancellation may wait up to 1 second.

Better:

```go
select {

case <-ctx.Done():
	return

case <-time.After(time.Second):
}
```

---

## 19. Atomic OR Mutex

Bad design:

```go
mu.Lock()
counter.Add(1)
mu.Unlock()
```

Atomic already provides synchronization.

Usually choose one strategy:

```text
Atomic
OR
Mutex
```

not both.

---

## 20. Data Race Example

Bad:

```go
done := false

go func() {
	done = true
}()

for !done {
}
```

May appear to work.

Still a data race.

Verify using:

```bash
go run -race main.go
```

---

## 21. Happens-Before Relationship

Channel send:

```go
x = 10

ch <- struct{}{}
```

Receiver:

```go
<-ch

fmt.Println(x)
```

Guaranteed:

```text
10
```

Reason:

```text
send happens-before receive
```

---

## 22. Buffered Channel FIFO

```go
ch := make(chan int, 1)

ch <- 1

go func() {
	ch <- 2
}()

fmt.Println(<-ch)
fmt.Println(<-ch)
```

Always:

```text
1
2
```

Never:

```text
2
1
```

Reason:

Channel preserves FIFO ordering.

---

## 23. Scheduler Ordering Is NOT Guaranteed

Bad assumption:

```go
go worker()
```

means:

```text
worker runs immediately
```

Reality:

```text
worker becomes runnable
```

Scheduler decides execution order.

Never rely on scheduling.

---

## 24. select{} Blocks Forever

```go
select {}
```

Blocks forever.

Often used to keep process alive during debugging.

Can hide deadlocks because main never exits.

---

## 25. Race Detector

Always run:

```bash
go test -race
```

or

```bash
go run -race main.go
```

Detects:

* Read/Write race
* Write/Write race
* Shared variable races

One of the first tools to use when debugging concurrency issues.
