package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int
	var x = 1
	// var x int = 1
	fmt.Println("Integers: ", strconv.Itoa(x))

	//float 32
	var bal float32 = 100.9
	fmt.Println(bal)

	//float 64
	var bal2 float64 = 100.912345678
	fmt.Println(bal2)

	//converting float632 to float64
	fmt.Println(float64(bal))
	//conveting float32 to float64
	fmt.Println(float32(bal2))

	//constants
	const length int = 10
	const width int = 5
	var area int = 0

	area = length * width
	fmt.Println("area is equal to: ", area)
}
