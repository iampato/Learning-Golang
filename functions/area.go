package main

import "fmt"

func calcArea(x int,y int) int {
	return x * y
}
func calcVolume(w,h,b int)int {
	return w * h *b
}

func main(){
	fmt.Println(calcArea(12,10))
	volume := calcVolume(2,10,5)
	fmt.Println("Volume is: ",volume)
}