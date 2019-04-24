package main

import "fmt"

type DataRetrival interface {
	getData() *Customer
}

type Customer struct {
	Id   int
	name string
}

func (Customer) getData() *Customer {
	return &Customer{Id: 1, name: "Priyank"}
}

func main() {
	var name DataRetrival = Customer{Id: 2, name: "Shah"}
	fmt.Println(name.getData().name)
}
