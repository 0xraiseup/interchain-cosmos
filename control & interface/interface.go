package main

import (
	"fmt"
	"reflect"
)

func PrintType(i interface{}) {
	switch i.(type) {
	case bool:
		fmt.Printf("Value of i is %v\n", i.(bool))
	case int, int32, int64, uint, uint16, uint32, uint64:
		fmt.Printf("Value of i is %v\n", i)
	default:
		fmt.Printf("I has another type\n")
	}
}

func main() {
	var i interface{}
	i = "string"
	PrintType(i)
	i = 34
	fmt.Println(reflect.TypeOf(i))
	PrintType(i)
	i = true
	PrintType(i)
}
