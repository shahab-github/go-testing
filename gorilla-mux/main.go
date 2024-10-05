package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Define a simple route
    r.HandleFunc("/", HomeHandler)

    // Start the server
    http.Handle("/", r)
    fmt.Println("Server is running on port 8000")
    http.ListenAndServe(":8000", nil)
}

// HomeHandler handles the home route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, World!"))
}

