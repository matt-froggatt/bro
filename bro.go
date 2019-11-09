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
	buildHandler := http.FileServer(http.Dir("frontend/"))
	router.PathPrefix("/").Handler(buildHandler)

	http.HandleFunc("/", helloWorld)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalln(err)
	}
}
