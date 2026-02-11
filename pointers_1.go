package main

import "fmt"

func main() {
	x := 10
	p := &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Pointer p stores:", p)
	fmt.Println("Value at address p:", *p)
}
