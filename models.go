package main

type Recipe struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Ingredients []string `json:"ingredients"`
}

type ShoppingList struct {
	ID      int      `json:"id"`
	Items   []string `json:"items"`
}


var recipes = []Recipe{
	{ID: 1, Name: "Pasta", Ingredients: []string{"Pasta", "Tomato Sauce", "Garlic"}},
	{ID: 2, Name: "Omelette", Ingredients: []string{"Eggs", "Salt", "Butter"}},
}

var shoppingLists = []ShoppingList{
	{ID: 1, Items: []string{"Milk", "Eggs", "Bread"}},
}
