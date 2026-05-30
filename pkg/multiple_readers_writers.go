package pkg

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Store struct {
	data map[int]string
	mu   sync.RWMutex
}

func Reader_Writer() {

	store := &Store{
		data: make(map[int]string),
	}

	var wg sync.WaitGroup

	numberOfReaders := 5
	numberOfWriters := 3

	// Writers
	for i := 0; i < numberOfWriters; i++ {

		wg.Add(1)

		go writer(
			store,
			i,
			&wg,
		)
	}

	// Readers
	for i := 0; i < numberOfReaders; i++ {

		wg.Add(1)

		go reader(
			store,
			i%3,
			&wg,
		)
	}

	wg.Wait()

	fmt.Println("All operations finished")
}

func writer(
	store *Store,
	key int,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	store.mu.Lock()

	defer store.mu.Unlock()

	value := "value-" + strconv.Itoa(key)

	store.data[key] = value

	fmt.Printf(
		"Writer wrote key=%d value=%s\n",
		key,
		value,
	)

	time.Sleep(time.Second)
}

func reader(
	store *Store,
	key int,
	wg *sync.WaitGroup,
) {

	defer wg.Done()

	time.Sleep(500 * time.Millisecond)

	store.mu.RLock()

	defer store.mu.RUnlock()

	value, ok := store.data[key]

	if ok {

		fmt.Printf(
			"Reader read key=%d value=%s\n",
			key,
			value,
		)

	} else {

		fmt.Printf(
			"Reader read key=%d NOT FOUND\n",
			key,
		)
	}
}
