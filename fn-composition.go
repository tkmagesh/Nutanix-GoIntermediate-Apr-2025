package main

import (
	"fmt"
	"log"
)

func main() {
	logAdd := logWrapper(add)
	logAdd(100, 200)
}

func logWrapper(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("An operation is performed")
		op(x, y)
	}
}

func add(x, y int) {
	fmt.Println("Add Result : ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result : ", x-y)
}
