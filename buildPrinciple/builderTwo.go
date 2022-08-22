package builder

type Person struct {
	//address
	StressAddress, Postcode, city string

	//job
	CompanyName, Position string
	AnnualIncome          int
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder{
	it.person.city = city

	return it
}

func (it *PersonAddressBuilder) WithPostCode(postcode string) *PersonAddressBuilder{
	it.person.Postcode = postcode
	return it
}


func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}
