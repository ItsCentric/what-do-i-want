package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {
    indexTemplate := template.Must(template.ParseFiles("src/index.html"))
    fileSystem := http.FileServer(http.Dir("src/static/"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileSystem))

    router.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
        requestParameters := request.URL.Query()
        returnedUrl := Url{ Url: requestParameters.Get("url"), Error: requestParameters.Get("error") }
        indexTemplate.Execute(responseWriter, returnedUrl)
    })
    router.HandleFunc("/api/process-item", func(responseWriter http.ResponseWriter, request *http.Request) {
        if request.Method != http.MethodPost {
            indexTemplate.Execute(responseWriter, nil)
            return
        }
        url := &Url{
            Url: request.PostFormValue("url"),
        }
        url.Validate()
        http.Redirect(responseWriter, request, fmt.Sprintf("/?url=%s&error=%s", url.Url, url.Error), 302)
    })

    
    err := http.ListenAndServe(":3000", router)
    if err != nil {
        fmt.Println("There was an error starting the server:", err.Error())
        os.Exit(1)
    }
}
