package main

import (
	"fmt"
	"time"
)

func doSomething(size int, c chan int) {
	for i := 0; i < size; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	c <- size
}

func main() {
	c, q := make(chan int), make(chan int)
	jobs := 5

	go func() {
		for i := 1; i <= jobs; i++ {
			doSomething(i*10, c)
		}
		q <- 0
	}()

	for {
		select {
		case x := <-c:
			fmt.Println(x)
		case <-q:
			fmt.Println("Finished")
			return
		default:
			fmt.Print("...")
			time.Sleep(time.Millisecond)
		}
	}
}
