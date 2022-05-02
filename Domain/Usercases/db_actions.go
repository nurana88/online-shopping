package Usercases

import (
	"database/sql"
	"fmt"

	"github.com/nurana88/online-shopping/Infrastructure"
	"golang.org/x/crypto/bcrypt"
)

const (
	queryInsertUser = "INSERT INTO users_shopping(name,lastname,email,password,date_created) VALUES(?,?,?,?,?);"
	queryFindEmail  = "SELECT id,name,lastname,email,date_created FROM users_shopping WHERE email=?;"

	queryFindUserByEmailAndPassword = "SELECT id,name,lastname,email,date_created FROM users_shopping WHERE email=? AND password=?;"

	queryUpdateUser = "UPDATE users_shopping SET name=?, lastname=?, email=? WHERE id=?;"
)

func (user User) InsertUser() *ApiErr {
	// Check if user already exists
	err := Infrastructure.DB.QueryRow(queryFindEmail, user.Email).Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.DateCreated)

	fmt.Println("Checking after queryRow...", err)
	fmt.Println("email...", user.Email)
	if err != sql.ErrNoRows {
		return NewInternalError("Email already exists")
	}

	fmt.Println("Email checked", user)

	// Make hashed password
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return NewInternalError("Password couldn't be hashed")
	}
	user.Password = string(hash)

	// Create new user
	stmt, err := Infrastructure.DB.Prepare(queryInsertUser)
	if err != nil {
		return NewInternalError("Cannot insert user in DB")
	}
	defer stmt.Close()

	insertUser, saveErr := stmt.Exec(user.Name, user.Lastname, user.Email, user.Password, user.DateCreated)
	if saveErr != nil {
		return NewInternalError("Couldn't execute user details")
	}

	userId, err := insertUser.LastInsertId()
	if err != nil {
		return NewInternalError("Couldn't insert userId")
	}
	user.Id = int(userId)
	fmt.Println("User was inserted", user)

	return nil
}

func (user *User) FindUserInDB() *ApiErr {
	fmt.Println("Checking LoginUser...", user)

	err := Infrastructure.DB.QueryRow(queryFindUserByEmailAndPassword, user.Email, user.Password).Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.DateCreated)

	fmt.Println("Checking after prepare...", user)
	fmt.Println("Checking after prepare...", err)
	if err == sql.ErrNoRows {
		return NewInternalError("No info found in DB...")
	}
	fmt.Println("pwd is...", user.Password)

	if err != nil {
		return NewInternalError("err in queryrow...")
	}

	fmt.Println("Password is correct", err)

	return nil
}
