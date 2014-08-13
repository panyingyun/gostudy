package main

import (
	"fmt"
	"time"
)

var c chan interface{}

func main() {
	c = make(chan interface{})
	go ready("Hangzhou", 2)
	//go ready("Shanghai", 1)
	//fmt.Println("waiting ...")
	//<-c
	//<-c
	timeout := time.After(5 * time.Second)
	for {
		select {
		case <-c:
			fmt.Println("OK")
		case <-timeout:
			fmt.Println("time is out")
			return
		}
	}
}

func ready(w string, sec int64) {
	time.Sleep(time.Duration(sec * 1e9))
	fmt.Println(w, "is ready!")
	c <- 1
}
