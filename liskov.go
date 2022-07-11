package main

import "fmt"

//Liskov Substitution Principle : primary deals with inheritance but inheritance isnt implemented in Go
/*
Liskov principle states that if you have some apI that takes a base class and works correctly with a base class
it should also work correctly with a derived class // but in go no base class or derived class

The Liskov substitution principle states that a standerdized function implemented should always work
in the general case without breaking  
*/

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
  r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square{
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected an area of", expectedArea, ", but got", actualArea)
}

func main(){

	rc := &Rectangle{2, 3}
	UseIt(rc)
}