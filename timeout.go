package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func timeout(ch <-chan string) {
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much!")
			return
		}
	}
}

func main() {
	c := boring("goo")
	timeout(c)
	fmt.Println("You're boing; I'm leaving.")
}
