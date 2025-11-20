package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Fix the race condition you created in exercise #210 by using package atomic
*/

func main() {
	var wg sync.WaitGroup
	var incrementer int64 = 0

	gs := 15
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("final value:", incrementer)
}
