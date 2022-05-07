package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person   //Campo anónimo
	Employee //Campo anónimo
	endDate  string
}

func (ftEmployee FullTimeEmployee) GetMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) GetMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	GetMessage() string
}

// Receives interface as parameter Target: Polymorphism
// Here the interface is being implemented
func getMessage(p PrintInfo) {
	fmt.Println(p.GetMessage())
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Sujeto"
	ftEmployee.age = 10
	ftEmployee.id = 55555555
	fmt.Printf("%v\n", ftEmployee)
	// Error because in Golang it's not possible apply polymorphism without interfaces.
	// GetMessage(ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
