package domain

import "net/http"

const (
	queryFindUser   = "SELECT id,name,last_name,email,date_created FROM users WHERE id=?;"
	queryInsertUser = "INSERT INTO users(name,last_name,email,date_created,password) VALUES(?,?,?,?,?);"
	queryUpdateUser = "UPDATE users SET name=?, last_name=?, email=? WHERE id=?;"
)

// Register user
func SaveUser(res http.ResponseWriter, req *http.Request) {
}

// Login user
func Login(res http.ResponseWriter, req *http.Request) {

}

// Check validity of user
func IsUserValid() {

}
