package main

import (
	"fmt"
)

func main() {
	//Println Printf
	fmt.Println("Golang, I am Michael.Pan!")
	i := 101
	fmt.Printf("i = %v \n", i)
	//var int string bool
	var cnt = 100
	str := "hangzhou,China"
	substr := str[1:]
	var isSuccess = true

	fmt.Printf("cnt = %v str = %v substr= %v isSuccess = %v \n", cnt, str, substr, isSuccess)
	//Const
	const MAXCOUNT = 30
	const NAME = "China"
	fmt.Printf("MAXCOUNT = %v , NAME = %v\n", MAXCOUNT, NAME)
	//calc and var +
	var j = 101
	var k int
	k = j + 19
	fmt.Printf("k = %v \n", k)
	//array
	var a [10]int
	a[0] = 10
	a[1] = 11
	fmt.Printf("a = %v \n", a)
	// array foreach
	for index, value := range a {
		fmt.Printf("index = %v, value = %v\n", index, value)
	}
	//slice add remove
	b := []int{1, 2, 3, 4, 5, 6}
	b = append(b, 7, 8, 9, 10, 11, 12)
	c := make([]int, 10, 20)
	copy(c, b)
	fmt.Printf("b = %v \nc = %v len(c) = %v cap(c) = %v\n", b, c, len(c), cap(c))
	//slice foreach
	for index, value := range c {
		fmt.Printf("index = %v, value = %v\n", index, value)
	}
	//map
	m := make(map[string]string)
	m["a"] = "China"
	m["b"] = "American"
	m["c"] = "Brizer"
	m["d"] = "Japan"
	fmt.Printf("m(1) = %v\n", m)
	delete(m, "c")
	fmt.Printf("m(2) = %v\n", m)
	value1, ok1 := m["c"]
	fmt.Printf("value1 = %v ok1 = %v\n", value1, ok1)
	value2, ok2 := m["d"]
	fmt.Printf("value2 = %v ok2 = %v\n", value2, ok2)
	for key, value := range m {
		fmt.Printf("key = %v value = %v\n", key, value)
	}

	//for exp1,2,3
	fmt.Printf("sum1 = %v,sum2 = %v,sum3 = %v\n", sum1(), sum2(), sum3())
	//switch
	fmt.Printf("str0 = %v\n", getStr(0))
	fmt.Printf("str1 = %v\n", getStr(1))
	fmt.Printf("str2 = %v\n", getStr(2))
	fmt.Printf("str3 = %v\n", getStr(3))
	fmt.Printf("str4 = %v\n", getStr(4))
	fmt.Printf("str5 = %v\n", getStr(5))
	fmt.Printf("str6 = %v\n", getStr(6))
	fmt.Printf("str6 = %v\n", getStr(100000))

	// slice append
	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:2]
	slice = append(slice, 6, 7, 8)
	fmt.Println(slice)
	fmt.Println(arr)

	// function
	fmt.Println(quicksum1(11, 12))
	fmt.Println(quicksum2(1, 2, 3, 4, 5, 6))

	// callback
	fmt.Println(quickcmp(10, 12, compare))
	fmt.Println(quickcmp(32, 12, compare))
	fmt.Println(quickcmp(12, 12, compare))

	//closure
	f1 := plus(10)
	f2 := plus(15)
	fmt.Println(f1(20))
	fmt.Println(f2(25))
}

func sum1() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func sum2() int {
	sum := 0
	i := 0
	for i < 10 {
		sum += i
		i++
	}
	return sum
}

func sum3() int {
	sum := 0
	i := 0
	for {
		sum += i
		i++
		if i == 10 {
			break
		}
	}
	return sum
}

func getStr(a int) string {
	var str string
	switch a {
	case 0:
		str = "China"
	case 1:
		str = "American"
	case 2, 3, 4, 5:
		str = "Japan"
	default:
		str = "others"
	}
	return str
}

func quicksum1(a int, b int) (sum int) {
	sum = a + b
	return
}

func quicksum2(a ...int) (sum int) {
	for _, value := range a {
		sum += value
	}
	return
}

func compare(a int, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func quickcmp(a int, b int, f func(int, int) int) (result int) {
	result = f(a, b)
	return
}

func plus(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
