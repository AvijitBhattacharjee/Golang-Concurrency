
1. Questions 

	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			total += i
			fmt.Println(i)
		}()
	}
	wg.Wait()

In this code wg.Add has been added inside the for loop which is not correct
also total+= i is Multiple goroutines read/write simultaneously.

so 

wg.Add(1)

for i:=0;i<5;i++ {
	go func(i int) {
		mu.Lock()
		defer mu.Unlock()

		total+= i
	}(i)
}


2. select {}
blocks forever

3. close(ch)
	ch <- 100
sending on a close channel deadlock

4. two consequitive lock screates deadlock

mu.RLock()

	if _, ok := cache[key]; !ok {

		mu.Lock()

		cache[key] = 1

Goroutine

RLock()
   ↓
readers = 1

Lock()
   ↓
wait for readers=0

But I am the reader

Deadlock


5. go func() {

			defer wg.Done()

			close(done)

		}()

multiple close 
close on a closed channel - deadlock



6. select {

case <-ctx.Done():
	return

case <-time.After(time.Second):
}

best cases for context cancellation


7. either atomic or mutex
not both bcz atomic already provides synchronzation

8. ch := make(chan int, 1)

ch <- 1

go func() {
	ch <- 2
}()

fmt.Println(<-ch)

fmt.Println(<-ch)


capacity = 1
size = 0

after ch<-1 size becomes full (1)
the main goroutine hits the first fmt <-ch and buffer becomes 0
ch<- 2
then 2nd fmt print 2

