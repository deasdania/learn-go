package main

import "fmt"

type Man struct {
	Name string
}

// cukum menggunakan * sebelum Man
func (man *Man) Married() {
	man.Name = "Mr." + man.Name
	fmt.Println("di Method", man.Name)
}

func main() {
	dani := Man{"Dani"}
	dani.Married()

	fmt.Println(dani.Name)
}
