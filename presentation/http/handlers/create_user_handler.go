package handlers

import (
	"fmt"
	"github.com/nurana88/online-shopping/logger"
	"github.com/sirupsen/logrus"
	"net/http"

	config "github.com/nurana88/online-shopping/config"
	"github.com/nurana88/online-shopping/domain"
)

type registerUserHandler struct {
	userCreator domain.UserCreateAction
}

func NewRegisterUserHandler(userCreator domain.UserCreateAction) http.Handler {
	return &registerUserHandler{userCreator: userCreator}
}

func (h *registerUserHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logrus.Info(req.Context(), "Serving Register handler")

	var user config.User
	//res.Header().Set("Content-Type", "json/application")
	err := req.ParseForm()
	if err != nil {
		fmt.Println("Error in registration handler", err)
		//logger.FromContext(req.Context()).WithError(err).Error("Error in parsing request")
	}

	user.Name = req.FormValue("name")
	user.Lastname = req.FormValue("lastname")
	user.Email = req.FormValue("email")
	user.Password = req.FormValue("pwd")
	user.PasswordRepeat = req.FormValue("pwd-repeat")

	fmt.Println("User in servehttp handler is", user)

	saveErr := h.userCreator.CreateUser(user)

	if saveErr != nil {
		fmt.Println("Error in registration handler", saveErr)
		logger.FromContext(req.Context()).WithError(saveErr).Error("Error in creating user handler")
		w.Write([]byte("Bad request 400"))
		//res.WriteHeader(http.StatusBadRequest)
	}
	//uuid := req.Context().Value("email")
	//log.Printf("[%v] Returning email - Healthy", uuid)
	//fmt.Println(result)
	//
	//bytes, err := ioutil.ReadAll(req.Body)
	//if err != nil {
	//	log.Println("Bytes error: ", err)
	//}
	//
	//err = json.Unmarshal(bytes, &user)
	//if err != nil {
	//	log.Println("Json error: ", err)
	//}

	//byte, newErr := json.Marshal(result)
	//if newErr != nil {
	//	log.Fatal("Error in marshaling user", newErr)
	//}

	http.Redirect(w, req, "/welcome", 302)
}
