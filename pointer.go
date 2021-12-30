package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	//address2 := &address1
	// or
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Depok"

	//fmt.Println(address1)
	//fmt.Println(address2)
	////{Depok Jawa Barat Indonesia}
	////&{Depok Jawa Barat Indonesia}
	//
	//address2 = &Address{"Malang", "Jawa Timur", "Indonesia"} //creating new pointer
	//
	//fmt.Println(address1)
	//fmt.Println(address2)
	//{Depok Jawa Barat Indonesia}
	//&{Malang Jawa Timur Indonesia}

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} //dengan menambahkan * merefer ke address1
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	//{Malang Jawa Timur Indonesia}
	//&{Malang Jawa Timur Indonesia}
	//&{Malang Jawa Timur Indonesia}

	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
