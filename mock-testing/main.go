package main

import time

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

func GetPersonByDNI(dni string)(Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM Persona Where ... (Imagina que aca viene una sentencia de este estilo.)
	return Person{}, nil
}

func GetEmployeeById(id int) (Employee, error) {
	time.Sleep.Second
	// SELECT * FROM Employee Where ... (Imagina que aca viene una sentencia de este estilo.)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
	
	e, err := GetEmployeeById(id)
	if error != nil{
		return ftEmployee, err
	}
ftEmployee.Employee = e

p, err := GetPersonByDNI(dni)
if err != nil{
	return ftEmployee, err
}
ftEmployee.Person = p 

return ftEmployee, nil
}


