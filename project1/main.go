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
	fmt.Printf("Hello %s, you are %d years old.", name, age)

	// Menue
	fmt.Println("\nChoose an option: ")
	fmt.Println("1- Check if you age is even or odd")
	fmt.Println("2- Enter tow number and see their sum")
	fmt.Println("3- Enter a number and see it square")
	fmt.Println("4- Enter a number and check if its positive, negative or zero")
	fmt.Println("5- Exit")
}
