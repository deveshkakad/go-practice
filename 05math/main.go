package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Maths in golang")

	//rand.Seed(30) will give me the same number as Seed is fixed with same value everytime

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5) + 1)
}
