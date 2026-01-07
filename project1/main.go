package main

import "fmt"

func main() {
	// Message de bienvenue
	fmt.Println("Welcome to Go Basics Validator")

	// Variables
	var name string
	var age int

	// Informations
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)

	// Message
	fmt.Printf("Hello %s, you are %d years old.\n", name, age)

	// boucle infinie
	for i := 0; i > 0; i++ {
		// Menue
		fmt.Println("\nChoose an option: ")
		fmt.Println("1- Check if you age is even or odd")
		fmt.Println("2- Enter tow number and see their sum")
		fmt.Println("3- Enter a number and see it square")
		fmt.Println("4- Enter a number and check if its positive, negative or zero")
		fmt.Printf("5- Exit\n")

		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Printf("GoodBye %s !", name)
			break
		} else if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice, please try again.")
		} else {
			// Options
			switch choice {
			case 1:
				if age%2 == 0 {
					fmt.Println("Your age is even !")
				} else {
					fmt.Println("Your age is odd !")
				}
			case 2:
				var nb1, nb2 int
				fmt.Println("Enter number1: ")
				fmt.Scanln(&nb1)
				fmt.Println("Enter number2: ")
				fmt.Scanln(&nb2)
				somme := nb1 + nb2
				fmt.Printf("The sum of %d and %d is %d\n", nb1, nb2, somme)
			case 3:
				var nb0 int
				fmt.Println("Enter a number: ")
				fmt.Scanln(&nb0)
				carrer := nb0 * nb0
				fmt.Printf("The square of %d is %d\n", nb0, carrer)
			case 4:
				var nb int
				fmt.Println("Enter a number: ")
				fmt.Scanln(&nb)
				if nb > 0 {
					fmt.Printf("%d is a positive number !\n", nb)
				} else if nb < 0 {
					fmt.Printf("%d is a negative number !\n", nb)
				} else {
					fmt.Printf("Is null !\n")
				}
			case 5:
				fmt.Printf("GoodBye %s !", name)
				break
			}
		}
	}
}
