package coreconcepts

import (
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)


// Run N goroutines 
// Cancel all if one fails

func ErrorGroup() {

	var g errgroup.Group

	for i := 1; i <= 5; i++ {

		id := i

		g.Go(func() error {

			if id == 3 {
				return errors.New("worker failed")
			}

			fmt.Println("worker completed:", id)

			return nil
		})
	}

	if err := g.Wait(); err != nil {

		fmt.Println("error:", err)
	}
}