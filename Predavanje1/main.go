package main

import "fmt"

func main() {
	//Prvi zadatak
	// var firstName string = "Lea"
	// var age int = 23
	// var height float64 = 1.75
	// var isStudent bool = true

	// fmt.Println(firstName, age, height, isStudent)

	//Drugi zadatak
	// var pi float64 = 3.14
	// var radius float64 = 5
	// var circle float64 = 2 * pi * radius

	// fmt.Println(circle)

	//Treci zadatak
	// var count int
	// var temperature float64
	// var isRaining bool
	// var message string

	// fmt.Println(count, temperature, isRaining, message)
	// fmt.Println("Count: ", count)
	// fmt.Println("Temperature: ", temperature)
	// fmt.Println("isRaining: ", isRaining)
	// fmt.Println("Message: ", message)

	//Cetvrti zadatak
	// var total int = 64
	// var discount float64 = 12.5
	// var finalPrice = float64(total) - discount

	// fmt.Println(finalPrice)

	//zad1
	var firstNumber int = 20
	var secondNumber int = 30
	firstNumber = firstNumber + secondNumber
	secondNumber = firstNumber - secondNumber
	firstNumber = firstNumber - secondNumber

	fmt.Println("First number is:", firstNumber)
	fmt.Println("Second number is:", secondNumber)

	//zad2
	var firstName string = "Lea"
	var lastName string = "Rajic"
	var fullName string = firstName + " " + lastName

	fmt.Println(fullName)
}
