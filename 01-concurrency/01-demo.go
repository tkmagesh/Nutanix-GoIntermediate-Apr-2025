package main

import (
	"fmt"
	"time"
)

func main() {
	// panic("dummy panic!")
	go f1() //scheduling the execution of f1() through the scheduler
	f2()

	// block the execution of main() so that the scheduler can look for other goroutines that are scheduled and execute them (cooperative multitasking)

	// DO NOT DO THIS!!
	// poor man's synchronization technique
	time.Sleep(2 * time.Second)

	fmt.Println("Done!")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
