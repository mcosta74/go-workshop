package main

import (
	"fmt"
)

func foo(a int, done chan<- struct{}) {
	fmt.Println("Hello ", a)
	done <- struct{}{}
}

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Bye")

	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(a int) {
			foo(a, ch)
		}(i)
	}

	var done int
	for {
		<-ch
		done++

		if done == 10 {
			break
		}
	}
}
