package main

/*
The interface segragation principle states that you should not put too much into an interface
sometimes it makessense to break the interface

basically try to break up an interface into seperate parts that people will definitely need.

*/

type Document struct{}



type Machine interface{
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{

}

func (m MultiFunctionPrinter) Print(d Document){

}

func (m MultiFunctionPrinter) Scan(d Document){

}
func (m MultiFunctionPrinter) Fax(d Document){

}


//an old fashioned printer
type OldFashionPrinter struct{
}

func (o OldFashionPrinter) Print(d Document){

}

func (o OldFashionPrinter) Fax(d Document){
	panic("Operatio not supported")
}

func (o OldFashionPrinter) Scan(d Document){

}

//ISP
type Printer interface{
 Printer(d Document)
}

type Scanner interface{
	Scan(d Document)
}

type MyPrinter struct{

}

func (m MyPrinter) Print(d Document){

}

type PhotoCopier struct{

}

func (p PhotoCopier) Scan(d Document){

}

func(p PhotoCopier) Print(d Document){

}

//an interface that represent a multi function device
type MultiFunctionDevice interface{
	Printer
	Scanner
}

type MultiFunctionMachine struct{
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document){
	
}


func main(){

}