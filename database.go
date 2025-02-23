package main

import (
	"database/sql"
	"fmt"
	"strings"
)

var db *sql.DB

func connectDB() {
	var err error
	db, err = sql.Open("sqlite3", "./recipes.db")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS recipes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		ingredients TEXT
	)`)
	if err != nil {
		fmt.Println("Error creating table:", err)
	}
}

func getRecipesFromDB() ([]Recipe, error) {
	rows, err := db.Query("SELECT id, name, ingredients FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []Recipe
	for rows.Next() {
		var r Recipe
		var ingredients string
		if err := rows.Scan(&r.ID, &r.Name, &ingredients); err != nil {
			return nil, err
		}
		r.Ingredients = strings.Split(ingredients, ",")
		recipes = append(recipes, r)
	}
	return recipes, nil
}
