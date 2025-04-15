/*
Refactor the below to follow "share memory by communicating"
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var start, end int
	start = 2
	end = 100
	primes := generatePrimes(start, end)
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) []int {
	primes := []int{}
	wg := &sync.WaitGroup{}
	mutex := sync.Mutex{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isPrime(no) {
				mutex.Lock()
				{
					primes = append(primes, no)
				}
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
