package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

//Recordar que en Go hay que especificar el tipo de retorno.
func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	//fmt.Printf("%+v\n", e)
	e.id = 1
	e.name = "Pedro"
	//e.SetId(5)
	fmt.Printf("%+v\n", e)
	e.SetId(7)
	//fmt.Printf("%+v\n", e)
	e.SetName("Jonas")
	//fmt.Printf("%+v\n", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
}
