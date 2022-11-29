package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/JozeFons/Web_Development_with_Go_v2/controllers"
)

func Must(t Template, err error) Template {
	controllers.CheckError(err)
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {

	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error on executing the template!", http.StatusInternalServerError)
		return
	}
}
