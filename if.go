package main

import "fmt"

func main() {
	var marks int = 600

	if marks < 0 || marks > 100 {
		fmt.Println("Invalid score")
	} else if marks >= 70 && marks <= 100 {
		fmt.Println("A grade !!")
	} else if marks >= 50 && marks < 70 {
		fmt.Println("B grade !!")
	} else if marks >= 30 && marks < 50 {
		fmt.Println("C grade !!")
	} else if marks >= 0 && marks < 30 {
		fmt.Println("D grade !!")
	} else {
		fmt.Println("Nothing evaluated to true")
	}
}
