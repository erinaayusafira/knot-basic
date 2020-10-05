package main

import "fmt"
import "net/http"
import "MVCKnotbasic/app/controllers"
import "MVCKnotbasic/app/assets/github.com/gorilla/mux"

func main(){

	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("").Subrouter()
	sub.Methods("GET").Path("/").HandlerFunc(controllers.GetData)

	http.Handle("/static/",
	http.StripPrefix("/static/",
	http.FileServer(http.Dir("assets"))))

fmt.Println("server started at localhost : 8000")
http.ListenAndServe(":8000", router)
} 