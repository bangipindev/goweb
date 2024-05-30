package route

import (
	"goweb/controllers/home"
	"goweb/controllers/user"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", home.Welcome)

	http.HandleFunc("/user", user.Index)
	http.HandleFunc("/user/create", user.Create)
	http.HandleFunc("/user/store", user.Store)
	http.HandleFunc("/user/edit", user.Edit)
	http.HandleFunc("/user/update", user.Update)
	http.HandleFunc("/user/delete", user.Destroy)
}
