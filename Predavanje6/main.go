package main

import "fmt"

//---------------------------------------------------------------------
//FUNKCIJE
//---------------------------------------------------------------------
/*
func main() {
	fmt.Println(greet())
}

func greet() string {
	return "Greetings!"
} 
*/
//vraća string i moramo ili spremiti u varijablu ili preko println pozvati

/*
func main() {
	fmt.Println(greet("CodeCamper"))
}

func greet(name string) string {
	return fmt.Sprintf("Greetings %s!", name)
}
*/
//vraća ime

//--------------------------------------------------------------------
//VARIJABILNI PARAMETRI
//--------------------------------------------------------------------

/*
func main() {
	greet("CodeCamper", "CodeGuru", "CodeWizard")
}

func greet(names ...string) {
	for _, name := range names{
		fmt.Printf("Greetings %s!\n", name)
	}
}
*/

/*
func main() {
	names := []string{"Code", "Camper"}
	greet(names...)
}

func greet(names ...string) {
	for _, name := range names{
		fmt.Printf("Greetings %s!\n", name)
	}
}
*/

//--------------------------------------------------------------------
//RETURN VALUES
//--------------------------------------------------------------------	

/*
func main() {
	//names := []string{"Code", "Camper"}
	p := Person{"John", "Doe"}
	name, lastname := greet(p)

	fmt.Println(name)
	fmt.Println(lastname)
}

func greet(person Person) (string, string) {
	return person.Name, person.Lastname
}
*/


// _ označava skrivanje varijable ako nam u tom trenu nije potrebna

	// name, _/*lastname*/ := greet(p)

	// fmt.Println(name)
	// //fmt.Println(lastname)


//--------------------------------------------------------------------
//NAMED RETURN PARAMETERS
//--------------------------------------------------------------------

/*
type Person struct{
	Name string
	Lastname string
}

func main() {
	//names := []string{"Code", "Camper"}
	p := Person{"John", "Doe"}
	name, lastname := greet(p)

	fmt.Println(name)
	fmt.Println(lastname)
}

func greet(person Person) (name string, lastname string) {
	//return person.Name, person.Lastname
	name = person.Name
	lastname = person.Lastname

	return
}
*/

//--------------------------------------------------------------------
//ERROR HANDLING
//--------------------------------------------------------------------

/*
import "errors"
func main() {
	r, err := divide(5, 0)
	if err != nil{
		fmt.Printf("Error while dividing two numbers: %v", err.Error())
		return
	}
	fmt.Println(r)
}

func divide(a,b int) (int, error){
	if b==0{
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
*/

/*
type DivisionError struct {
	FirstNumber int
	SecondNumber int
	Message string

}

func (de *DivisionError) Error() string {
	return fmt.Sprintf("division error: %d / %d, err: %s", de.FirstNumber, de.SecondNumber, de.Message)
}
func main() {
	r, err := divide(5, 0)
	if err != nil{
		fmt.Printf("Error while dividing two numbers: %v", err.Error())
		return
	}
	fmt.Println(r)
}

func divide(a,b int) (int, error){
	if b==0{
		return 0, &DivisionError{a, b, "cannot divide by zero"}
	}
	return a / b, nil
}
*/


//--------------------------------------------------------------------
//ZADATAK KALKULATOR
//--------------------------------------------------------------------

func main(){
	result := calculator(10, 4, "/")
	fmt.Printf("Result: %.2f", result)
}

func calculator(firstNumber int, secondNumber int, operator string)(float64){
	switch operator{
		case "+":
			return float64(firstNumber) + float64(secondNumber)
		case "-":
			return float64(firstNumber) - float64(secondNumber)
		case "*":
			return float64(firstNumber) * float64(secondNumber)
		case "/":
			return float64(firstNumber) / float64(secondNumber)
		default:
			return 0
	}
}

