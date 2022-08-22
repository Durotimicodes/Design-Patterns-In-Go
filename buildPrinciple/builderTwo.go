package builder

/*


 */
type Person struct {
	//address
	StressAddress, Postcode, city string

	//job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

//address builder
type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.city = city

	return b
}

func (b *PersonAddressBuilder) At(location string) *PersonAddressBuilder {
	b.person.StressAddress = location

	return b
}

func (b *PersonAddressBuilder) WithPostCode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

//Job builder
type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (jb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	jb.person.AnnualIncome = annualIncome

	return jb
}

func (jb *PersonJobBuilder) PlaceOfWork(companyName string) *PersonJobBuilder {
	jb.person.CompanyName = companyName

	return jb
}

func (jb *PersonJobBuilder) AsA(occupation string) *PersonJobBuilder {
	jb.person.Position = occupation

	return jb
}

//person
func (b *PersonBuilder) BuildPerson() *Person {
	return b.person
}
