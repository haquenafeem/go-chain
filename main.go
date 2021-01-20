package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (person Person) PersonIn() Person {
	return person
}

func (person Person) PrintName() Person {
	fmt.Println("Name : ", person.Name)
	return person
}

func (person Person) PrintAge() {
	fmt.Println("Age : ", person.Age)
}

func main() {
	person := Person{Name: "Tony Stark", Age: 42}
	person.PersonIn().PrintName().PrintAge()
}
