// using mutex

package main

import (
	"fmt"
	"sync"
)

/* Encapsulating the state and "concurrent safe state manipulation" logic together */
type SafeCounter struct {
	count int
	sync.Mutex
}

func (sc *SafeCounter) Add(delta int) {
	sc.Lock()
	{
		sc.count += delta
	}
	sc.Unlock()
}

var sc SafeCounter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", sc.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	sc.Add(1)

}
