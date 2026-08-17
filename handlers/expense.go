package handlers

import (
	"fmt"
	"net/http"
)

func Expense(w http.ResponseWriter, r *http.Request){
	fmt.Println(w)

}