package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func handler(responseWriter http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(responseWriter, "You requested: %s\n", request.URL.Path)
}

func main() {
    router.HandleFunc("/", handler)

    err := http.ListenAndServe(":3000", router)
    fileSystem := http.FileServer(http.Dir("src/static/"))
    router.Handle("/src/static/", http.StripPrefix("/src/static/", fileSystem))
    if err != nil {
        fmt.Println("There was an error starting the server:", err.Error())
        os.Exit(1)
    }
}
