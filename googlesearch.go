package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web    = fakeSearch("web")
	Image  = fakeSearch("image")
	Video  = fakeSearch("video")
	Web1   = fakeSearch("web1")
	Image1 = fakeSearch("image1")
	Video1 = fakeSearch("video1")
	Web2   = fakeSearch("web2")
	Image2 = fakeSearch("image2")
	Video2 = fakeSearch("video2")
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google10(query string) (results []Result) {
	//fmt.Println("results = ", results)
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

func Google20(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func Google21(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func Google30(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google10("golang")
	//results := Google20("golang")
	//results := Google21("golang")
	//results := Google30("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// test first function
// func main() {

// 	rand.Seed(time.Now().UnixNano())
// 	start := time.Now()
// 	result := First("golang",
// 		fakeSearch("replica 1"),
// 		fakeSearch("replica 2"))
// 	elapsed := time.Since(start)
// 	fmt.Println(result)
// 	fmt.Println(elapsed)
// }
