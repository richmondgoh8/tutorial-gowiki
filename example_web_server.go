package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
// Web Listener
func main() {
    http.HandleFunc("/", handler) // Handle all request with root
    log.Fatal(http.ListenAndServe(":8080", nil))
    //http://localhost:8080/willPrintThis
}