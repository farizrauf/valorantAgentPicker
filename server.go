package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting on http://localhost:8080...")
	http.Handle("/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
