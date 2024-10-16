package main
import "fmt"

//zadatak1

// type Pravokutnik struct {
// 	sirina float64
// 	visina float64
// }

// func (p Pravokutnik) povrsina(){
// 	fmt.Println(p.sirina*p.visina)
// }

// func (p Pravokutnik) opseg(){
// 	fmt.Println(2*(p.sirina+p.visina))
// }

// func main(){
// 	pravokutnik1 := Pravokutnik{10,20}
// 	pravokutnik1.povrsina()
// 	pravokutnik1.opseg()
// }

//zadatak2
type Adresa struct {
	grad string
	ulica string
}

type Osoba struct{
	ime string
	godine int
	adresa Adresa
}

func (o Osoba) ispis(){
	fmt.Printf("%s, %d godina, zivi u %s, %s", o.ime, o.godine, o.adresa.grad, o.adresa.ulica)
}

func main(){
	adresa1 := Adresa{"Mostar", "Ante Starcevica"}
	osoba1 := Osoba{"Iva Ivić", 20, adresa1} 
	osoba1.ispis()
}