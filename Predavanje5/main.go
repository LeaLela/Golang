package main

import "fmt"


func main() {

	// for i := 0; i < 10; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	if i == 7 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }


	//zadatak

	// broj := 5
    // faktorijel := 1

    // for i := 1; i <= broj; i++ {
    //     faktorijel *= i
    // }

    // fmt.Printf("Faktorijel od %d je %d\n", broj, faktorijel)

	// var number, iterator int

	// number =5
	// var result int = 1

	// for iterator = 1; iterator <= number; iterator++ {
	// 	result *= iterator
	// }
	// fmt.Println("The factorial of", number, " is",result)	
	// fmt.Println("The factorial of", number, " is", factorial(5))


	//array

	// myArray := [5]int{10, 20, 30, 40, 50}
	// sum := 0
	
	// for i:=0; i<len(myArray); i++ {
	// 	sum += myArray[i]
	// }
	// average := float64(sum)/float64(len(myArray))
	// fmt.Printf("The average of the array elements is %.2f\n", average)


	//slice

	//mySlice := []int{22,1133,818,181,718}
	//sum := 0
	// fmt.Println("Length of myArray after appending 120: ",len(mySlice))
	// fmt.Println("Capacity of myArray after appending 120: ",cap(mySlice))
	// mySlice = append(mySlice,120)
	// fmt.Println("Length of myArray after appending 120: ",len(mySlice))
	// fmt.Println("Capacity of myArray after appending 120: ",cap(mySlice))
	// mySlice = append(mySlice,12)
	// fmt.Println("Length of myArray after appending 120: ",len(mySlice))
	// fmt.Println("Capacity of myArray after appending 120: ",cap(mySlice))

	// for i:=0; i<len(mySlice); i++ {
	// 	sum += mySlice[i]
	// }
	// average := float64(sum)/float64(len(mySlice))
	// fmt.Printf("The average of the array elements is %.2f\n", average)


	//map

	// phones := map[string]int{
	// 	"phone1": 300,
	// 	"phone2": 500,
	// 	"phone3": 120,
	// }
	// fmt.Println(phones)

	// phone2 := phones["phone2"]
	// fmt.Println(phone2)

	// fmt.Print(phones["nema"])

	// if price, exists := phones["phone3"]; exists {
	// 	fmt.Printf("price is %v",price)
	// }else {
	// 	fmt.Println("Phone not found")
	// }

	names := map[string]int{
		"Ivan": 10,
		"Marko": 20,
		"Ana": 30,
	}
	names["Jure"] = 40
	names["Ivan"] = 45

	fmt.Println(names)
}
