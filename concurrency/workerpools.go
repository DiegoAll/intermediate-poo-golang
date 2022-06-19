package main

import (
	"fmt"
)

func main() {
	// Numbers that the Fibonacci series will be calculated.
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	// We will use this channel as if it were a semaphore.
	jobs := make(chan int, len(tasks))
	// In this channel all results will be transmitted
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		// Create several workers that are attentive.
		go Worker(i, jobs, results)
	}

	// Tasks are iterated over and sent through the jobs channel
	for _, value := range tasks {
		jobs <- value
	}
	// The jobs channel is closed, everything necessary has already been transmitted.
	close(jobs)

	// The results are read
	for r := 0; r < len(tasks); r++ {
		<-results
	}

}

// id: It will help to identify which Worker is working and which is executing.
// jobs:  Here the Worker receives all the jobs that are being assigned. (Valor al que deseo calcularle la serie)
//results: It serves to send which are the values that we have calculated in the Fibonacci series.
func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Workers with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

// Implementation based on recursion
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
