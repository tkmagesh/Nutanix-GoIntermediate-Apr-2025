/*
Refactor the below to follow "share memory by communicating"
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	primesCh := generatePrimes()
	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

// refactor the following so that the prime generation stops after 10 seconds
func generatePrimes() <-chan int {
	primesCh := make(chan int)
	// timeoutCh := timeout(10 * time.Second)
	timeoutCh := time.After(10 * time.Second)
	go func() {
	LOOP:
		for no := 2; ; no++ {
			select {
			case <-timeoutCh:
				fmt.Println("Timeout occurred")
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
