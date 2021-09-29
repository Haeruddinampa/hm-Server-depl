package router

import (
	"homework-server-depl/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRoute() {
	var port string
	if port == "" {
		port = "8080"
	}
	log.Println("Starting development server at", port)
	log.Println("Quit the server with CTRL-C.")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/follower/{username}", controller.Follower).Methods("GET")
	myRouter.HandleFunc("/{userid}/detail", controller.DataUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}
