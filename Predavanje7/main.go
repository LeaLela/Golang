package main

import "fmt"
import "errors"
/*
type BankAccount struct {
	owner string
	balance float64
}

func (a *BankAccount) Deposit(amount float64) {
    a.balance += amount
    fmt.Printf("Deposited %.2f to %s's account. New balance: %.2f\n", amount, a.owner, a.balance)
}

func (a *BankAccount) Withdraw(amount float64) error{
	if a.balance < amount {
		return errors.New("insufficient funds")
	}

	a.balance -= amount
    fmt.Printf("Withdrew %.2f from %s's account. New balance: %.2f\n", amount, a.owner, a.balance)
	return nil
}
func main() {
	account := BankAccount{
		owner: "John",
		balance: 100.0,
	}
	account.Deposit(200)

	err := account.Withdraw(160.0)
    if err != nil {
        fmt.Println("Error:", err)
    }

	
}
*/

type Book struct {
	Title string
	Author string
	Quantity int
}

func (b *Book) AddBooks(amount int){
	b.Quantity += amount
	fmt.Printf("Added %d books to %s\n", amount, b.Title)
}

func (b *Book) RemoveBooks(amount int) error{
	if b.Quantity < amount {
		return errors.New("not enough books in stock")
	}

	b.Quantity -= amount
	fmt.Printf("Removed %d books from %s\n", amount, b.Title)
	return nil
}

func (b *Book) printBookInfo(){
	fmt.Printf("Book Title: %s\nAuthor: %s\nAvailable Stock: %d\n", b.Title, b.Author, b.Quantity)
}

func main() {
	book := Book{
		Title: "Harry Potter",
		Author: "JK Rowling",
		Quantity: 10,
	}

	book.AddBooks(25)

	err := book.RemoveBooks(20)
	if err != nil{
		fmt.Println("Error:", err)
	}

	book.printBookInfo()

}