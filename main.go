package main

import (
	"invest/structs"
	"log"
)

func main() {

	var creditAsigner structs.Assigner
	b3, b5, b7, err := creditAsigner.Assign(6701)

	if err != nil {

		log.Println(err.Error())
	}

	log.Println(b3, b5, b7)

}
