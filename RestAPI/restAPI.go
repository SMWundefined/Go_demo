package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Article JSON imports for the API
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

//Articles being defined in the func below when the /article is hit
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "demoTitle", Desc: "This is the description of the Title", Content: "Hello World from Go"},
	}
	fmt.Println("Endpoint reached for All Articles")
	json.NewEncoder(w).Encode(articles)
}

//main page of the http endpoint for the API

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpont Reached")
}

//handling the API
func handlingRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	handlingRequest()
}
