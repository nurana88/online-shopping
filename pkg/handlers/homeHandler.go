package handlers

import (
	"net/http"

	"github.com/nurana88/online-shopping/pkg/render"
)

func Home(res http.ResponseWriter, req *http.Request) {
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
