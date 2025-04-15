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
	end = 200
	primesCh := generatePrimes(start, end)
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) <-chan int {
	nosCh := generateData(start, end)
	primesCh := startWorkers(5, nosCh)
	return primesCh
}

func startWorkers(workerCount int, nosCh <-chan int) <-chan int {
	primesCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for range workerCount {
			wg.Add(1)
			go worker(nosCh, primesCh, wg)
		}
		wg.Wait()
		close(primesCh)
	}()
	return primesCh
}

func worker(nosCh <-chan int, primesCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for no := range nosCh {
		if isPrime(no) {
			primesCh <- no
		}
	}
}

func generateData(start, end int) <-chan int {
	nosCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			nosCh <- no
		}
		close(nosCh)
	}()
	return nosCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
