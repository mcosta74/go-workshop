package main

import (
	"fmt"
	"time"
)

func foo(a int, res chan<- int) {
	time.Sleep(time.Second) // emulate complex computation
	res <- (a * a)
}

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Bye")

	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(a int) {
			foo(a, ch)
		}(i)
	}

	var done int
	squares := make([]int, 0)
	for {
		res := <-ch
		squares = append(squares, res)
		done++

		if done == 10 {
			break
		}
	}
	fmt.Println("List of squares:", squares)
}
