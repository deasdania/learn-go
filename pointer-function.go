package main

import "fmt"

type Alamat struct {
	City     string
	Province string
	Country  string
}

func ChangeCountryToIndonesia(alamat *Alamat) {
	alamat.Country = "Indonesia"
}

func main() {
	alamat1 := Alamat{"Malang", "Jawa Timur", ""}

	ChangeCountryToIndonesia(&alamat1)
	fmt.Println(alamat1)
}
