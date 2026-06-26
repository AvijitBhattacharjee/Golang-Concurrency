package pkg

import (
	"fmt"
	"sync"
)


func OddEvenPrint() {

	var wg sync.WaitGroup
	var oddch =make(chan struct{})
	var evench =make(chan struct{})

	wg.Add(1)
	go printOdd(&wg, oddch, evench)

	wg.Add(1)
	go printEven(&wg, evench, oddch)

	evench <- struct{}{}
	wg.Wait()
	fmt.Println("process finished")
}

func printOdd(wg *sync.WaitGroup, myturn, nextturn chan struct{}) {

	defer wg.Done()

	for i:=1;i<10;i=i+2 {
		<-myturn
		fmt.Println("Odd is = ", i)
		if i != 9 {nextturn <- struct{}{}}
		
	}
}


func printEven(wg *sync.WaitGroup, myturn, nextturn chan struct{}) {

	defer wg.Done()

	for i:=0;i<10;i=i+2 {
		<-myturn
		fmt.Println("Odd is = ", i)
		nextturn <- struct{}{}
	}
}
