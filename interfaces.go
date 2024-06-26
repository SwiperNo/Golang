package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Samson",
		Breed: "Chihuaha",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Koko",
		Color:         "Red",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animals says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Grunt"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
