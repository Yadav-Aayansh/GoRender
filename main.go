package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default if PORT not set
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Render!")
	})

	fmt.Println("Server running on port", port)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
