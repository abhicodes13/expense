package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	expenses := []Expense{
		{
			ID:       1,
			Title:    "Burger",
			Amount:   18.99,
			Category: "Food",
		},
		{
			ID:       2,
			Title:    "Gas",
			Amount:   40.00,
			Category: "Transportation",
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Expense API is running!")
	})

	http.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var expense Expense

			// Decode JSON from the request body
			err := json.NewDecoder(r.Body).Decode(&expense)

			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			// Validate title
			if expense.Title == "" {
				http.Error(w, "Title is required", http.StatusBadRequest)
				return
			}

			// Validate amount
			if expense.Amount <= 0 {
				http.Error(w, "Amount must be greater than 0", http.StatusBadRequest)
				return
			}

			// Validate category
			if expense.Category == "" {
				http.Error(w, "Category is required", http.StatusBadRequest)
				return
			}

			// Generate ID
			expense.ID = len(expenses) + 1

			// Add expense to our slice
			expenses = append(expenses, expense)

			// Tell client the resource was created
			w.WriteHeader(http.StatusCreated)

			// Send the new expense back as JSON
			json.NewEncoder(w).Encode(expense)

			return
		}

		// GET /expenses
		json.NewEncoder(w).Encode(expenses)
	})

	fmt.Println("Server is running on 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server failed:", err)
	}
}