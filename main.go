package main

import (
	"fmt"
	builderdesign "github/Durotimicode-design-patterns/builderDesign"
	factoriesdesign "github/Durotimicode-design-patterns/factoriesDesign"
)

func main() {

	builderdesign.BuildUnorderList()

	pb := builderdesign.NewPersonBuilder()

	pb.Lives().At("123 Birmingham City").In("London").
		WithPostCode("ABQ124").Works().PlaceOfWork("Victoria Island").
		AsA("Software Engineer").Earning(70000)

	person := pb.BuildPerson()
	fmt.Println(person)

	//notification

	ff := factoriesdesign.NewPerson("Oluwadurotimi", 31)

	fmt.Println("factory function:", ff)

	p := factoriesdesign.NewPersonInterface("Edmond", 32)
	p.SayHello()
	fmt.Println(p)

	developerFactory := factoriesdesign.NewEmployeeFactory("developer", 60000)
	managerFacyory := factoriesdesign.NewEmployeeFactory("manager", 80000)

	developer := developerFactory("Lydia")
	manager := managerFacyory("Lauretta")

	fmt.Println(developer)
	fmt.Println(manager)

	bossFactory := factoriesdesign.NewEmployeeFactory("CEO", 100000)
	bossFactory.AnnualIncome = 10000
	fmt.Println(&bossFactory)

}
