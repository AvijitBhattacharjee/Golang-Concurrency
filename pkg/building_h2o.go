package pkg

import (
	"fmt"
	"sync"
)

type H2O struct {
	oxygenChan   chan func()
	hydrogenChan chan func()
	done         chan struct{}
}

func NewH2O() *H2O {

	h := &H2O{
		oxygenChan:   make(chan func()),
		hydrogenChan: make(chan func()),
		done:         make(chan struct{}, 3),
	}

	go h.makeH2O()

	return h
}

func (h *H2O) makeH2O() {

	for {

		// wait for exactly 2 Hydrogen
		h1 := <-h.hydrogenChan
		h2 := <-h.hydrogenChan

		// wait for exactly 1 Oxygen
		o := <-h.oxygenChan

		// release atoms
		h1()
		h2()
		o()

		fmt.Println(" -> H2O formed")

		// unblock waiting goroutines
		h.done <- struct{}{}
		h.done <- struct{}{}
		h.done <- struct{}{}
	}
}

func (h *H2O) Hydrogen(
	releaseHydrogen func(),
) {

	h.hydrogenChan <- releaseHydrogen

	<-h.done
}

func (h *H2O) Oxygen(
	releaseOxygen func(),
) {

	h.oxygenChan <- releaseOxygen

	<-h.done
}

func Build_H2O() {

	h2o := NewH2O()

	// 4H + 2O = 2 water molecules
	input := "OOHHHH"

	var wg sync.WaitGroup

	wg.Add(len(input))

	for _, ch := range input {

		if ch == 'H' {

			go func() {

				defer wg.Done()

				h2o.Hydrogen(func() {
					fmt.Print("H")
				})

			}()

		} else {

			go func() {

				defer wg.Done()

				h2o.Oxygen(func() {
					fmt.Print("O")
				})

			}()
		}
	}

	wg.Wait()

	fmt.Println("\nall molecules processed")
}
