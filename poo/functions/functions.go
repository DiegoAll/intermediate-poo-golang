package main

import (
	"fmt"
	"time"
)

// Example 1

func main() {
	/*
		x := 5
		y := func() int {
			return x * 2
		}()
		fmt.Println(y)
		z := 10
		w := func() int {
			return z * 2
		}()
		fmt.Println(w)
		//It's a bad practice, anti pattern!! Dry (Don`t repeat yourself)
	*/
	c := make(chan int)
	go func() {
		fmt.Println("Starting Function")
		time.Sleep(5 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}
