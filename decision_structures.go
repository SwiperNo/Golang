package main

import "log"

func main() {

	var isTrue bool

	isTrue = false

	if isTrue {

		log.Println("isTrue is", isTrue)

	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("Cat is cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 100

	if myNum > 99 && !isTrue {
		log.Println("my Number is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue {
		log.Println("We got a different vaule this time!")
	} else if myNum == 101 || isTrue {
		log.Println(" It is 101")
	} else if myNum > 1000 && isTrue == false {
		log.Println("We are just getting silly now")
	}

	myVar := "dog"
	switch myVar {
	case "cat":
		log.Println("dog is set to cat")

	case "dog":
		log.Println("dog is set to dog")

	case "fish":
		log.Println("dog is set to fish")

	default:
		log.Println("Unknow value")
	}

}
