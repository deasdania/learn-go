package main

import (
	"fmt"
	"net/http"
	"time"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "hello world"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("server starter at %s\n", address)

	// menggunakan listenAndServe
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	// menggunakan Server Kelebihan menggunakan http.Server salah satunya adalah kemampuan untuk mengubah beberapa konfigurasi default web server Go.
	server := new(http.Server)
	server.Addr = address
	errS := server.ListenAndServe()
	if errS != nil {
		fmt.Println(errS.Error())
	}

	// seperti
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10
}
