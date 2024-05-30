package main

import (
	"goweb/config"
	"goweb/route"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	route.HandleRequest()
	log.Println("Server running at localhost:4000")
	http.ListenAndServe("localhost:4000", nil)
}
