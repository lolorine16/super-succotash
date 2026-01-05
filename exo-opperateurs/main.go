package main

import "fmt"

func main() {
	var nb1 int
	var nb2 int

	fmt.Println("Entrer le premier nombre:")
	fmt.Scan(&nb1)
	fmt.Println("Entrer le deuxieme nombre:")
	fmt.Scan(&nb2)

	carre := nb1 * nb1
	fmt.Printf("Le carrer de %d est %d \n", nb1, carre)

	somme := nb1 + nb2
	mult := nb1 * nb2

	if nb1 > nb2 {
		fmt.Printf("%d est plus grand que %d leur somme est %d ; Leur multiplaction %d ; ", nb1, nb2, somme, mult)
		soustrac := nb1 - nb2
		div := nb1 / nb2
		fmt.Printf("Leur soustraction: %d ; Leur division %d", soustrac, div)
	} else {
		fmt.Printf("%d est plus petit que %d leur somme est %d ; Leur multiplaction %d ; ", nb1, nb2, somme, mult)
		soustrac := nb2 - nb1
		div := nb2 / nb1
		fmt.Printf("Leur soustraction: %d ; Leur division %d ", soustrac, div)
	}
}
