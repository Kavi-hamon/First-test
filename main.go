package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("")
	a := willgiveFunc()
	var b any
	b = a()
	fmt.Println(reflect.ValueOf(b).Kind())
	var c int
	c = b.(int) 
	fmt.Println(c)
}

func willgiveFunc() func() any {
	return return2
}

func return2() any {
	fmt.Println("2 called")
	return 2
}
