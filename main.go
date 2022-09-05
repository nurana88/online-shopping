package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/nurana88/online-shopping/Infrastructure"
	"github.com/nurana88/online-shopping/Presentation/HTTP/Handlers"
	domain "github.com/nurana88/online-shopping/domain"
	"github.com/nurana88/online-shopping/pkg/handlers"
)

func main() {

	dataSourceName := "root:pwd@tcp(127.0.0.1:3306)/testdb"

	var err error
	DB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Couldn't connect to DB", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("Can not be connected in Ping", err)
		return
	}
	fmt.Println("Successfully connected to DB...")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/registerauth", Handlers.NewRegisterUserHandler(domain.NewDbUserUsercase(Infrastructure.NewDbActions(DB))))
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/loginauth", Handlers.LoginUser(domain.NewDbUserUsercase(Infrastructure.NewDbActions(DB))))

	http.HandleFunc("/welcome", handlers.Welcome)

	// http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("assets/"))))

	fmt.Println("Starting session on :9000...")
	log.Fatal(http.ListenAndServe(":9000", nil))

	// log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("../templates"))))
}
