package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// root context (non cancellable)
	rootCtx := context.Background()

	// context to share data across context hierarchy
	valCtx := context.WithValue(rootCtx, "root-key", "root-value")

	// cancellable context
	timeoutCtx, cancelFn := context.WithTimeout(valCtx, 5*time.Second)
	fmt.Println("Will timeout after 5 seconds, but Hit ENTER to stop anytime if needed....")

	go func() {
		fmt.Scanln()
		// send the cancellation signal by "unblocking" the context.Done() channel
		cancelFn()
	}()
	doneCh := generateData(timeoutCtx)
	<-doneCh
	switch timeoutCtx.Err() {
	case context.Canceled:
		fmt.Println("Programmatic cancellation!")
	case context.DeadlineExceeded:
		fmt.Println("Timeout based cancellation!")
	}
	fmt.Println("Done")
}

func generateData(ctx context.Context) <-chan struct{} {
	doneCh := make(chan struct{})
	fmt.Println("[generateData] root-key :", ctx.Value("root-key"))

	// a context to add more data
	parentValCtx := context.WithValue(ctx, "parent-key", "parent-value")
	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(1)
		evenTimeoutCtx, cancelFn := context.WithTimeout(parentValCtx, 2*time.Second)
		defer cancelFn()
		go generateEven(evenTimeoutCtx, wg)

		wg.Add(1)
		oddTimeoutCtx, cancelFn := context.WithTimeout(parentValCtx, 4*time.Second)
		defer cancelFn()
		go generateOdd(oddTimeoutCtx, wg)
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
		wg.Wait()
		close(doneCh)
	}()
	return doneCh
}

func generateEven(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[generateEven] root-key :", ctx.Value("root-key"))
	fmt.Println("[generateEven] parent-key :", ctx.Value("parent-key"))
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

func generateOdd(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[generateOdd] root-key :", ctx.Value("root-key"))
	fmt.Println("[generateOdd] parent-key :", ctx.Value("parent-key"))
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
