package main

import (
	f "fmt"
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Path
	f.Println(link)
}

func main() {
	http.HandleFunc("/", redirect)
	f.Println("Server running on port 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
