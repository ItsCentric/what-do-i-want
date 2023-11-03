package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func registerRoutes() {
    fs.WalkDir(os.DirFS("src/"), ".", func(path string, directory fs.DirEntry, err error) error {
        if path == "." {
            return nil
        }
        if err != nil {
            log.Fatal(err)
        }
        if directory.IsDir() {
            router.HandleFunc("/" + strings.TrimSuffix(path, "/index.html"), httpHandler)
        } else if path == "index.html" {
            router.HandleFunc("/", httpHandler)
        }
        return nil
    })
}
    
func httpHandler(responseWriter http.ResponseWriter, request *http.Request) {
    requestedPage := request.URL.Path + "/index.html"
    htmlTemplate := template.Must(template.ParseFiles("src/" + requestedPage))
    htmlTemplate.Execute(responseWriter, nil)
}

func main() {
    registerRoutes()
    err := http.ListenAndServe(":3000", router)
    fileSystem := http.FileServer(http.Dir("src/static/"))
    router.Handle("/src/static/", http.StripPrefix("/src/static/", fileSystem))
    if err != nil {
        fmt.Println("There was an error starting the server:", err.Error())
        os.Exit(1)
    }
}
