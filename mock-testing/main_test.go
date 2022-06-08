package main

import testing

func TestGetFullTimeEmployeeById(t *testing.T){
	table := []struct{
		id int
		dni string
		mockFunc func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id: 1,
			dni: "1",
			mockFunc: func(){
				GetEmployeeById = func(id int) (Employee, error){
					return Employee{
						Id: 1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error){
					//De manera explicita devuelve la persona que yo requiero
					return Person{
						DNI: ,
						Name: "John Doe",
						Age: 35,
					}
				}, nil
			},
		}
	}
}