package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	var tmp, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/index", func(rw http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		err = tmp.ExecuteTemplate(rw, "index", data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(rw http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		err = tmp.ExecuteTemplate(rw, "about", data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
