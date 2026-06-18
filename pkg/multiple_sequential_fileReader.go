package pkg

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	fmt.Println("Implementing Files Reader writer with rate limiter in sequence")

	var numberOfFiles = 20
	var numberOfonsumers = 5
	var files = make(chan int, numberOfFiles)
	var wg sync.WaitGroup
	var limiter = time.Tick(2*time.Second)

	wg.Add(1)
	go producer(&wg, numberOfFiles, files)

	for i:=0;i<numberOfonsumers;i++ {
		wg.Add(1)
		go consumer(&wg, files, limiter, i+1)
	}

	wg.Wait()
	fmt.Println("All files read successfully")
}

func producer(wg *sync.WaitGroup, numberOfFiles int, files chan<- int) {
	defer wg.Done()
	for i:=0;i<numberOfFiles;i++ {
		files <- i
	}
	close(files)
}

func consumer(wg *sync.WaitGroup, files <-chan int, limiter <-chan time.Time, id int) {
	defer wg.Done()
	for file := range files {
		fmt.Printf("%d Reading file %d\n", id, file)
		<-limiter
		fmt.Printf("%d Completed file %d\n", id, file)
	}
}	