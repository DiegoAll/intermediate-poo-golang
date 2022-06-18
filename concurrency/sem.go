package main

import (
	"fmt"
	"sync"
	"time"
)

// c := [][][][][]
// c := [goRoutine1][goRoutine2]
// c:= [goRoutine3][goRoutine2]

func main() {
	c := make(chan int, 5)
	var wg sync.WaitGroup
	for i := 0; i < 100000000000; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	// Set the program in the blocking state until the counter reaches 0
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("id %d finished\n", i)
	// Send values through the previously created channel.
	<-c
}
