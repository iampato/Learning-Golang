package main

import "fmt"

func main() {
	var i int = 0

	for i < 5 {
		fmt.Println("Value of i is: ", i)
		i++
		if i == 3 {
			break
		} else {
			continue
			//fmt.Println("I will start at ",i)
		}
	}
}
