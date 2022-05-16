package handlers

import (
	"fmt"
	"net/http"

	"github.com/nurana88/online-shopping/pkg/render"
)

func Home(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my2-cookie")
	if err != nil {
		fmt.Println("Cookie wasn't found...")
		cookie = &http.Cookie{
			Name:     "my2-cookie",
			Value:    "Index2 file",
			HttpOnly: true,
			//MaxAge:   60,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println("Cookie is", cookie)
	render.RenderTemplate(res, "index.html")
}

func Register(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "register.html")
}

func Login(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "login.html")
}
func Shop(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "shop.html")
}
func Welcome(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "welcome.html")
}
