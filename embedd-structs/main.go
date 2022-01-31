package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
	age   int
	home  address
}

type address struct {
	country string
	zipCode string
}

func main() {

	engineers := make([]person, 0, 5)
	jcama := person{}
	jcama.fName = "Juan"
	jcama.lName = "Camacho"
	jcama.age = 26
	jcama.home.country = "Mexico"
	jcama.home.zipCode = "07040"
	engineers = append(engineers, jcama)

	fbar := person{}
	fbar.fName = "Fernando"
	fbar.lName = "Barrera"
	fbar.age = 25
	fbar.home.country = "Mexico"
	fbar.home.zipCode = "07040"

	engineers = append(engineers, fbar)
	fmt.Println("Total of Engineers: ", len(engineers), engineers)

}
