package main

import (
	_ "basic-api/models"
	"basic-api/repository"
	_ "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, " Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/all", returnAllArticles)
	//log.Fatal(http.ListenAndServe(":8081", nil))
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", repository.ReturnAllArticles)
	myRouter.HandleFunc("/article/{id}", repository.ReturnArticleById)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
