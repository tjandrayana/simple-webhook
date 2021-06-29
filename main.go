package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Serving static files
	router.HandleFunc("/me-here", MeHere).Methods(http.MethodPost)
	router.HandleFunc("/ping", Ping).Methods(http.MethodGet)

	fmt.Println("Serving requests on port 1234")
	err := http.ListenAndServe(":1234", router)
	log.Fatal(err)
}

func MeHere(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[MeHere] unable to read body, err: ", err.Error())
		return
	}
	log.Printf("[MeHere] body %s\n", string(body))
	return

}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"pong"}`))
}
