package main

import (
	"fmt"
)

type City struct {
	name string
	code int
}

func main() {

	// pointer
	i := 10
	p := &i
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(*p)
	*p = 20
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(*p)

	//new var is diff
	var c1 City
	c2 := new(City)
	fmt.Println(c1) //it is city
	fmt.Println(c2) //it is address of city

	//init
	c3 := City{name: "hangzhou", code: 310000}
	c4 := new(City)
	c4.name = "shanghai"
	c4.code = 330000
	fmt.Println(c3) //it is city
	fmt.Println(c4) //it is pointer of city

	//make only create slice 、map 、channel
	//there are not pointer
	s := make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))
	m := make(map[int]string)
	m[10] = "Google"
	m[19] = "Microsoft"
	fmt.Println(m)

}
