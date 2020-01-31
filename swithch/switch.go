package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, "as")

	/*
		A simple switch statement
	*/
	switch i {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	}

	/*
		You can use commas to separate multiple expressions in the same case statements
		We use the default case as well
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Its a weekday")
	}

	/*
		Switch without an expression is an alternative way to write if/else logic
		Here we also show how case expressions can be non-constants
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
