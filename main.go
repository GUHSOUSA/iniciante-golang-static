package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "Post request sucesso")
	name := r.FormValue("name")
	address := r.FormValue("addres")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprint(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Não encontrada", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Esse method não é suportado", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func home() {
	fileServer := http.FileServer(http.Dir("./site"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("servidor esta rodando na porta 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
