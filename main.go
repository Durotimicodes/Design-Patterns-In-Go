package main

import (
	"fmt"
	builder "github/Durotimicode-design-patterns/buildPrinciple"
)

func main() {

	builder.BuildUnorderList()

	pb := builder.NewPersonBuilder()

	pb.Lives().At("123 Birmingham City").In("London").
		WithPostCode("ABQ124").Works().PlaceOfWork("Victoria Island").
		AsA("Software Engineer").Earning(70000)

	person := pb.BuildPerson()
	fmt.Println(person)

	//notification

}
