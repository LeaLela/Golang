package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int

}

// type Book struct{
// 	Title string
// 	Author string
// 	Year int
// 	Genre string
// 	Publisher Publisher
// }

// type Publisher struct{
// 	Name string
// 	Address string
// }

func main() {
	person1 := Person{"Ivo", "Ivic", 30}
	// person2 := Person{
	// 	FirstName: "Marko",
	// 	LastName:  "Markovic",
	// 	Age:       40,
	// }
	person1.greet()

	// fmt.Println(person1.FirstName)
	// fmt.Println(person1.LastName)
	// fmt.Println(person1.Age)


	// fmt.Println(person2.FirstName)
	// fmt.Println(person2.LastName)
	// fmt.Println(person2.Age)

	//task1/task2
	// publisher1 := Publisher{"Penguin", "London"}

	// book1 := Book{"Harry Potter", "JK Rowling", 1997, "fiction", publisher1}
	// book2 := Book{"Lord of the Rings", "JRR Tolkien", 1954, "fiction", publisher1}
	// book3 := Book{"The Hobbit", "JRR Tolkien", 1937, "fiction", publisher1}

	// fmt.Println("Book info: ", book1.Author, book1.Title, book1.Year, book1.Year, book1.Publisher.Name, book1.Publisher.Address)
	// fmt.Println("Book info: ", book2.Author, book2.Title, book2.Year, book2.Year, book2.Publisher.Name, book2.Publisher.Address)
	// fmt.Println("Book info: ", book3.Author, book3.Title, book3.Year, book3.Year, book3.Publisher.Name, book3.Publisher.Address)

	
}

func (p Person) greet() {
	fmt.Printf("Hello, my name is: %s %s and I am %d years old", p.FirstName, p.LastName, p.Age)
}

