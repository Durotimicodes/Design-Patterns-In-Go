package factoriesdesign

import "fmt"

type PersonTwo interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, i am %d years old\n", p.name, p.age)
}

func NewPersonInterface(name string, age int) PersonTwo {
	return &person{name, age}
}
