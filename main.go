package main // nom du packet auquel le fichier appartient

import "fmt" // lib fmt

func main() {
	fmt.Println("Hello, World!")

	/*
	 * bool : stocke un bouleen
	 * string : stocke une chaine de caracteres
	 * int, int8, int16, int32, int64 : stockent un entier signe code sur n bits
	 * uint, ... , uint64 : stockent un entier non-signe sur n bits
	 * byte : alias pour uint8
	 * rune : alias pour int32, represente un caracteres Unicode
	 * float32, float64 : represente des nombres flottants codés sur 32 ou 64 bits
	 * TYPES SPECIAUX
	 * error : Type pour les messages d'erreurs
	 * interface : definit et decrit les methodes exactes qu'un autre type doit avoir
	 */

	// VARIABLES
	/*
	 * var name type = value  --declaration en precisant son type
	 * name := value --declaration en mode “devine”
	 */
	var firstname string = "Laureen"
	lastname := "EKON"
	var age int
	age = 18
	fmt.Printf("Bonjour %s %s, tu as %d ans\n", firstname, lastname, age)

	// CONSTANTES
	const pi = 3.14
	const g = 9.81

	// declaration groupées
	const (
		height = 180
		grade  = 2
	)
}
