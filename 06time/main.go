package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Golang time package practice")
	presentTime := time.Now()
	//fmt.Println("Present time:", presentTime)
	fmt.Println("Format data:", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(1999, time.May, 19, 5, 0, 0, 0, time.UTC)
	fmt.Println("createdDate:", createdDate)

	fmt.Println("Create Time format:", createdDate.Format("01-02-2006 Monday"))
}
