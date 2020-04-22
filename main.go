package main

import (
	"fmt"

	_ "github.com/yuanyu90221/gotest/foo"
)

const (
	Monday = iota
	Tuesday
	Wednesday
)

func main() {
	foo := "Hello"
	bar := 10
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	// fmt.Println(bar.HelloWorld())
	fmt.Println("Hello Main")
}
