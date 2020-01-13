package main

import "fmt"

func main(){
	var age =18;

	switch age {
	  case 18:
		  fmt.Println("You are too young ");
	  case 30:
		   fmt.Println("Work hard for a promotion ");
	  case 45:
		  fmt.Println("Hope you have started to save");
	  case 55:
		  fmt.Println("You should retire tomorrow");
	  default: 
		  fmt.Println("Sorry wewe ni anestor");
	}
}