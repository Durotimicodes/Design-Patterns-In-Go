package solid

import "fmt"
//Open closed principle OCP
/*
This states that a type should be open for extension, but closed for modification
bascially you want to be able to extend the code/function
*/

//Specification pattern, decribes the OCP well

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

//THE ABOVE CODE WILL VOILATE THE OPEN CLOSED PRINCIPLE when your boss askes to modify the code
//to retify this we use the Specification Principle of the interface
//a composite specification if implemented when you are required to add two or more specification to
//you pre existing function/method

type Specification interface { // here you are going to test whether the function satisfies a criteria
	IsSatisfied(p *Product) bool
}

//create a color specification
type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

//create a size specification
type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct{}

func (bf *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

//composite specification
type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}


var apple = Product{"Apple", green, small}
var tree = Product{"Tree", green, large}
var house = Product{"House", blue, large}

var products = []Product{apple, tree, house}

f := Filter{}

for _, v := range f.FilterByColor(products, blue) {
	fmt.Printf("- %s is blue\n", v.name)
}

fmt.Printf("Green products (new): \n")

greenSpec := ColorSpecification{green}

bf := BetterFilter{}
for _, v := range bf.Filter(products, greenSpec) {
	fmt.Printf("- %s is green\n", v.name)
}

largeSpec := SizeSpecification{large}
lgSpec := AndSpecification{greenSpec, largeSpec}

fmt.Printf("Large green products:\n")

for _, v := range bf.Filter(products, lgSpec) {
	fmt.Printf("-%s is green\n", v.name)
}