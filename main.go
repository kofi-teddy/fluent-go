package main

import (
	"fluent-go/routes"
	"fmt"
	"net/http"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, world!")
// }

// func main() {
// 	// fmt.Println("Hello, world!")
// 	http.HandleFunc("/", handler)
// 	fmt.Println("Server is running on http://localhost:8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	// start the server
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
