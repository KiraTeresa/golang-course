package main

import (
	"fmt"
	"runtime"
	"sync"
)

//prevent race conditions by using a mutex

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // lock this, no one else can use it
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() // unlock this and make available for others again
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("count:", counter)
}
