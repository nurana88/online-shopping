package Handlers

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/nurana88/online-shopping/config"
	usercases "github.com/nurana88/online-shopping/domain"
)

// Register user

func NewRegisterUserHandler(userCreator *usercases.DBUserUsercase) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var user config.User
		//res.Header().Set("Content-Type", "json/application")
		err := req.ParseForm()
		if err != nil {
			log.Fatal("Error in parsing request:", err)
		}

		user.Name = req.FormValue("name")
		user.Lastname = req.FormValue("lastname")
		user.Email = req.FormValue("email")
		user.Password = req.FormValue("pwd")
		user.PasswordRepeat = req.FormValue("pwd-repeat")

		result, saveErr := userCreator.CreateUser(user)
		if saveErr != nil {
			log.Println(saveErr)
			res.Write([]byte("Bad request 400"))
			//res.WriteHeader(http.StatusBadRequest)
		}

		fmt.Println(result)

		http.Redirect(res, req, "/welcome", 302)

		// bytes, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	log.Println("Bytes error: ", err)
		// }

		// err = json.Unmarshal(bytes, &user)
		// if err != nil {
		// 	log.Println("Json error: ", err)
		// }

		// byte, newErr := json.Marshal(result)
		// if newErr != nil {
		// 	log.Fatal("Error in marshaling user", newErr)
		// }

	}

}

func LoginUser(userCreator *usercases.DBUserUsercase) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var request config.LoginDetails
		//res.Header().Set("Content-Type", "json/application")
		err := req.ParseForm()
		if err != nil {
			log.Fatal("Error in parsing request:", err)
		}

		request.Email = req.FormValue("email")
		request.Password = req.FormValue("pwd")

		result, saveErr := userCreator.GetUser(request)
		if saveErr != nil {
			log.Println(saveErr)
			res.Write([]byte("Bad request 400"))
			//res.WriteHeader(http.StatusBadRequest)
		}

		fmt.Println(result)

		http.Redirect(res, req, "/welcome", 302)
	}
}

// // Update user
// func UpdateUser(res http.ResponseWriter, req *http.Request) {

// }

// // Check validity of user
// func IsUserValid() {

// }
