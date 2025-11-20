package main

import (
	"fmt"
	"sync"
)

/*
Fix the race condition you created in the previous exercise by using a mutex
	● it makes sense to remove runtime.Gosched()
*/

func main() {
	var wg sync.WaitGroup
	incrementer := 0
	gs := 15
	wg.Add(gs)

	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incrementer)
}
