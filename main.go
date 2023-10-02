package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/JuanPO17/go-gorm-restapi/routes"
)



func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	
	http.ListenAndServe(":3000", r)
}

