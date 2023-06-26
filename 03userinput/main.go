package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the Name:")

	name, _ := reader.ReadString('\n')
	fmt.Println("Name:", name)

	fmt.Printf("Enter the Age:")
	age, _ := reader.ReadString('\n')
	fmt.Println("age:", age)
	fmt.Printf("Type of age is :%T \n", age) // Type will come as string
}
