package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to out Ratings app")
	fmt.Println("Please enter rating from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	//Need to add 1 to the input

	//numAdded := input +1 This will give us error as input is a string and we are trying to add integer to the string
	//numAdded, err := strconv.ParseInt(input, 10, 64)
	//The Above format will give us error as Error is: strconv.ParseInt: parsing "3\n": invalid syntax
	//This is because as we are entering the rating from the keyboard and entering the enter it is taking \n character in the end

	numAdded, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64) //TrimSpace func will remove the \n character added in the end
	if err != nil {
		fmt.Println("Error is:", err)
	} else {
		fmt.Println("Added 1 to input and result is:", numAdded+1)
	}
}
