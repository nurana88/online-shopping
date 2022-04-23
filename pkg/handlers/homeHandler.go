package handlers

import (
	"net/http"

	"github.com/nurana88/shop-template/pkg/render"
)

func Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "index.html")
	res.WriteHeader(http.StatusOK)
}

func Shop(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "shop.html")
	res.WriteHeader(http.StatusOK)
}

func renderFile(res http.ResponseWriter, html string) {

}
