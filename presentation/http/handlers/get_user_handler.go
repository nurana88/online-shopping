package handlers

import (
	"fmt"
	"github.com/nurana88/online-shopping/config"
	"github.com/nurana88/online-shopping/domain"
	"log"
	"net/http"
)

type getUserHandler struct {
	userCreator domain.UserGetAction
}

func NewGetUserHandler(userCreator domain.UserGetAction) http.Handler {
	return &getUserHandler{userCreator: userCreator}
}

func (h *getUserHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var request config.LoginDetails
	//res.Header().Set("Content-Type", "json/application")
	err := req.ParseForm()
	if err != nil {
		log.Fatal("Error in parsing request:", err)
	}

	request.Email = req.FormValue("email")
	request.Password = req.FormValue("pwd")

	result, saveErr := h.userCreator.GetUser(request)
	if saveErr != nil {
		log.Println(saveErr)
		w.Write([]byte("Bad request 400"))
		//res.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println(result)

	http.Redirect(w, req, "/welcome", 302)
}
