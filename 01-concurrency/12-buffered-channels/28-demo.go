/* channel behavior */

package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 3)
	fmt.Printf("cap(ch) : %d, len(ch) : %d\n", cap(ch), len(ch))
	ch <- 100
	fmt.Printf("cap(ch) : %d, len(ch) : %d\n", cap(ch), len(ch))
	ch <- 200
	fmt.Printf("cap(ch) : %d, len(ch) : %d\n", cap(ch), len(ch))
	ch <- 300
	fmt.Printf("cap(ch) : %d, len(ch) : %d\n", cap(ch), len(ch))
	fmt.Println(<-ch)
	fmt.Printf("cap(ch) : %d, len(ch) : %d\n", cap(ch), len(ch))
	fmt.Println(<-ch)
	fmt.Printf("cap(ch) : %d, len(ch) : %d\n", cap(ch), len(ch))
	fmt.Println(<-ch)
	fmt.Printf("cap(ch) : %d, len(ch) : %d\n", cap(ch), len(ch))

}
