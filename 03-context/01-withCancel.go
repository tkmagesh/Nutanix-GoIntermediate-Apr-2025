package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// root context (non cancellable)
	rootCtx := context.Background()

	// cancellable context
	cancelCtx, cancelFn := context.WithCancel(rootCtx)
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		// send the cancellation signal by "unblocking" the context.Done() channel
		cancelFn()
	}()
	doneCh := generateData(cancelCtx)
	<-doneCh
	fmt.Println("Done")
}

func generateData(ctx context.Context) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		go generateEven(ctx)
		go generateOdd(ctx)
	LOOP:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[generateData] cancellation signal received")
				break LOOP
			default:
				time.Sleep(1 * time.Second)
				fmt.Printf("Current Time : %v\n", time.Now())
			}

		}
		close(doneCh)
	}()
	return doneCh
}

func generateEven(ctx context.Context) {
LOOP:
	for evenNo := 0; ; evenNo += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("[generateEven] cancellation signal received")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Even : %d\n", evenNo)
		}

	}
}

func generateOdd(ctx context.Context) {
LOOP:
	for oddNo := 1; ; oddNo += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("[generateOdd] cancellation signal received")
			break LOOP
		default:
			time.Sleep(700 * time.Millisecond)
			fmt.Printf("Odd : %d\n", oddNo)
		}
	}
}
