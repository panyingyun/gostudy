package main

import (
	"fmt"
)

//filter并发模式的实现

//filter1 input nums here one by one
func filter1(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

//fileter2 get input and out = input*input one by one
//fileter2 具有相同的输入流和输出流，所以它可以调用N次

func filter2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// out := filter2(filter2(filter1(1, 2, 3, 4, 5, 6)))
	// for result := range out {
	// 	fmt.Println("result = ", result)
	// }
	done := make(chan struct{})
	defer close(done)

}
