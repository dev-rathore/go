package main

import (
	"fmt"
	"sync"
	// "time"
)

// Multithreading
// Concurrency

// With the help of go routines, we can run multiple tasks at the same time, we can write background workers, we can do CPU-intensive tasks concurrently

// func task(id int) {
// 	fmt.Println("Doing task:", id)
// }

// func main() {
	// for i := 0; i <= 10; i++ {
		// go task(i)

		// go func(i int) {
		// 	fmt.Println("Doing task:", i)
		// }(i)
	// }

	// time.Sleep(time.Second * 2)
// }

// WaitGroup

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing task:", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}
