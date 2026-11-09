package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error  Starting Server", err)
	}

}
