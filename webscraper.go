package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Product struct {
	Name  string
	Image string
	Price int64
}

func GetProductInformation(url string) Product {
	collyClient := colly.NewCollector()
	product := Product{}

	collyClient.IgnoreRobotsTxt = false
	collyClient.Limit(&colly.LimitRule{
		Delay: 2 * time.Second,
	})
	collyClient.OnHTML("h1", func(element *colly.HTMLElement) {
		id := element.Attr("id")
		class := element.Attr("class")
		if strings.Contains(id, "title") || strings.Contains(class, "title") {
			product.Name = strings.TrimSpace(element.Text)
		}
	})
	collyClient.OnHTML("img", func(element *colly.HTMLElement) {
		if product.Image != "" {
			return
		}
		alt := element.Attr("alt")
		if alt == product.Name {
			product.Image = element.Attr("src")
		}
	})
	collyClient.OnHTML("span[class]", func(element *colly.HTMLElement) {
		if product.Price != 0 {
			return
		}
		class := element.Attr("class")
		testAttribute := element.Attr("data-test")
		if strings.Contains(strings.ToLower(class), "price") || strings.Contains(strings.ToLower(testAttribute), "price") {
			fmt.Println("------")
			fmt.Println("class:", class)
			fmt.Println("text:", element.Text)
			candidatePrice := strings.Trim(element.Text, "$.")
			price, _ := strconv.ParseInt(candidatePrice, 0, 8)
			fmt.Println(candidatePrice)
			product.Price = price
		}
	})

	collyClient.Visit(url)
	collyClient.Wait()

	return product
}
