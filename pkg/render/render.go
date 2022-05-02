package render

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func RenderTemplate(res http.ResponseWriter, html string) {
	tpl = template.Must(template.ParseFiles("templates/" + html))
	err := tpl.Execute(res, nil)
	if err != nil {
		fmt.Println("Error in render template", err)
	}
}
