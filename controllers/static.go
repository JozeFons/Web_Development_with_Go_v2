package controllers

import (
	"html/template"
	"net/http"
	"github.com/JozeFons/Web_Development_with_Go_v2/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer: "Yes, We offer a free trial for 30 days on paid plans.",
		},
		{
			Question: "What are your support hours?",
			Answer: "We have support staff answering emails 24/7, though response times may be a bit slower on weekends.",
		},
		{
			Question: "How do I contact support?",
			Answer: `"Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com"`,
		},
		{
			Question: "Where is your office located?",
			Answer: `"Our entire team is remote!"`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
