package user

import (
	"goweb/entities"
	"goweb/models/usermodel"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	usermodel := usermodel.GetAll()
	data := map[string]interface{}{
		"users": usermodel,
	}
	temp, err := template.ParseFiles("views/user/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/user/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		user := entities.User{
			Name:      r.Form.Get("name"),
			Email:     r.Form.Get("email"),
			Password:  r.Form.Get("password"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		log.Println(user)
		if true := usermodel.Store(user); !true {
			temp, _ := template.ParseFiles("views/user/create.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/user/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		user := usermodel.GetById(id)
		data := map[string]interface{}{
			"user": user,
		}
		log.Println(data)
		temp.Execute(w, data)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		idString := r.Form.Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		user := entities.User{
			Id:        id,
			Name:      r.Form.Get("name"),
			Email:     r.Form.Get("email"),
			UpdatedAt: time.Now(),
		}
		log.Println(user)
		if true := usermodel.Update(user); !true {
			temp, _ := template.ParseFiles("views/user/edit.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	}
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}
	log.Println(id)
	usermodel.Delete(id)
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}
