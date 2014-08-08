package main

import (
	"fmt"
)

func main() {
	//create slice
	s := make([]int, 10)
	fmt.Println(s)
	// init
	for i, _ := range s {
		s[i] = i
	}
	fmt.Println(s)
	// sum all elemets who are even number
	fmt.Println(sum(s, isDouble))
	// sum all elemets who are single number
	fmt.Println(sum(s, isSingle))
}

/**
 callback function : elemet who is even number
**/
func isDouble(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

/**
 callback function : elemet who is single number
**/
func isSingle(a int) bool {
	if a%2 == 0 {
		return false
	} else {
		return true
	}
}

/**
	sum all element who are func number
**/
func sum(s []int, f func(int) bool) int {
	sum := 0
	for _, v := range s {
		if f(v) {
			sum += v
		}
	}
	return sum
}
