package main

import (
	"fmt"
)

func main() {
	countA := createCounter(2)
	countB := createCounter(100)
	for i := 0; i < 5; i++ {
		fmt.Printf("(A:%d B:%d)\n", <-countA, <-countB)
	}
}

func createCounter(start int) chan int {
	next := make(chan int)
	go func(i int) {
		for {
			next <- i
			i++
		}
	}(start)
	return next
}
