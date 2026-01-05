package main

import "fmt"

func main() {
	var nb1 int
	var nb2 int

	fmt.Println("Entrer le premier nombre:")
	fmt.Scan(&nb1)
	fmt.Println("Entrer le deuxieme nombre:")
	fmt.Scan(&nb2)

	somme := nb1 + nb2

	if nb1 > nb2 {
		fmt.Printf("%d est plus grand que %d leur somme est %d ", nb1, nb2, somme)
	} else {
		fmt.Printf("%d est plus petit que %d leur somme est %d ", nb1, nb2, somme)
	}
}
