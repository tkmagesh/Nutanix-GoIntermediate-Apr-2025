package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for range 5 {
		fmt.Println(<-ch)
	}
}

func genNos(ch chan<- int) {
	for no := range 5 {
		time.Sleep(500 * time.Millisecond)
		ch <- (no + 1) * 10
	}
}
