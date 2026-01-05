package main

import "fmt"

func main() {
	var prenom string = "Laureen"
	nom := "EKON"

	fmt.Println("Entrer votre age :")
	var age int
	fmt.Scanln(&age)

	if age%2 == 0 {
		fmt.Println("Votre age est pair")
	} else {
		fmt.Println("Votre age est impaire")
	}

	fmt.Printf("Bonjour %s %s, tu as %d ans\n", prenom, nom, age)

	fmt.Println("Entrer votre fruit preferer :")
	var fruit string
	fmt.Scan(&fruit)
	fmt.Printf("Votre fruit preferer est %s \n", fruit)
}
