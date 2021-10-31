package main

import (
	"log"
	"net/http"
	"os"

	"GoRestApi/routes"
)

func main() {

	router := routes.MovieRoutes()

	http.Handle("/api", router)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
