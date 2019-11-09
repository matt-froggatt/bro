package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	if _, err := fmt.Fprintf(w, "Hello World") ; err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := mux.NewRouter()
	buildHandler := http.FileServer(http.Dir("frontend/build"))
	router.PathPrefix("/").Handler(buildHandler)
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/build/static")))
	router.PathPrefix("/static/").Handler(staticHandler)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalln(err)
	}
}
