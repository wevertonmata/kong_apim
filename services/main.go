package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        message := os.Getenv("MESSAGE")
        fmt.Fprintf(w, message)
    })

    port := ":8083"
    if p := os.Getenv("PORT"); p != "" {
        port = ":" + p
    }

    fmt.Printf("Server listening on port %s\n", port)
    http.ListenAndServe(port, nil)
}
