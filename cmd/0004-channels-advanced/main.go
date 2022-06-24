package main

import (
	"fmt"
	"time"
)

type FooIn struct {
	a, b int
}

func producer(idx int, in <-chan *FooIn, out chan<- int) {
	for {
		x := <-in
		fmt.Printf("Producer[%d]: received %v\n", idx, x)
		out <- (x.a + x.b)
	}
}

func consumer(in <-chan int) {
	for {
		fmt.Println("Consumer: received ", <-in)
	}
}

func main() {
	inputs := make(chan *FooIn)
	results := make(chan int)

	for i := 0; i < 10; i++ {
		idx := i
		go producer(idx, inputs, results)
	}
	go consumer(results)

	for i := 0; i < 100; i++ {
		inputs <- &FooIn{3, i}
	}
	time.Sleep(5 * time.Second)
}
