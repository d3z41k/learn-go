package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)

	ch1 <- 9

	select {
		case val := <-ch1:
			fmt.Println("chan-1 val", val)
		case ch2 <- 1:
			fmt.Println("chan-1 val to chan-2")
		default:
			fmt.Println("default case")

	}
}
