package solid

import "fmt"

//Dependency Inversion Principle
//HLM (high level module) should not depend on LLM(low level module)
//low level : in memory db
//high level : functionalities etc
//abstraction : means the interface
//both should depend on abstractions

/*

 */
type Relationship int

const (
	Parent Relationship = iota
	Child
	Silbling
)



type Person struct{
	name string
}

type Info struct {
	from *Person
	relationship Relationship
	to *Person
}

//low-level module
type Relationships struct{
	relations []Info
}


//high-level module
type Research struct{
	//break DIP
	relationships Relationships
}

func (r *Relationships) AddParentAndChild(parent, child *Person){
	r.relations = append(r.relations, Info{parent, Parent, child})

	r.relations = append(r.relations, Info{child, Child, parent})
}

func(r *Research) Investigate(){
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("John has a child called", rel.to.name)
		}
	}
}




