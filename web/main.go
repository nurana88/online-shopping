package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// http.HandleFunc("/", handlers.Home)
	// http.HandleFunc("/shop", handlers.Shop)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../templates"))))

	fmt.Println("Starting session on :9000...")
	log.Fatal(http.ListenAndServe(":9000", nil))

	// log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("../templates"))))
}
