package main

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

type Specification interface { // here you are going to test whether the function satisfies a criteria
	IsSatisfied(p *Product) bool
}

//to check the color
type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

//to check the size specification
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

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")

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
}
