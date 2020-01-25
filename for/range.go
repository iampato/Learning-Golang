package main

import "fmt"

func main() {
	var mylist = []int{0, 1, 2, 3, 4, 5}

	for i := range mylist {
		fmt.Println(mylist[i])
	}
}
