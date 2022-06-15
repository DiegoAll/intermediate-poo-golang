package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					// De manera explicita devuelve la persona que yo requiero
					return Person{
						Name: "John Doe",
						Age:  35,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  35,
					DNI:  "1",
					Name: "John Doe",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}
	// Save the original functions
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error , got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDni
}

// Tener en cuenta las comas de mas al cerrar las llaves en el repositorio commit despues del cambio (mock)
