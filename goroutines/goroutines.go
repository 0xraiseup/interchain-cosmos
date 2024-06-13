package main

import (
	"fmt"
	"time"
)

func doSomething(size int) {
	for i := 0; i < size; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	go doSomething(10)
	go doSomething(10)
	time.Sleep(10 * time.Second)
}
