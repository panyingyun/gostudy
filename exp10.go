package main

import (
	"fmt"
	"runtime"
	"time"
)

func sum(a []int64, c chan int64, flag string) {
	//fmt.Println("a = ", a)
	fmt.Println("time = ", time.Now(), flag)
	var sum int64 = 0
	//sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println("sum = ", sum, flag)
	fmt.Println("time = ", time.Now(), flag)
	c <- sum // send sum to c
}

// func loop() {
// 	for {
// 	}
// }

func main() {
	tStart := time.Now()
	runtime.GOMAXPROCS(4)
	// number := runtime.NumCPU()
	// fmt.Println("CPU number = ", number)

	fmt.Println("Gorountine number = ", runtime.NumGoroutine())
	//
	s := make([]int64, 100000000)
	for index, _ := range s {
		s[index] = int64(index * 2)
	}

	c1 := make(chan int64)
	c2 := make(chan int64)

	go sum(s[:len(s)/2], c1, "Google")
	go sum(s[len(s)/2:], c2, "Apple")
	//go loop()
	fmt.Println("Gorountine number = ", runtime.NumGoroutine())

	sum1 := <-c1
	fmt.Println(sum1)
	sum2 := <-c2
	fmt.Println(sum2)
	fmt.Println(sum1 + sum2)
	tEnd := time.Now()
	fmt.Println("total exec time = ", tEnd.Sub(tStart))
}
