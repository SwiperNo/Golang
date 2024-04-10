package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	log.Default().Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Default().Println(myMap2["First"])
	log.Println(myMap2["Second"])

	myMap3 := make(map[string]User)

	me := User{
		FirstName: "Kyle",
		LastName:  "Jones",
	}

	myMap3["me"] = me

	log.Println(myMap3["me"].FirstName, myMap3["me"].LastName)

	var mySlice []string

	mySlice = append(mySlice, "kyle")
	mySlice = append(mySlice, "Jones")
	mySlice = append(mySlice, "Is using")
	mySlice = append(mySlice, "a slice")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[4:8])

	names := []string{"one", "two", "seven", "four"}

	log.Println(names[1:3])

}
