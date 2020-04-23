package main

import (
	"fmt"

	_ "github.com/yuanyu90221/gotest/foo"
)

func main() {
	foo := func() string {
		return "Hello World"
	}
	fmt.Println(foo())
	bar := func() {
		fmt.Println("Hello World 2")
	}
	bar()

	func() {
		fmt.Println("Hello World3")
	}()
}
