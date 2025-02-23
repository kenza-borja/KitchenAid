package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	log.Println("Starting server...")
	mux := http.NewServeMux()

	mux.HandleFunc("/recipes", GetRecipes)
	mux.HandleFunc("/shopping-lists", GetShoppingLists)

	fmt.Println("Server running on http://localhost:8080")
	err := 	http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
