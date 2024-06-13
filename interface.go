package main

import "fmt"

func main() {
	var i interface{}
	i = 2
	fmt.Println(i)
	i = "Test"
	fmt.Println(i)
	s, ok := i.(string)
	fmt.Println(s, ok)
}
