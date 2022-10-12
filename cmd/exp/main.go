package main

import (
	"os"
	"html/template"
)

type User struct {
	Name string
	Bio string
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User {
		Name: "Joe",
		Bio: `<script>alert("Haha, you have been hacked!");</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}