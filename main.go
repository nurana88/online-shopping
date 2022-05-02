package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nurana88/online-shopping/Presentation/HTTP/Handlers"
	"github.com/nurana88/online-shopping/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/registerauth", Handlers.RegisterUser)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/loginauth", Handlers.LoginUser)

	http.HandleFunc("/welcome", handlers.Welcome)

	// http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("assets/"))))

	fmt.Println("Starting session on :9000...")
	log.Fatal(http.ListenAndServe(":9000", nil))

	// log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("../templates"))))
}
