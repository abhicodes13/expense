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
        json.NewEncoder(w).Encode(expenses)
    })

    fmt.Println("Server is running on 8080")

    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        fmt.Println("Server failed:", err)
    }
}