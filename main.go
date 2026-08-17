package main

import (
	"fmt"
	"net/http"
)

func main (){
	 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Expense API is running!")
    })
	
	fmt.Println("Server is running on 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil{
		fmt.Println("SErver failed", err)
	}



}

