package main

import (
	builder "github/Durotimicode-design-patterns/buildPrinciple"
)

func main() {

	builder.BuildUnorderList()

	rc := &Rectangle{2, 3}
	rg :=
		UseIt(rc)

	parent := Person{"John"}
	child1 := Person{"Edmond"}
	child2 := Person{"Mary"}

	//low level
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{relationships}
	r.Investigate()


	//

}
