/*
Refactor the below to follow "share memory by communicating"
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	primesCh := generatePrimes(stopCh)
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

// refactor the following so that the prime generation stops after 10 seconds
func generatePrimes(stopCh <-chan struct{}) <-chan int {
	primesCh := make(chan int)
	go func() {
	LOOP:
		for no := 2; ; no++ {
			select {
			case <-stopCh:
				fmt.Println("stop signal received")
				break LOOP
			default:
				if isPrime(no) {
					primesCh <- no
					time.Sleep(500 * time.Millisecond)
				}
			}
		}
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
