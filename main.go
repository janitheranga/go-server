package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.\n", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed.\n", http.StatusMethodNotAllowed)
		return
	}

	_, _ = fmt.Fprint(w, "Hello World!\n")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed.\n", http.StatusMethodNotAllowed)
		return
	}

	_, _ = fmt.Fprintf(w, "POST request is successfull\n")
	name := r.FormValue("name")
	//if name == "" {
	//	http.Error(w, "Name is empty.\n", http.StatusBadRequest)
	//}
	address := r.FormValue("address")
	//if address == "" {
	//	http.Error(w, "Address is empty.\n", http.StatusBadRequest)
	//}

	_, _ = fmt.Fprintf(w, "Name is %s", name)
	_, _ = fmt.Fprintf(w, "Address is %s", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
