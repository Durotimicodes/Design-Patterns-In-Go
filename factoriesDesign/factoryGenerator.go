package factoriesdesign

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

//create factories for specific roles in the company: functional approach
func NewEmployeeFactory(position string, AnnualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, AnnualIncome}
	}
}

//create factories for specific roles in the company:structural approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}
