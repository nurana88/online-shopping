package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {

	http.HandleFunc("/", Home)

	fmt.Println("Starting session...")
	log.Fatal(http.ListenAndServe(":9000", nil))

}

func Home(res http.ResponseWriter, req *http.Request) {
	//Return home page index.html
}
