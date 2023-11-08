package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type PageData struct {
	Url     Url
	Product Product
}

var router = mux.NewRouter()

func main() {
	indexTemplate := template.Must(template.ParseFiles("src/index.html"))
	fileSystem := http.FileServer(http.Dir("src/static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileSystem))

	router.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		requestParameters := request.URL.Query()
		priceParameter := requestParameters.Get("price")
		returnedPrice, err := strconv.ParseInt(priceParameter, 0, 8)
		fmt.Println(returnedPrice)
		if err != nil && priceParameter != "" {
			log.Fatal("Something went wrong parsing the integer")
		}
		returnedUrl := Url{Url: requestParameters.Get("url"), Error: requestParameters.Get("error")}
		returnedProduct := Product{Name: requestParameters.Get("name"), Image: requestParameters.Get("image"), Price: returnedPrice}
		data := PageData{Url: returnedUrl, Product: returnedProduct}
		fmt.Println(data)
		indexTemplate.Execute(responseWriter, data)
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
		if url.Error != "" {
			http.Redirect(responseWriter, request, fmt.Sprintf("/?url=%s&error=%s", url.Url, url.Error), 302)
			return
		}
		product := GetProductInformation(url.Url)
		http.Redirect(responseWriter, request, fmt.Sprintf("/?name=%s&image=%s&price=%d", product.Name, product.Image, product.Price), 303)
	})

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("There was an error starting the server:", err.Error())
		os.Exit(1)
	}
}
