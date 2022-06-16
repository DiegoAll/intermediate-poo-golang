package main

import "fmt"

func main() {

	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 6
	c <- 9

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
