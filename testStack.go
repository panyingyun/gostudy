package main

import (
	"fmt"
	"github.com/panyingyun/gostudy/stacker"
)

func main() {
	a := stacker.NewStack(2)
	a.Push("Google")
	a.Push(100)
	a.Push([]int{1, 2, 3})
	fmt.Println(a)
	fmt.Println(a.Len())
	fmt.Println(a.Cap())

	out, error := a.Pop()
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(out)
	}

	out, error = a.Pop()
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(out)
	}

	out, error = a.Pop()
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(out)
	}

	out, error = a.Pop()
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(out)
	}

}
