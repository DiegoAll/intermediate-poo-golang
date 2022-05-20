package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	return
}

func getCalculadora(x, y int) (suma int, resta int, multiplicacion int) {
	suma = x + y
	resta = x - y
	multiplicacion = x * y
	return
}

//Improvement: returns error when the function don't receives parameters.
func enhancedSum(values ...int) (int, error) {
	if len(values) > 0 {
		total := 0
		for _, num := range values {
			total += num
		}
		return total, nil
	} else {
		return 0, fmt.Errorf("This function receives at least one argument")
	}
}

func main() {
	// fmt.Println(sum(7))
	// fmt.Println(sum(7, 2, 1))
	// fmt.Println(sum(7, 2, 6, 3, 1))
	// printNames("Alice", "Bob", "Charles", "Dave")
	fmt.Println(sum())
	//fmt.Println(getValues(2))
	//println(getCalculadora(7, 5))
	fmt.Println(enhancedSum(3, 6, 1))
	fmt.Println(enhancedSum())
}
