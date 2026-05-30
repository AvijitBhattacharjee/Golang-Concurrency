package pkg

import "fmt"

func generate(nums ...int) <-chan int {

	out := make(chan int)

	go func() {

		defer close(out)

		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		defer close(out)

		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func filterEven(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		defer close(out)

		for n := range in {

			if n%2 == 0 {
				out <- n
			}
		}
	}()

	return out
}

func PipelineConcurrency() {

	gen := generate(1, 2, 3, 4, 5)

	sq := square(gen)

	filtered := filterEven(sq)

	for val := range filtered {
		fmt.Println(val)
	}
}
