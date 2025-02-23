package main

import (
	"encoding/json"
	"net/http"
)

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func GetShoppingLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shoppingLists)
}
