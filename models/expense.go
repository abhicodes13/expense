package models

type Expense struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
}
