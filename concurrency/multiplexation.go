package main

import (
	"fmt"
	"time"
)

// i: Time to sleep to DoSomething().
// c: A channel is used to send this data.
// param: This parameter will change depending on which channel is being used.
func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	// Finally we are going to send the param through our channel
	c <- param
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)

		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}
