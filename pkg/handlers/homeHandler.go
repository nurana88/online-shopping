package handlers

import (
	"net/http"

	"github.com/nurana88/shop-template/pkg/render"
)

func Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "index.html")
}

func Shop(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "shop.html")
}

func renderFile(res http.ResponseWriter, html string) {

}
