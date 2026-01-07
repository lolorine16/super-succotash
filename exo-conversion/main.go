package main

import (
	"fmt"
	"strconv"
)

func main() {
	// age en string
	var age string
	var dadAge int
	fmt.Println("Entrer votre age : ")
	fmt.Scanln(&age)
	fmt.Println("Entrer l'age de votre papa : ")
	fmt.Scanln(&dadAge)

	// conversion de l'age string en int
	n, err := strconv.Atoi(age)
	// gestion d'erreur
	if err != nil {
		fmt.Println("Invalid Input") // entrer invalide
	}

	dadAgeOld := dadAge - n
	fmt.Printf("Votre papa avait %d a votre naissance !", dadAgeOld)
}
