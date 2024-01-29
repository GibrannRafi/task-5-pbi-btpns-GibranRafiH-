package main

import (
	"log"
	"net/http"

	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/controllers/authcontroller"
	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/controllers/photocontroller"
	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/middlewares"
	"github.com/GibrannRafi/task-5-pbi-btpns-GibranRafiH-/app/models"
	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/user/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/user/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/user/logout", authcontroller.Logout).Methods("GET")
	r.HandleFunc("/user", authcontroller.Index).Methods("GET")
	r.HandleFunc("/user/update/{id}", authcontroller.UpdateUser).Methods("PUT")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/photo", photocontroller.Create).Methods("POST")
	api.HandleFunc("/photo", photocontroller.Index).Methods("GET")
	api.HandleFunc("/photo/show/{id}", photocontroller.Show).Methods("GET")
	api.HandleFunc("/update/{id}", photocontroller.Update).Methods("PUT")
	api.HandleFunc("/delete/{id}", photocontroller.Delete).Methods("DELETE")
	api.Use(middleware.JWTMiddleware)
	log.Fatal(http.ListenAndServe(":8080", r))
}
