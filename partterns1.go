package main

import (
	"fmt"
	//"math/rand"
	"runtime"
	"time"
)

func boring(msg string, c chan<- string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
		//time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
func main() {
	tStart := time.Now()
	runtime.GOMAXPROCS(4)
	ch := make(chan string, 10)
	go boring("google", ch)
	fmt.Println("I'm listening.")
	for i := 0; i < 5; i++ {
		fmt.Println("receive str:", <-ch)
	}
	fmt.Println("You're boing; I'm leaving. ", time.Now().Sub(tStart))
}
