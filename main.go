package main

import (
	"fmt"

	_ "github.com/yuanyu90221/gotest/foo"
)

func add(i, j int) int {
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
}

func foo() func() int {
	return func() int {
		return 100
	}
}
func main() {
	fmt.Println(add(1, 2))
	a := 1
	b := 2
	a, b = swap(a, b)
	fmt.Println("a:", a, "b:", b)

	a, b = b, a
	fmt.Println("a:", a, "b:", b)

	bar := foo()
	fmt.Printf("%T\n", bar)
	fmt.Println(bar())

	bar2 := func(i, j float32) float32 {
		return i + j
	}

	fmt.Printf("%T\n", bar2)
	fmt.Println(bar2(1, 45+2.3))
}
