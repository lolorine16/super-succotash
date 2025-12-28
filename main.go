package main // nom du packet auquel le fichier appartient

import "fmt" // lib fmt

func main() {
	/*
		*Exercice
		• Créer un programme qui déclare une variable de manière explicite myNum1 , lui
		attribue la valeur 50 et l’affiche.
		• Faire la même chose avec myNum2 , en la déclarant de manière implicite.
	*/
	var myNum1 int = 50
	fmt.Println(myNum1)

	myNum2 := 50
	fmt.Println(myNum2)
}
