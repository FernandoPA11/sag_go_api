package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":3000", router)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
