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
	primesCh := generatePrimes(start, end)
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) <-chan int {
	primesCh := make(chan int)
	nosCh := make(chan int)
	workerCount := 5
	go func() {
		for no := start; no <= end; no++ {
			nosCh <- no
		}
		close(nosCh)
	}()
	go func() {
		wg := &sync.WaitGroup{}
		for workerId := range workerCount {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for no := range nosCh {
					fmt.Printf("Worker - %d, processing %d\n", workerId, no)
					if isPrime(no) {
						primesCh <- no
					}
				}
			}()
		}
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
