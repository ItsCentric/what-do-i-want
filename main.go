package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(responseWriter, "You requested: %s\n", request.URL.Path)
}

func main() {
    http.HandleFunc("/", handler)

    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        fmt.Println("There was an error starting the server:", err.Error())
        os.Exit(1)
    }
}
