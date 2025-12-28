package main // nom du packet auquel le fichier appartient

import "fmt" // lib fmt

func main() {
	/*
			* addition -> +
			* soustraction -> -
			* multiplication -> *
			* division -> /
			* modulo -> %
		  * incrémentation -> ++
		  * decrémentation -> --
	*/

	// OPERATEURS D'ASSIGNEMENT
	/*
	* a = b -> a := b
	* a := b -> bar a = b
	* a += b -> a = a + b
	* a -= b -> a = a - b
	* a *= b -> a = a * b
	* a /= b -> a = a / b
	* a %= b -> a = a % b
	 */

	// OPERATEURS RELATIONNELS
	/*
	* a == b -> égal à
	* a != b -> différent de
	* a > b  -> supérieur à
	* a < b  -> inférieur à
	* a >= b -> supérieur ou égal à
	* a <= b -> inférieur ou égal à
	 */

	// OPERATEURS LOGIQUES
	/*
	* && et
	* || ou
	* ! diff
	 */

	a := 10
	b := 3
	fmt.Println("Addition :", a+b)
	fmt.Println("Soustraction :", a-b)
	fmt.Println("Multiplication :", a*b)
	fmt.Println("Division :", a/b)
	fmt.Println("Modulo :", a%b)
}
