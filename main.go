package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2) // menetukan jumlah core yang diaktifkan untuk eksekusi program

	go print(5, "halo")
	go print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
	// Fungsi tersebut bisa menampung parameter bertipe interface{} berjumlah tak terbatas.
}
