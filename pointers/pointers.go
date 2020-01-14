package main

import "fmt"

func main() {
	var n int = 200
	var np *int

	np = &n
	fmt.Printf("The address of n is : %x\n", &n)
	fmt.Printf("The address of np is : %x\n", np)
	fmt.Printf("the value of np is : %d", *np)
}
