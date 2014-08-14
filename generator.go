package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
golang partens:
generator: function that returns a channel
*/
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	// ch := boring("google")
	// fmt.Println("I'm listening.")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("receive str:", <-ch)
	// }

	google := boring("google")
	apple := boring("apple")
	for i := 0; i < 5; i++ {
		fmt.Println(<-google)
		fmt.Println(<-apple)
	}

	fmt.Println("You're boing; I'm leaving.")
}
