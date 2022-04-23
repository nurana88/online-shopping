package domain

type User struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
}
