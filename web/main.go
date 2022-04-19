package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nurana88/shop-template/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/shop", handlers.Shop)

	fmt.Println("Starting session on :9000...")
	log.Fatal(http.ListenAndServe(":9000", nil))

	// http.Handle("/temp/", http.StripPrefix("/temp", http.FileServer(http.Dir("../templates"))))

	// log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("../templates"))))
}
