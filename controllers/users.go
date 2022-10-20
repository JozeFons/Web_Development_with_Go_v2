package controllers

import (
	"net/http"
	"github.com/JozeFons/Web_Development_with_Go_v2/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
