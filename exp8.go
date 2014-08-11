package main

import (
	"fmt"
)

type city struct {
	name string
}

func (c city) Put(name string) {
	c.name = name
}

func (c city) GetName() string {
	return c.name
}

type country struct {
	name string
}

func (c country) Put(name string) {
	c.name = name
}

func (c country) GetName() string {
	return c.name
}

type IName interface {
	Put(string)
	GetName() string
}

//interface type and query
func printname(p interface{}) {
	ivalue, ok := p.(IName)
	if !ok {
		fmt.Println("It is not a IName interface obj:", p)
		return
	}
	switch ivalue.(type) {
	case *city:
		fmt.Println("It is a *city: ", ivalue.GetName())
	case *country:
		fmt.Println("It is a *country: ", ivalue.GetName())
	case city:
		fmt.Println("It is a city: ", ivalue.GetName())
	case country:
		fmt.Println("It is a country: ", ivalue.GetName())
	default:
		fmt.Println("It is other IName interface")
	}

}

func main() {
	var c1, c2, c3, c4 interface{}
	c1 = city{name: "Hangzhou"}
	c2 = country{name: "US"}
	c3 = &city{name: "Shanghai"}
	c4 = &country{name: "Japan"}
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	//print name of object has interface IName
	printname(c1)
	printname(c2)
	printname(c3)
	printname(c4)
	//print name of object not has interface IName
	printname(10)
	printname("abc")
}
