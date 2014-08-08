package main

import (
	"fmt"
	"github.com/panyingyun/gostudy/algorithms"
)

func main() {
	s := []int{10, 5, 6, 3}
	fmt.Println(s)
	bubble.Sort(s)
	fmt.Println(s)
}
