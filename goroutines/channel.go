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
	c := make(chan int)
	go doSomething(10, c)
	go doSomething(20, c)
	go doSomething(30, c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z)
}
