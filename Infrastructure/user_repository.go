package Infrastructure

import (
	"database/sql"
	"errors"
	"fmt"

	config "github.com/nurana88/online-shopping/config"
)

const (
	queryInsertUser = "INSERT INTO users_shopping(name,lastname,email,password,date_created) VALUES(?,?,?,?,?);"
	queryFindEmail  = "SELECT id,name,lastname,email,date_created FROM users_shopping WHERE email=?;"

	queryFindUserByEmailAndPassword = "SELECT id,name,lastname,email,date_created FROM users_shopping WHERE email=? AND password=?;"

	queryUpdateUser = "UPDATE users_shopping SET name=?, lastname=?, email=? WHERE id=?;"
)

type DbActions struct {
	db *sql.DB
}

func NewDbActions(newDB *sql.DB) *DbActions {
	return &DbActions{db: newDB}
}

func (a *DbActions) InsertUser(user config.User) error {
	// Check if user already exists
	err := a.db.QueryRow(queryFindEmail, user.Email).Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.DateCreated)

	fmt.Println("Checking after queryRow...", err)
	fmt.Println("email...", user.Email)
	if err != sql.ErrNoRows {
		return errors.New("Email already exists...")
	}

	fmt.Println("Email checked", user)

	user.Password = config.GetMd5(user.Password)

	// Create new user
	stmt, err := a.db.Prepare(queryInsertUser)
	if err != nil {
		return errors.New("Cannot insert user in DB")
	}
	defer stmt.Close()

	insertUser, saveErr := stmt.Exec(user.Name, user.Lastname, user.Email, user.Password, user.DateCreated)
	if saveErr != nil {
		return errors.New("Couldn't execute user details")
	}

	userId, err := insertUser.LastInsertId()
	if err != nil {
		return errors.New("Couldn't insert userId")
	}
	user.Id = int(userId)
	fmt.Println("User was inserted", user)

	return nil
}

func (a *DbActions) FindUserInDB(user config.User) error {
	fmt.Println("Checking LoginUser...", user)

	err := a.db.QueryRow(queryFindUserByEmailAndPassword, user.Email, user.Password).Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.DateCreated)

	fmt.Println("Checking after prepare...", user)
	fmt.Println("Checking after prepare...", err)
	if err == sql.ErrNoRows {
		return errors.New("No info found in DB...")
	}
	fmt.Println("pwd is...", user.Password)

	if err != nil {
		return errors.New("err in queryrow...")
	}

	fmt.Println("Password is correct", err)

	return nil
}
