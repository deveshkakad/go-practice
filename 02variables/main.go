package main

import "fmt"

// rollno := 25 This is invalid syntax,to declare and initialize global variables ,need to use var keyword

//var rollno = 25 Valid sysntax
//svar rollno int =25 Valid Syntax

func main() {
	var username string = "Devesh"
	fmt.Println(username)
	fmt.Printf("Type :%T \n", username)

	var isEnabled bool = true
	fmt.Println(isEnabled)
	fmt.Printf("Type:%T \n", isEnabled)
}
