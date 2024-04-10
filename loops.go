package main

import "log"

func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "cow", "cat", "mario"}

	for _, animal := range animals {
		log.Println(animal)
	}

	animals_map := make(map[string]string)
	animals_map["dog"] = "Fido"
	animals_map["cat"] = "PussInBoots"

	for animalType, animals_map := range animals_map {
		log.Println(animalType, animals_map)
	}

	var firstLine = "Once upon a mudnight dreary"
	firstLine = "x"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"odejon", "Zuniga", "jacob@wonder.com", 100})
	users = append(users, User{"Kyle", "Jones", "jacob@wonder.com", 30})
	users = append(users, User{"Vivian", "Jones-Zuniga", "jacob@wonder.com", 30})
	users = append(users, User{"koko", "Jones-Zuniga", "jacob@wonder.com", 15})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
