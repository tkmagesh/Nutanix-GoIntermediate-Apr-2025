package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("All the data received!")
}

func genNos(ch chan<- int) {
	count := rand.Intn(20)
	fmt.Printf("[genNos] count : %d\n", count)
	for no := range count {
		ch <- (no + 1) * 10
	}
	close(ch)
}
