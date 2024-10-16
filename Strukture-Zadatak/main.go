package main
import "fmt"


//ZADATAK
type Book struct{
	title string
	author string
	publisher Publisher
	price float64
}
type Publisher struct{
	name string
	location string
	booksPublished int
}

func (b Book) printDetails(){
	fmt.Printf("Book: %s, author: %s, publisher: %v, price: %.2f", b.title, b.author, b.publisher, b.price)
}

func (b1 *Book) applyDiscount(discountPercentage float64){
	discount := b1.price * discountPercentage/100
	b1.price = b1.price - discount
}

func (b2 *Book)updatePublisher(newPublisher Publisher){
	b2.publisher = newPublisher
}

func (p Publisher)printInfo(){
	fmt.Printf("Name: %s, Location: %s, Books published: %d", p.name, p.location, p.booksPublished)
}

func (p *Publisher) publishBook(){
	p.booksPublished++
}

func (p *Publisher)changeLocation(newLocation string){
	p.location = newLocation
}

func main(){
	publisher1 := Publisher{"Penguin", "London", 2000}
	newPublisher := Publisher{"New York Times", "New York", 2000}
	book1 := Book{"Harry Potter", "JK Rowling", publisher1, 20.00}

	newPublisher.changeLocation("Mostar")
	fmt.Println("")

	newPublisher.publishBook()
	fmt.Println("")
	
	book1.applyDiscount(10)
	fmt.Println("")

	book1.updatePublisher(newPublisher)
	fmt.Println("")

	newPublisher.printInfo()
	fmt.Println("")

	book1.printDetails()
}