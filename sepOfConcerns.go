package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

//SINGLE RESPONSIBILITY PRINCIPLE
/*
The single responsibility principle states that a type should only have one primary responsible as result
it should have only one reason to change, and the reason must be related to its pry responsibility.
*/
//count the no of entries

var entryCount int = 0

type Journal struct {
	entries []string
}

//convert to string
func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

//function to add an entry to the journal and return the position of the entry
func (j *Journal) AddEntry(text string) int {

	entryCount++

	entry := fmt.Sprintf("%d : %s", entryCount, text)

	//add entry to entries
	j.entries = append(j.entries, entry)

	return entryCount

}

//function to remove entry from journal
func (j *Journal) RemoveEntry(index int) {

}

//SEPERATION OF CONCERNS

//function save entries to file
func (j *Journal) SaveToFile(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

//function to load file
func (j *Journal) Loadfile(filename string) {

}

//function to load from the web
func (j *Journal) LoadFromWeb(url *url.URL) {

}

var LineSeperator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeperator)), 0644)
}

type Persistence struct {
	lineSeperator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeperator)), 0644)
}

func main() {

	j := Journal{}

	j.AddEntry("I became more wise today")
	j.AddEntry("I became better in Christ today")
	fmt.Println(j.String())

	SaveToFile(&j, "journal.txt")

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
