package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//Declaración de manera explicita
	var x int
	x = 8
	//Declaración de manera implicita
	y := 7.0

	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("10", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	myValue2, err := strconv.ParseInt("NaN", 0, 64)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue2)
	}

	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])

	s := []int{1, 2, 3}

	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	//add a new value into the slice
	s = append(s, 16)
	fmt.Println(s)

	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	//c := make(chan int)
	//go doSomething(c)
	//<-c
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
