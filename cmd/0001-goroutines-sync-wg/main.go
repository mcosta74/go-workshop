package main

import (
	"fmt"
	"sync"
)

func foo(a int) {
	fmt.Println("Hello ", a)
}

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Bye")

	var g sync.WaitGroup

	for i := 0; i < 10; i++ {
		g.Add(1)
		go func(a int) {
			defer g.Done()
			foo(a)
		}(i)
	}
	g.Wait()
}
