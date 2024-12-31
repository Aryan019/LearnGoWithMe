package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time Study of Golang ")

	// Package to deal with time is time
	presentTime := time.Now()
	// fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
}
